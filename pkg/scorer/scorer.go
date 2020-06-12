package scorer

import (
	"context"

	"github.com/sdclarke/scores/pkg/proto/scoreupdate"
)

type Scorer interface {
	UpdateScores(ctx context.Context, request *scoreupdate.UpdateScoresRequest) (*scoreupdate.UpdateScoresResponse, error)
}
