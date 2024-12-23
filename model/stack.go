package model

type Stack int

const (
	Standard Stack = iota
	Special
	Rejected
)

const (
	StandardStr string = "STANDARD"
	SpecialStr  string = "SPECIAL"
	RejectedStr string = "REJECTED"
)

func (s Stack) String() string {
	switch s {
	case Standard:
		return StandardStr
	case Special:
		return SpecialStr
	default:
		return RejectedStr
	}
}
