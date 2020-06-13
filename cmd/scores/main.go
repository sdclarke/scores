package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
	"github.com/sdclarke/scores/pkg/scorer"

	"google.golang.org/grpc"
)

func main() {
	scores := struct {
		Mean    []struct{} `json:"mean"`
		Median  []struct{} `json:"median"`
		Maximum []struct{} `json:"maximum"`
	}{
		Mean:    []struct{}{},
		Median:  []struct{}{},
		Maximum: []struct{}{},
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
	scorer := scorer.NewMaximumScorer(
		scorer.NewMedianScorer(
			scorer.NewMeanScorer()))

	s := grpc.NewServer()
	scoreupdate.RegisterScorerServer(s, scorer)

	sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to create listening socket for :8080: %#v", err)
	}

	serveErrors := make(chan error)
	go func() { serveErrors <- s.Serve(sock) }()
	log.Printf("Ready")
	<-serveErrors
}
