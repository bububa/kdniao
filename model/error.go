package model

type Error struct {
	Reason string
}

func (e Error) Error() string {
	return e.Reason
}
