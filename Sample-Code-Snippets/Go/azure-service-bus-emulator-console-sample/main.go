package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {
	connectionString := "Endpoint=sb://127.0.0.1;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
	queueName := "queue.1"
	// Create a Service Bus client
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}

	// Send a message
	sendMessage(client, queueName, "Hello, Azure Service Bus!")

	// Receive messages
	receiveMessages(client, queueName)
}

func sendMessage(client *azservicebus.Client, queueName, message string) {
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("failed to create sender: %s", err)
	}
	defer sender.Close(context.Background())

	msg := &azservicebus.Message{
		Body: []byte(message),
	}

	err = sender.SendMessage(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	fmt.Println("Message sent:", message)
}

func receiveMessages(client *azservicebus.Client, queueName string) {
	receiver, err := client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		log.Fatalf("failed to create receiver: %s", err)
	}
	defer receiver.Close(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messages, err := receiver.ReceiveMessages(ctx, 1, nil)
	if err != nil {
		log.Fatalf("failed to receive messages: %s", err)
	}

	for _, msg := range messages {
		fmt.Println("Received message:", string(msg.Body))
		receiver.CompleteMessage(ctx, msg, nil)
	}
}
