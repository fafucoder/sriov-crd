package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	netDirectory     = "/sys/class/net"
	sysBusPci        = "/sys/bus/pci/devices"
	configuredVfFile = "sriov_numvfs"
	totalVfFile      = "sriov_totalvfs"
)

func GetVFInfo(pciAddr string) (string, int, error) {
	var vfID int

	pf, err := GetPFName(pciAddr)
	if err != nil {
		return "", vfID, err
	}

	vfID, err = GetVFID(pciAddr, pf)
	if err != nil {
		return "", vfID, err
	}

	return pf, vfID, nil
}

func GetPFName(pciAddr string) (string, error) {
	path := filepath.Join(sysBusPci, pciAddr, "physfn", "net")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			path := filepath.Join(sysBusPci, pciAddr, "net")
			files, err = ioutil.ReadDir(path)
			if err != nil {
				return "", err
			}
			if len(files) < 1 {
				return "", fmt.Errorf("no interface name found for device %s", pciAddr)
			}
			return files[0].Name(), nil
		}
		return "", err
	} else if len(files) > 0 {
		return files[0].Name(), nil
	}

	return "", fmt.Errorf("the PF name is not found for device %s", pciAddr)
}

func GetVFID(addr string, pfName string) (int, error) {
	var id int
	vfTotal, err := GetVFNum(pfName)
	if err != nil {
		return id, err
	}
	for vf := 0; vf <= vfTotal; vf++ {
		vfDir := filepath.Join(netDirectory, pfName, "device", fmt.Sprintf("virtfn%d", vf))
		_, err := os.Lstat(vfDir)
		if err != nil {
			continue
		}
		pciinfo, err := os.Readlink(vfDir)
		if err != nil {
			continue
		}
		pciaddr := filepath.Base(pciinfo)
		if pciaddr == addr {
			return vf, nil
		}
	}
	return id, fmt.Errorf("unable to get VF ID with PF: %s and VF pci address %v", pfName, addr)
}

func GetVFNum(ifName string) (int, error) {
	var vfTotal int

	sriovFile := filepath.Join(netDirectory, ifName, "device", configuredVfFile)
	if _, err := os.Lstat(sriovFile); err != nil {
		return vfTotal, fmt.Errorf("failed to open the sriov_numfs of device %q: %v", ifName, err)
	}

	data, err := ioutil.ReadFile(sriovFile)
	if err != nil {
		return vfTotal, fmt.Errorf("failed to read the sriov_numfs of device %q: %v", ifName, err)
	}

	if len(data) == 0 {
		return vfTotal, fmt.Errorf("no data in the file %q", sriovFile)
	}

	sriovNumfs := strings.TrimSpace(string(data))
	vfTotal, err = strconv.Atoi(sriovNumfs)
	if err != nil {
		return vfTotal, fmt.Errorf("failed to convert sriov_numfs(byte value) to int of device %q: %v", ifName, err)
	}

	return vfTotal, nil
}

func GetDriverName(pciAddr string) (string, error) {
	driverLink := filepath.Join(sysBusPci, pciAddr, "driver")
	driverInfo, err := os.Readlink(driverLink)
	if err != nil {
		return "", fmt.Errorf("error getting driver info for device %s %v", pciAddr, err)
	}

	return filepath.Base(driverInfo), nil
}

func GetPciAddress(ifName string, vf int) (string, error) {
	var pciaddr string
	vfDir := filepath.Join(netDirectory, ifName, "device", fmt.Sprintf("virtfn%d", vf))
	dirInfo, err := os.Lstat(vfDir)

	if err != nil {
		return pciaddr, fmt.Errorf("can't get the symbolic link of virtfn%d dir of the device %q: %v", vf, ifName, err)
	}

	if (dirInfo.Mode() & os.ModeSymlink) == 0 {
		return pciaddr, fmt.Errorf("No symbolic link for the virtfn%d dir of the device %q", vf, ifName)
	}

	pciInfo, err := os.Readlink(vfDir)
	if err != nil {
		return pciaddr, fmt.Errorf("can't read the symbolic link of virtfn%d dir of the device %q: %v", vf, ifName, err)
	}

	pciaddr = filepath.Base(pciInfo)
	return pciaddr, nil
}

func GetVFNames(pciAddr string) ([]string, error) {
	var names []string
	netDir := filepath.Join(sysBusPci, pciAddr, "net")
	if _, err := os.Lstat(netDir); err != nil {
		return nil, fmt.Errorf("GetNetNames: no net directory under pci device %s: %q", pciAddr, err)
	}

	fInfos, err := ioutil.ReadDir(netDir)
	if err != nil {
		return nil, fmt.Errorf("GetNetNames: failed to read net directory %s: %q", netDir, err)
	}

	names = make([]string, 0)
	for _, f := range fInfos {
		names = append(names, f.Name())
	}

	return names, nil
}

func IsSriovPF(pciAddr string) bool {
	totalVfFilePath := filepath.Join(sysBusPci, pciAddr, totalVfFile)
	if _, err := os.Stat(totalVfFilePath); err != nil {
		return false
	}

	return true
}

func IsSriovVF(pciAddr string) bool {
	totalVfFilePath := filepath.Join(sysBusPci, pciAddr, "physfn")
	if _, err := os.Stat(totalVfFilePath); err != nil {
		return false
	}
	return true
}
