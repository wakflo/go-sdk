package core

type ActionSettings struct {
	// Loop *LoopSettings `json:"loop,omitempty"`
	Branch   *BranchSettings   `json:"_"`
	Router   *RouterSettings   `json:"router,omitempty"`
	Parallel *ParallelSettings `json:"parallel,omitempty"`
}
