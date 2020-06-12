package scorer

import (
	"context"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
)

type meanScorer struct {
	scores map[string]int
}

func NewMeanScorer(teams []string) Scorer {
	scores := make(map[string]int)
	for _, name := range teams {
		scores[name] = 0
	}
	return &meanScorer{
		scores: scores,
	}
}

func (ms *meanScorer) UpdateScores(ctx context.Context, request *scoreupdate.UpdateScoreRequest) (*scoreupdate.UpdateScoresResponse, error) {
	return nil, nil
}
