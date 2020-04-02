package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

type PF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PFSpec `json:"spec"`
}

type PFSpec struct {
	IfName   string `json:"ifName"`
	PfName   string `json:"pfName"`
	Driver   string `json:"driver"`
	NodeName string `json:"nodeName"`
	Vendor   string `json:"vendor"`
	Product  string `json:"product"`
	VfNum    int    `json:"vfNum"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PF `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

type VF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VFSpec `json:"spec"`
}

type VFSpec struct {
	VFID         int    `json:"vfID"`
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

type VFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VF `json:"items"`
}
