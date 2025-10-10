package main

import (
	"context"
	desc "github.com/olezhek28/microservices_course/week_1/grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	noteId  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}

	defer conn.Close()

	client := desc.NewNoteV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	read, err := client.Get(ctx, &desc.GetRequest{Id: noteId})
	if err != nil {
		log.Fatalf("could not get note: %v", err)
	}

	log.Printf("note: %v", read)
}
