package v1beta1

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
	WaitAverage float32        `json:"waitAverage,omitempty"` // average wait time
	Load        ServerLoadSpec `json:"load,omitempty"`        // server load statistics
}

// Specifications of server load statistics
type ServerLoadSpec struct {
	ArrivalRate float32 `json:"arrivalRate"`          // req/min
	AvgLength   int     `json:"avgLength"`            // number of tokens
	ArrivalCOV  float32 `json:"arrivalCOV,omitempty"` // coefficient of variation of inter-request arrival time
	ServiceCOV  float32 `json:"serviceCOV,omitempty"` // coefficient of variation of request service time
}
