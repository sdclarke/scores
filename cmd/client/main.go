package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

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
	
	fmt.Printf("\nMean Scores:\n")
	printScores(response.MeanScores)

	fmt.Printf("\nMedian Scores:\n")
	printScores(response.MedianScores)

	fmt.Printf("\nMaximum Scores:\n")
	printScores(response.MaximumScores)

	// log.Printf("Mean Scores: %#v", response.MeanScores)
	// log.Printf("Median Scores: %#v", response.MedianScores)
	// log.Printf("Maximum Scores: %#v", response.MaximumScores)
}

func printScores(m map[string]float64) {
	n := map[float64][]string{}
	var a []float64
	for k, v := range m {
			n[v] = append(n[v], k)
	}
	for k := range n {
			a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(a)))
	for _, k := range a {
			for _, s := range n[k] {
					fmt.Printf("%s, %.2f\n", s, k)
			}
	}

}
