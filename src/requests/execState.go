package requests

type ExecState int

const (
	Blocked ExecState = 0
	Running ExecState = 1
	Ready   ExecState = 2
)
