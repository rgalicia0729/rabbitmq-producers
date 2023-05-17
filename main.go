package main

import (
	"fmt"
	"github.com/rgalicia0729/rabbitmq-producers/src/simple"
)

func main() {
	// Publica 5 mensajes en una cola simple
	for i := 1; i <= 5; i++ {
		content := simple.Content{
			Message: fmt.Sprintf("Message %v", i),
		}

		simple.Publish(&content)
	}
}
