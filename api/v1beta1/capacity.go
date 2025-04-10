package v1beta1

// Data about accelerator type availability
type CapacityData struct {
	Count []AcceleratorCount `json:"count"` // count of accelerator types
}

// Count of accelerator types in the system
type AcceleratorCount struct {
	Type  string `json:"type"`  // name of accelerator type
	Count int    `json:"count"` // number of available units
}
