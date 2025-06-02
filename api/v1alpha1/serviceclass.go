package v1alpha1

// Data related to a service class SLOs
type ServiceClassData struct {
	Spec []ServiceClassSpec `json:"serviceClasses"`
}

// Specifications of SLO data for a combination of a service class and a model
type ServiceClassSpec struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // service class name

	// +kubebuilder:validation:MinLength=1
	Model string `json:"model"` // model name

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	Priority int `json:"priority,omitempty"` // (non-negative) priority (lower value is higher priority)

	SLO_ITL float32 `json:"slo-itl,omitempty"` // inter-token latency (msec)
	SLO_TTW float32 `json:"slo-ttw,omitempty"` // request waiting time (msec)
	SLO_TPS float32 `json:"slo-tps,omitempty"` // throughput (tokens/sec)
}
