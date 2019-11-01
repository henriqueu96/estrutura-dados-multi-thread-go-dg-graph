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

type ManagementMessageType string;

const (
	enterNewNode        ManagementMessageType = "enterNewNode" //liberando nodo para entrar no grapho
	newNodeAppeared     ManagementMessageType = "newNodeAppeared" // nodo pergunta se tem conflitos
	hasConflictMessage  ManagementMessageType = "hasConflictMessage" // tem conflito
	endsConflictMessage ManagementMessageType = "endsConflictMessage" // nodo mais antigo avisa que terminou as mensagens de conflito
	decreaseConflict    ManagementMessageType = "decreaseConflict" // decrementa um conflito
	endFunc             ManagementMessageType = "endFunc" // termina execução
	leavingNode ManagementMessageType = "leavingNode" // nodo pedindo pra sair
	wantToDelete ManagementMessageType = "wantToDelete" //querendo deletar o nodo seguinte
	UpdateLastInManagementChannel = "UpdateLastInManagementChannel"
	AddRequest = "AddRequest"
)


var MessageTypes = map[ManagementMessageType] string{
	enterNewNode:   "enterNewNode",
	newNodeAppeared:  "newNodeAppeared",
	hasConflictMessage:     "hasConflictMessage",
	endsConflictMessage:    "endsConflictMessage",
	decreaseConflict:    "decreaseConflict",
	endFunc:    "endFunc",
	leavingNode: "leavingNode",
	wantToDelete: "wantToDelete",
	UpdateLastInManagementChannel: "UpdateLastInManagementChannel",
	AddRequest: "AddRequest",
}

