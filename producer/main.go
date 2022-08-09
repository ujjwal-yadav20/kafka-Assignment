package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "example-123",
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
		Async:                  true,
		Completion: func(messages []kafka.Message, err error) {

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, val := range messages {
				fmt.Printf("messages sent, offset %d, key %s, val %s \n", val.Offset, val.Key, val.Value)
			}
		},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key-1"),
			Value: []byte("value-1"),
		},
		kafka.Message{
			Key:   []byte("key-2"),
			Value: []byte("value-2"),
		},
		kafka.Message{
			Key:   []byte("key-3"),
			Value: []byte("value-3"),
		},
		kafka.Message{
			Key:   []byte("key-4"),
			Value: []byte("value-4"),
		},
	)

	if err != nil {
		log.Fatal("Failed in writing messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("Failed To Close Writer:", err)
	}
}
