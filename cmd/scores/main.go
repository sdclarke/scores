package main

import (
	"log"
	"net"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
	"github.com/sdclarke/scores/pkg/scorer"

	"google.golang.org/grpc"
)

func main() {
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
