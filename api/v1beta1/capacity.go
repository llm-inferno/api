package v1beta1

// Data about accelerator type availability
type CapacityData struct {
	Count []AcceleratorCount `json:"count"` // count of accelerator types
}

// Count of accelerator types in the system
type AcceleratorCount struct {
	// +kubebuilder:validation:MinLength=1
	Type string `json:"type"` // name of accelerator type

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	Count int `json:"count,omitempty"` // number of available units
}
