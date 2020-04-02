package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Context("Checking GetPFName Function", func() {
		It("PF Exists", func() {
			result, err := GetPFName("0000:af:06.0")
			Expect(err).NotTo(HaveOccurred(), "Existing VF should not return an error")
			Expect(result).To(Equal("enp175s0f1"), "Existing VF should return correct PF name")
		})

		It("PF Not Exists", func() {
			result, err := GetPFName("0000:af:07.0")
			Expect(result).To(Equal(""))
			Expect(err).To(HaveOccurred(), "Not Exists VF Should return error")
		})
	})

	Context("Checking GetVFID function", func() {
		It("VF ID Exists", func() {
			result, err := GetVFID("0000:af:06.0", "enp175s0f1")
			Expect(result).To(Equal(0), "Existing VF should return correct VF index")
			Expect(err).NotTo(HaveOccurred(), "Existing VF should not return an error")
		})
		It("VF ID Not Exists", func() {
			_, err := GetVFID("0000:af:06.0", "enp175s0f2")
			Expect(err).To(HaveOccurred(), "Not existing interface should return an error")
		})
	})

	Context("Checking GetVFInfo function", func() {
		It("VF Pci Exists", func() {
			result, id, err := GetVFInfo("0000:af:06.0")
			Expect(id).To(Equal(0), "Existing VF should return correct VF index")
			Expect(err).NotTo(HaveOccurred(), "Existing VF should not return an error")
			Expect(result).To(Equal("enp175s0f1"), "Existing VF should return correct PF name")
		})

		It("VF Pci Not Exists", func() {
			_, _, err := GetVFInfo("0000:af:06.6")
			Expect(err).To(HaveOccurred(), "Not Existing VF should not return an error")
		})
	})

	Context("Checking GetVFNum Function", func() {
		It("PF Exists", func() {
			result, err := GetVFNum("enp175s0f1")
			Expect(result).To(Equal(2), "Existing PF should return correct VFs count")
			Expect(err).NotTo(HaveOccurred(), "Existing PF should not return an error")
		})
		It("PF Not Exists", func() {
			_, err := GetVFNum("enp175s0f2")
			Expect(err).To(HaveOccurred(), "Not existing PF should return an error")
		})
	})

	Context("Checking GetVFName Function", func() {
		It("VF Pci Exists", func() {
			result, err := GetVFNames("0000:af:06.0")
			Expect(result).To(ContainElement("enp175s6"), "Existing VF Pci should return vf name")
			Expect(err).NotTo(HaveOccurred(), "Existing VF Pci should not return an error")
		})
		It("VF Pci Not Exists", func() {
			_, err := GetVFNum("0000:af:06.4")
			Expect(err).To(HaveOccurred(), "Not Existing VF Pci return an error")
		})
	})

	Context("Checking IsSriovPF Function", func() {
		It("Is Sriov PF", func() {
			exists := IsSriovPF("0000:af:00.1")
			Expect(exists).To(BeTrue(), "Is Sriov PF")

			exists = IsSriovPF("0000:af:00.2")
			Expect(exists).NotTo(BeTrue(), "Not Is Sriov PF")
		})
	})

	Context("Checking IsSriovVF Function", func() {
		It("Is Sriov VF", func() {
			exists := IsSriovVF("0000:af:06.0")
			Expect(exists).To(BeTrue(), "Is Sriov VF")

			exists = IsSriovVF("0000:af:06.6")
			Expect(exists).NotTo(BeTrue(), "Not Is Sriov VF")
		})
	})
})
