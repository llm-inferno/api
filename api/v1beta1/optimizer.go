package v1beta1

// Data related to Optimizer
type OptimizerData struct {
	Spec OptimizerSpec `json:"optimizer"`
}

// Specifications for optimizer data
type OptimizerSpec struct {
	Unlimited     bool `json:"unlimited"`     // unlimited number of accelerator types (for capacity planning and/or cloud)
	Heterogeneous bool `json:"heterogeneous"` // heterogeneous accelerators assigned to same inference server
	MILPSolver    bool `json:"milpSolver"`    // use MILP solver to optimize
	UseCplex      bool `json:"useCplex"`      // use CPLEX solver for MILP problem
}

type AllocationSolution struct {
	Spec map[string]AllocationData `json:"allocations"` // map of server names to allocation data
}
