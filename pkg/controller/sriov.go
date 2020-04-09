package controller

import (
	"fmt"
	cnitypes "github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/fafucoder/sriov-crd/pkg/apis/sriov"
	v1 "github.com/fafucoder/sriov-crd/pkg/apis/sriov/v1"
	sriovclient "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned/typed/sriov/v1"
	"github.com/fafucoder/sriov-crd/pkg/utils"
	"github.com/intel/multus-cni/logging"
	"github.com/jaypipes/ghw"
	"github.com/vishvananda/netlink"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"os"
	"strconv"
)

const (
	netClass = 0x02
)

func CreateSriovVfCrd(client sriovclient.K8sCniCncfIoV1Interface, vfSpec v1.VFSpec, name string) error {
	if client == nil {
		return fmt.Errorf("no client set")
	}

	vfCrd, err := client.VFs().Get(name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			_, err := client.VFs().Create(&v1.VF{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
				Spec: vfSpec,
			})
			if err != nil {
				return fmt.Errorf("failed create sriov vf crd: %v", err)
			}

			return nil
		}

		return err
	}

	vfCrd.Spec = vfSpec
	_, err = client.VFs().Update(vfCrd)

	return err
}

func DeleteSriovVfCrd(client sriovclient.K8sCniCncfIoV1Interface, name string) error {
	if client == nil {
		return fmt.Errorf("not client set")
	}
	err := client.VFs().Delete(name, &metav1.DeleteOptions{})
	if err != nil && !k8serrors.IsNotFound(err) {
		return err
	}

	return nil
}

func CreateSriovVfSpec(r cnitypes.Result, podName, podNamespace, deviceID string) (v1.VFSpec, error) {
	vfSpec := v1.VFSpec{}

	result, err := current.NewResultFromResult(r)
	if err != nil {
		return vfSpec, fmt.Errorf("error convert the type.Result to current.Result: %v", err)
	}

	if deviceID == "" {
		return vfSpec, fmt.Errorf("the device id is empty")
	}

	pf, vf, err := utils.GetVFInfo(deviceID)
	logging.Errorf("get vf info: %v %v", pf, vf)
	if err != nil {
		return vfSpec, fmt.Errorf("failed get vf info:%v", err)
	}

	vfSpec.VFID = vf
	vfSpec.PFName = pf
	vfSpec.DeviceID = deviceID
	vfSpec.PodNamespace = podNamespace
	vfSpec.PodName = podName
	vfSpec.NodeName = os.Getenv("KUBE_NODE_NAME")

	vfNames, err := utils.GetVFNames(deviceID)
	if err == nil && len(vfNames) > 0 {
		vfSpec.VFName = vfNames[0]
	}

	for _, ifs := range result.Interfaces {
		if ifs.Sandbox != "" {
			vfSpec.NetName = ifs.Name
			vfSpec.MacAddress = ifs.Mac
		}
	}

	return vfSpec, nil
}

func CreateSriovPfCrd(client kubernetes.Interface, sriovClient sriovclient.K8sCniCncfIoV1Interface) error {
	nodeList, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return logging.Errorf("failed get node list: %v", err)
	}

	for _, node := range nodeList.Items {
		sriovPfSpec, err := discoverHostDevices(node.Name)
		if err != nil {
			logging.Errorf("failed get this host devices: %v", err)
			continue
		}

		for _, pfSpec := range sriovPfSpec {
			if pfSpec.NodeName == "" || pfSpec.PfName == "" {
				continue
			}

			pfCrd, err := sriovClient.PFs().Get(fmt.Sprintf("%s.%s", pfSpec.PfName, pfSpec.NodeName), metav1.GetOptions{})
			if err != nil {
				if k8serrors.IsNotFound(err) {
					apiVersion := fmt.Sprintf("%s/%s", sriov.GroupName, "v1")
					_, err := sriovClient.PFs().Create(&v1.PF{
						ObjectMeta: metav1.ObjectMeta{
							Name: fmt.Sprintf("%s.%s", pfSpec.PfName, pfSpec.NodeName),
						},
						TypeMeta: metav1.TypeMeta{
							APIVersion: apiVersion,
							Kind:       "PF",
						},
						Spec: pfSpec,
					})
					if err != nil {
						klog.Errorf("failed create sriov pf crd: %v", err)
						return err
					}

					return nil
				}

				return err
			}

			pfCrd.Spec = pfSpec
			_, err = sriovClient.PFs().Update(pfCrd)

			return err
		}
	}

	return err
}

func discoverHostDevices(nodeName string) ([]v1.PFSpec, error) {
	var pfList []v1.PFSpec

	pci, err := ghw.PCI()
	if err != nil {
		return pfList, logging.Errorf("error getting PCI Info: %v", err)
	}

	devices := pci.ListDevices()
	if len(devices) == 0 {
		return pfList, nil
	}

	for _, device := range devices {
		devClass, err := strconv.ParseInt(device.Class.ID, 16, 64)
		if err != nil {
			logging.Errorf("failed get device class %v", err)
			continue
		}

		// only interested in network class
		if devClass != netClass {
			continue
		}

		vendor := device.Vendor
		vendorName := vendor.Name
		if len(vendor.Name) > 20 {
			vendorName = string([]byte(vendorName)[0:17]) + "..."
		}

		product := device.Product
		productName := product.Name
		if len(product.Name) > 40 {
			productName = string([]byte(productName)[0:37]) + "..."
		}

		//validate device has default route, exclude device in-use in host
		if isDefaultRoute, _ := hasDefaultRoute(device.Address); isDefaultRoute {
			continue
		}

		if utils.IsSriovPF(device.Address) {
			var vfNum int
			var ifName string

			pfName, err := utils.GetPFName(device.Address)
			if err != nil {
				logging.Errorf("failed get PF name: %v", err)
				continue
			}

			if !utils.IsSriovVF(device.Address) {
				vfNum, err = utils.GetVFNum(pfName)
				if err != nil {
					logging.Errorf("failed get vf num: %v", err)
					continue
				}
			}

			driverName, err := utils.GetDriverName(device.Address)
			if err != nil {
				logging.Errorf("failed get driver name: %v", err)
				continue
			}

			netDevs, _ := utils.GetVFNames(device.Address)
			if len(netDevs) > 0 {
				ifName = netDevs[0]
			}

			pf := v1.PFSpec{
				IfName:   ifName,
				PfName:   pfName,
				Driver:   driverName,
				Product:  productName,
				Vendor:   vendorName,
				VfNum:    vfNum,
				NodeName: nodeName,
			}

			logging.Debugf("discoverHostDevices: get pf spec, %#v", pf)
			pfList = append(pfList, pf)
		}
	}

	return pfList, nil
}

func hasDefaultRoute(pciAddr string) (bool, error) {
	ifNames, err := utils.GetVFNames(pciAddr)
	if err != nil {
		return false, fmt.Errorf("error trying get net device name for device %s", pciAddr)
	}

	if len(ifNames) > 0 { // there's at least one interface name found
		for _, ifName := range ifNames {
			link, err := netlink.LinkByName(ifName)
			if err != nil {
				continue
			}

			//get route list, check route dest name is empty
			routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
			for _, r := range routes {
				if r.Dst == nil {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
