package controller

import (
	cnicurrent "github.com/containernetworking/cni/pkg/types/current"
	"github.com/fafucoder/sriov-crd/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

var _ = Describe("Controller", func() {
	client := testing.SriovClient()

	Context("Checking CreateSriovVfSpec Function", func() {
		It("Exists VF Pci", func() {
			vfSpec, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Exists Device Should return vf spec")
			Expect(vfSpec.PodName).To(Equal("podName"))
			Expect(vfSpec.PodNamespace).To(Equal("podNamespace"))
			Expect(vfSpec.PFName).To(Equal("enp175s0f1"))
			Expect(vfSpec.VFID).To(Equal(0))
		})

		It("Exists Node Environment", func() {
			err := os.Setenv("KUBE_NODE_NAME", "node-2")
			Expect(err).NotTo(HaveOccurred())

			vfSpec, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Exists Device Should return vf spec")
			Expect(vfSpec.PodName).To(Equal("podName"))
			Expect(vfSpec.PodNamespace).To(Equal("podNamespace"))
			Expect(vfSpec.PFName).To(Equal("enp175s0f1"))
			Expect(vfSpec.VFID).To(Equal(0))
			Expect(vfSpec.NodeName).To(Equal("node-2"))
		})

		It("Exists Result", func() {
			cniResult := &cnicurrent.Result{
				CNIVersion: "0.4.0",
				Interfaces: []*cnicurrent.Interface{
					&cnicurrent.Interface{
						Name:    "net1",
						Mac:     "92:79:27:01:7c:cf",
						Sandbox: "/proc/1123/ns/net",
					},
				},
			}

			vfSpec, err := CreateSriovVfSpec(cniResult, "podName", "podNamespace", "0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Exists Device Should return vf spec")
			Expect(vfSpec.PodName).To(Equal("podName"))
			Expect(vfSpec.PodNamespace).To(Equal("podNamespace"))
			Expect(vfSpec.PFName).To(Equal("enp175s0f1"))
			Expect(vfSpec.VFID).To(Equal(0))
			Expect(vfSpec.NetName).To(Equal("net1"))
			Expect(vfSpec.MacAddress).To(Equal("92:79:27:01:7c:cf"))
		})

		It("VF Pci Not Exists", func() {
			_, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "0000:af:06.6")
			Expect(err).To(HaveOccurred(), "Not Exists  VF Pci Should return an error")
		})

		It("VF Pci Is Empty", func() {
			_, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "")
			Expect(err).To(HaveOccurred(), "Empty VF Pci Should return an error")
		})
	})

	Context("Checking CreateSriovVfCrd Function", func() {
		It("Create VF", func() {
			vfSpec, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Exists Device Should return vf spec")

			err = CreateSriovVfCrd(client, vfSpec, "sample-vfcrd")
			Expect(err).NotTo(HaveOccurred())

			vf, err := client.VFs().Get("sample-vfcrd", metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(vf.Spec.PodName).To(Equal("podName"))
			Expect(vf.Spec.PodNamespace).To(Equal("podNamespace"))
			Expect(vf.Spec.PFName).To(Equal("enp175s0f1"))
			Expect(vf.Spec.VFID).To(Equal(0))
		})
	})

	Context("Checking DeleteSriovVfCrd Function", func() {
		It("Delete SriovVF", func() {
			vfSpec, err := CreateSriovVfSpec(&cnicurrent.Result{CNIVersion: "0.4.0"}, "podName", "podNamespace", "0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Exists Device Should return vf spec")

			err = CreateSriovVfCrd(client, vfSpec, "sample-vfcrd")
			Expect(err).NotTo(HaveOccurred())

			_, err = client.VFs().Get("sample-vfcrd", metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())

			err = client.VFs().Delete("sample-vfcrd", &metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred())

			_, err = client.VFs().Get("sample-vfcrd", metav1.GetOptions{})
			Expect(err).To(HaveOccurred())
		})
	})
})
