package v1beta1

// All data related to the system (accelerators, models, service classes, ...)
type SystemData struct {
	Spec SystemSpec `json:"system"`
}

// Specifications for system data
type SystemSpec struct {
	// static data
	Accelerators   AcceleratorData  `json:"acceleratorData"`  // accelerator data
	Models         ModelData        `json:"modelData"`        // model data
	ServiceClasses ServiceClassData `json:"serviceClassData"` // service class data
	Servers        ServerData       `json:"serverData"`       // server data
	Optimizer      OptimizerData    `json:"optimizerData"`    // optimizer data

	// dynamic data
	Capacity CapacityData `json:"capacityData"` // data about accelerator type availability
}
