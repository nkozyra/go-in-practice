package ch11

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var opts []grpc.DialOption

func ConnectToGRPC() error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := NewChatServiceClient(conn)

	meta, err := client.RouteComments(context.Background(), &Comment{Username: "Nick", Text: "Hello World!", Sent: timestamppb.New(time.Now())})
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Message length: %d, previous comments: %d", meta.CommentLength, meta.PreviousCommentCount))

	defer conn.Close()
	return nil
}
