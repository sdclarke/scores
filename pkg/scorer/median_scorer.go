package scorer

import (
	"context"
	"sort"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
)

type medianScorer struct {
	base   Scorer
	scores map[string]float64
}

func NewMedianScorer(base Scorer) Scorer {
	scores := make(map[string]float64)
	return &medianScorer{
		base:   base,
		scores: scores,
	}
}

func (ms *medianScorer) UpdateScores(ctx context.Context, request *scoreupdate.UpdateScoresRequest) (*scoreupdate.UpdateScoresResponse, error) {
	values := []float64{}
	for name, value := range request.Scores {
		ms.scores[name] += value
		values = append(values, value)
	}
	sort.Float64s(values)
	var median float64
	if length := len(values); length%2 == 1 {
		median = values[(length-1)/2]
	} else {
		median = (values[length/2] + values[(length/2)-1]) / 2
	}
	ms.scores[request.Turn] += median
	response, err := ms.base.UpdateScores(ctx, request)
	if err != nil {
		return response, err
	}
	response.MedianScores = ms.scores
	return response, nil
}
