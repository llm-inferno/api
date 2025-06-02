package v1alpha1

// Data related to a Model
type ModelData struct {
	PerfData []ModelAcceleratorPerfData `json:"models"` // performance data for model on accelerators
}

// Specifications for a combination of a model and accelerator data
type ModelAcceleratorPerfData struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // model name

	Acc          string  `json:"acc"`          // accelerator name
	AccCount     int     `json:"accCount"`     // number of accelerator units used by model
	Alpha        float32 `json:"alpha"`        // alpha parameter of ITL
	Beta         float32 `json:"beta"`         // beta parameter of ITL
	MaxBatchSize int     `json:"maxBatchSize"` // max batch size based on average number of tokens per request
	AtTokens     int     `json:"atTokens"`     // average number of tokens per request assumed in max batch size calculation
}
