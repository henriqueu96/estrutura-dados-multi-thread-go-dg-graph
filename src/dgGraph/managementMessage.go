package dgGraph

type ManagementMessage struct {
	messageType ManagementMessageType
	parameter interface{}
}

func NewManagementMessage(messageType ManagementMessageType, parameter interface{}) ManagementMessage{
	return ManagementMessage{
		messageType: messageType,
		parameter:  parameter,
	}
}


type ManagementMessageType int;

const (
	newNodeAppeared   ManagementMessageType = 0
	incrementConflict ManagementMessageType = 1
	decreaseConflict  ManagementMessageType = 2
)