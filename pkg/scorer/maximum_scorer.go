package scorer

import (
	"context"
	"sort"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
)

type maximumScorer struct {
	base   Scorer
	scores map[string]float64
}

func NewMaximumScorer(base Scorer) Scorer {
	scores := make(map[string]float64)
	return &maximumScorer{
		base:   base,
		scores: scores,
	}
}

func (ms *maximumScorer) UpdateScores(ctx context.Context, request *scoreupdate.UpdateScoresRequest) (*scoreupdate.UpdateScoresResponse, error) {
	values := []float64{}
	for name, value := range request.Scores {
		ms.scores[name] += value
		values = append(values, value)
	}
	sort.Float64s(values)
	ms.scores[request.Turn] += values[len(values)-1]
	response, err := ms.base.UpdateScores(ctx, request)
	if err != nil {
		return response, err
	}
	response.MaximumScores = ms.scores
	return response, nil
}
