package dgGraph

type dgNodeStatus int;

const (
	entering dgNodeStatus = 0
	waiting dgNodeStatus = 1
	ready dgNodeStatus = 2
	leaving dgNodeStatus = 3
	deleting dgNodeStatus = 4
)
