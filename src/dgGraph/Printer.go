package dgGraph

import "fmt"

var shouldPrint = true;

type Printer interface {
	toString() string
}

func PrintMessage(message ManagementMessage, receiver Printer, sender Printer){
	if shouldPrint {
		messageType := MessageTypes[message.messageType]
		fmt.Println("Event:" + messageType + " Receiver: " + receiver.toString() + " sender: " + sender.toString() )
	}
}

func PrintMessageWithoutSender(message ManagementMessage, receiver Printer){
	if shouldPrint {
		messageType := MessageTypes[message.messageType]
		fmt.Println("Event:" + messageType + " Receiver: " + receiver.toString())
	}
}