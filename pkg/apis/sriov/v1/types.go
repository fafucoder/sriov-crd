package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

type SriovPF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SriovPFSpec `json:"spec"`
}

type SriovPFSpec struct {
	IfName   string `json:"ifName"`
	PfName   string `json:"pfName"`
	Driver   string `json:"driver"`
	NodeName string `json:"nodeName"`
	Vendor   string `json:"vendor"`
	Product  string `json:"product"`
	VfNum    int    `json:"vfNum"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SriovPFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SriovPF `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

type SriovVF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SriovVFSpec `json:"spec"`
}

type SriovVFSpec struct {
	DeviceID     string `json:"deviceID"`
	PFName       string `json:"pfName"`
	PodName      string `json:"podName"`
	PodNamespace string `json:"podNamespace"`
	NodeName     string `json:"nodeName"`
	NetName      string `json:"netName"`
	IPAddress    string `json:"ipAddress"`
	MacAddress   string `json:"macAddress"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SriovVFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SriovVF `json:"items"`
}
