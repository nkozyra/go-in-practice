package main

import (
	"context"
	"log"
	"os"
	"time"

	kafka "https://github.com/segmentio/kafka-go"
)

func main() {
	topic := "media"
	kafkaHost := os.Getenv("KAFKA_HOST")
	if kafkaHost == "" {
		panic("KAFKA_HOST environment variable not set")
	}
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, 0)
	if err != nil {
		panic(err)
	}
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)

	message := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(message)
		if err != nil {
			break
		}
	}
	log.Println(string(message))
	// get details of media upload, transcode media or pass to another service
	batch.Close()
	conn.Close()
}
