package scorer

import (
	"context"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
)

type meanScorer struct {
	scores map[string]float64
}

func NewMeanScorer() Scorer {
	scores := make(map[string]float64)
	return &meanScorer{
		scores: scores,
	}
}

func (ms *meanScorer) UpdateScores(ctx context.Context, request *scoreupdate.UpdateScoresRequest) (*scoreupdate.UpdateScoresResponse, error) {
	total := float64(0)
	for name, value := range request.Scores {
		ms.scores[name] += value
		total += value
	}
	mean := total / float64(len(request.Scores))
	ms.scores[request.Turn] += mean
	return &scoreupdate.UpdateScoresResponse{
		Success:    true,
		MeanScores: ms.scores,
	}, nil
}
