package paginator

import (
	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/errors"
	"github.com/ant1k9/api-crawler/internal/pkg/paginator/numeric"
)

type (
	Paginator interface {
		NextPage(payload string) (string, error)
	}
)

func New(cfg config.Paginator) (Paginator, error) {
	switch cfg.Type {
	case numeric.Type:
		return numeric.New(cfg), nil
	default:
		return nil, errors.ErrPaginatorDoesNotExist
	}
}
