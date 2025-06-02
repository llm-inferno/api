package v1alpha1

// Data related to an Accelerator
type AcceleratorData struct {
	Spec []AcceleratorSpec `json:"accelerators"` // accelerator specs
}

// Specifications for accelerator data
type AcceleratorSpec struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // name of accelerator

	// +kubebuilder:validation:MinLength=1
	Type string `json:"type"` // name of accelerator type (e.g. A100)

	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	Multiplicity int `json:"multiplicity,omitempty"` // number of cards of type for this accelerator

	MemSize int       `json:"memSize,omitempty"` // GB
	MemBW   int       `json:"memBW,omitempty"`   // GB/sec
	Power   PowerSpec `json:"power,omitempty"`   // power consumption specs
	Cost    float32   `json:"cost"`              // cents/hr
}

// Specifications for Accelerator power consumption data (Watts)
type PowerSpec struct {
	// +kubebuilder:validation:Minimum=0
	Idle int `json:"idle"` // idle power

	// +kubebuilder:validation:Minimum=0
	Full int `json:"full"` // full utilization power

	// +kubebuilder:validation:Minimum=0
	MidPower int `json:"midPower"` // power at inflection point

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	MidUtil float32 `json:"midUtil"` // utilization at inflection point
}
