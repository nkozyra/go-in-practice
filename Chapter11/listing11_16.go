package ch11

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func ConnectToGRPC() error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // # A
	if err != nil {
		return err
	}

	client := NewChatServiceClient(conn) // # B

	meta, err := client.RouteComments(context.Background(), &CommentRequest{Username: "Nick", Text: "Hello World!", Sent: timestamppb.New(time.Now())}) // # C
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Message length: %d, previous comments: %d", meta.CommentLength, meta.PreviousCommentCount)) // # D

	defer conn.Close()
	return nil
}
