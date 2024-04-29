package pdfgenerator

type Action string

const (
	ActionUp    Action = "up"
	ActionDown  Action = "down"
	ActionLeft  Action = "left"
	ActionRight Action = "right"
)

func (d Action) String() string {
	return string(d)
}

func (d Action) Valid() bool {
	switch d {
	case ActionLeft, ActionRight, ActionUp, ActionDown:
		return true
	default:
		return false
	}
}

func (d Action) Reverse() Action {
	switch d {
	case ActionLeft:
		return ActionRight
	case ActionRight:
		return ActionLeft
	case ActionUp:
		return ActionDown
	case ActionDown:
		return ActionUp
	default:
		return ActionDown
	}
}
