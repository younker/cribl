package model

// An object can be dispatched to a stack if we can reason about it's mass and
// size
type Dispatchable interface {
	IsHeavy() bool
	IsBulky() bool
}

// Dispatch rules:
//
// - STANDARD: standard packages (those that are not bulky or heavy) can be
// handled normally.
//
// - SPECIAL: packages that are either heavy or bulky can't be handled
// automatically.
//
// - REJECTED: packages that are both heavy and bulky are rejected.
func Dispatch(obj Dispatchable) Stack {
	switch {
	case obj.IsHeavy() && obj.IsBulky():
		return Rejected
	case obj.IsHeavy() || obj.IsBulky():
		return Special
	default:
		return Standard
	}
}
