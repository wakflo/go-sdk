package core

type ActionSettings struct {
	// Loop *LoopSettings `json:"loop,omitempty"`
	Branch *BranchSettings `json:"branch,omitempty"`
	Router *RouterSettings `json:"router,omitempty"`
}
