package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/sdclarke/scores/pkg/proto/configuration"
	"github.com/sdclarke/scores/pkg/proto/scoreupdate"

	"google.golang.org/grpc"
)

type scoresJson struct {
	Mean    []scoreMap `json:"mean"`
	Median  []scoreMap `json:"median"`
	Maximum []scoreMap `json:"maximum"`
}

type scoreMap struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

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

	meanScores := []scoreMap{}
	for k, v := range response.MeanScores {
		meanScores = append(meanScores, scoreMap{Name: k, Score: v})
	}
	medianScores := []scoreMap{}
	for k, v := range response.MedianScores {
		medianScores = append(medianScores, scoreMap{Name: k, Score: v})
	}
	maximumScores := []scoreMap{}
	for k, v := range response.MaximumScores {
		maximumScores = append(maximumScores, scoreMap{Name: k, Score: v})
	}

	scores := scoresJson{
		Mean:    meanScores,
		Median:  medianScores,
		Maximum: maximumScores,
	}

	scoresJsonString, err := json.MarshalIndent(scores, "", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal scores: %#v", err)
	}

	scoresFile, err := os.OpenFile("interface/src/assets/scores.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatalf("Failed to open file scores.json: %#v", err)
	}
	fmt.Fprintf(scoresFile, string(scoresJsonString))
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
