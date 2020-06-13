package main

import (
	"context"
	"log"
	"os"

	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/sdclarke/scores/pkg/proto/configuration"
	"github.com/sdclarke/scores/pkg/proto/scoreupdate"

	"google.golang.org/grpc"
)

func main() {
	var configuration configuration.ClientConfiguration
	if err := util.UnmarshalConfigurationFromFile(os.Args[1], &configuration); err != nil {
		log.Fatalf("Failed to read configuration from %s: %s", os.Args[1], err)
	}
	address := configuration.Address
	ctx := context.Background()
	c, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial %#v: %#v", address, err)
	}
	client := scoreupdate.NewScorerClient(c)

	updateRequest := &scoreupdate.UpdateScoresRequest{
		Scores: configuration.Scores,
		Turn:   configuration.Turn,
	}

	response, err := client.UpdateScores(ctx, updateRequest)
	if err != nil {
		log.Fatalf("Something went wrong when updating scores: %#v", err)
	}
	log.Printf("Mean Scores: %#v", response.MeanScores)
	log.Printf("Median Scores: %#v", response.MedianScores)
	log.Printf("Maximum Scores: %#v", response.MaximumScores)
}
