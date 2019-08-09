package requests

type RequestType string;

const (
	Write RequestType = "Write"
	Read  RequestType = "Read"
)
