package kdniao

type Error struct {
	Reason string
}

func (this Error) Error() string {
	return this.Reason
}
