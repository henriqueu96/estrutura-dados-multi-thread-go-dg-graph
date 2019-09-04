package dgGraph

type ManagementMessage struct {
	messageType ManagementMessageType
	parameter   interface{}
}

func NewManagementMessage(messageType ManagementMessageType, parameter interface{}) ManagementMessage {
	return ManagementMessage{
		messageType: messageType,
		parameter:   parameter,
	}
}

type ManagementMessageType int;

const (
	enterNewNode        ManagementMessageType = 0 //liberando nodo para entrar no grapho
	newNodeAppeared     ManagementMessageType = 1 // nodo pergunta se tem conflitos
	hasConflictMessage  ManagementMessageType = 2 // tem conflito
	endsConflictMessage ManagementMessageType = 3 // nodo mais antigo avisa que terminou as mensagens de conflito
	decreaseConflict    ManagementMessageType = 4 // decrementa um conflito
	endFunc             ManagementMessageType = 5 // termina execução



)
