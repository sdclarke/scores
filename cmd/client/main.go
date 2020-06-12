package main

import (
	"context"
	"log"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	c, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial localhost:8080: %#v", err)
	}
	client := scoreupdate.NewScorerClient(c)
	log.Printf("%#v", client)

	updateRequest := &scoreupdate.UpdateScoresRequest{
		Scores: map[string]float64{
			"scott": 1,
			"snow":  2,
			"john":  4,
		},
		Turn: "jeff",
	}

	response, err := client.UpdateScores(ctx, updateRequest)
	log.Printf("%#v, %#v", response, err)
}
