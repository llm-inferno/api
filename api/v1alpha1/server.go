package v1alpha1

// Data related to a Server
type ServerData struct {
	Spec []ServerSpec `json:"servers"`
}

// Specifications for a server
type ServerSpec struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // server name

	// +kubebuilder:validation:MinLength=1
	Class string `json:"class"` // service class name

	// +kubebuilder:validation:MinLength=1
	Model string `json:"model"` // model name

	KeepAccelerator bool `json:"keepAccelerator"` // option to not change accelerator
	MinNumReplicas  int  `json:"minNumReplicas"`  // minimum number of replicas
	MaxBatchSize    int  `json:"maxBatchSize"`    // overriding value for the maximum batch size

	CurrentAlloc AllocationData `json:"currentAlloc,omitempty"` // current allocation
	DesiredAlloc AllocationData `json:"desiredAlloc,omitempty"` // desired allocation
}

// Data about a server allocation
type AllocationData struct {
	Accelerator string         `json:"accelerator"`           // accelerator name
	NumReplicas int            `json:"numReplicas"`           // number of replicas
	MaxBatch    int            `json:"maxBatch,omitempty"`    // max batch size
	Cost        float32        `json:"cost,omitempty"`        // cost of allocation
	ITLAverage  float32        `json:"itlAverage,omitempty"`  // average ITL
	TTFTAverage float32        `json:"ttftAverage,omitempty"` // average TTFT
	Load        ServerLoadSpec `json:"load,omitempty"`        // server load statistics
}

// Specifications of server load statistics
type ServerLoadSpec struct {
	ArrivalRate  float32 `json:"arrivalRate"`  // req/min
	AvgInTokens  int     `json:"avgInTokens"`  // average number of input tokens
	AvgOutTokens int     `json:"avgOutTokens"` // average number of output tokens
}
