package v1alpha1

// Data related to a Model
type ModelData struct {
	PerfData []ModelAcceleratorPerfData `json:"models"` // performance data for model on accelerators
}

// Specifications for a combination of a model and accelerator data
type ModelAcceleratorPerfData struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // model name

	Acc          string       `json:"acc"`          // accelerator name
	AccCount     int          `json:"accCount"`     // number of accelerator units used by model
	MaxBatchSize int          `json:"maxBatchSize"` // max batch size based on average number of tokens per request
	AtTokens     int          `json:"atTokens"`     // average number of tokens per request assumed in max batch size calculation
	DecodeParms  DecodeParms  `json:"decodeParms"`  // parameters for estimating decode time
	PrefillParms PrefillParms `json:"prefillParms"` // parameters for estimating prefill time
}

// Parameters for estimating decode time = alpha + beta * batchSize (msec); batchSize > 0
type DecodeParms struct {
	Alpha float32 `json:"alpha"` // base
	Beta  float32 `json:"beta"`  // slope
}

// Parameters for estimating prefill time = gamma + delta * inputTokens * batchSize (msec); inputTokens, batchSize > 0
type PrefillParms struct {
	Gamma float32 `json:"gamma"` // base
	Delta float32 `json:"delta"` // slope
}
