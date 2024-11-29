package runner

type ActionType int

const (
	ATPlan ActionType = iota
	AtApply
	ATLock
	ATUnlock
)

type Action struct {
	Type ActionType `json:"type"`
	// todo
}
