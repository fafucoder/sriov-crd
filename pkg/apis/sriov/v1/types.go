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
	IfName   string `json:"if_name"`
	PfName   string `json:"pf_name"`
	Driver   string `json:"driver"`
	NodeName string `json:"node_name"`
	Vendor   string `json:"vendor"`
	Product  string `json:"product"`
	VfNum    int    `json:"vf_id"`
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
	DeviceID     string `json:"device_id"`
	PFName       string `json:"pf_name"`
	PodName      string `json:"pod_name"`
	PodNamespace string `json:"pod_namespace"`
	NodeName     string `json:"node_name"`
	NetName      string `json:"net_name"`
	IPAddress    string `json:"ip_address"`
	MacAddress   string `json:"mac_address"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SriovVFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SriovVF `json:"items"`
}
