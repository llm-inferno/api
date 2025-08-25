package v1alpha1

// Data related to a service class SLOs
type ServiceClassData struct {
	Spec []ServiceClassSpec `json:"serviceClasses"`
}

// Specification of a service class
type ServiceClassSpec struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"` // service class name

	// +kubebuilder:default=100
	// +kubebuilder:validation:Minimum=0
	Priority int `json:"priority"` // [1,100] priority (lower value is higher priority)

	ModelTargets []ModelTarget `json:"modelTargets"` // target SLOs for models
}

// Specification of SLO targets for a model
type ModelTarget struct {
	// +kubebuilder:validation:MinLength=1
	Model string `json:"model"` // model name

	SLO_ITL float32 `json:"slo-itl,omitempty"` // inter-token latency (msec)
	SLO_TTW float32 `json:"slo-ttw,omitempty"` // request waiting time (msec)
	SLO_TPS float32 `json:"slo-tps,omitempty"` // throughput (tokens/sec)
}
