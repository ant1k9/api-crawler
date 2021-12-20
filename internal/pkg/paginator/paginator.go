package paginator

import (
	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/errors"
	"github.com/ant1k9/api-crawler/internal/pkg/paginator/list"
	"github.com/ant1k9/api-crawler/internal/pkg/paginator/numeric"
)

type (
	Paginator interface {
		NextPage(payload string) (string, error)
	}

	PluginAdjuster interface {
		AdjustPlugin(plugin string) string
	}
)

func New(cfg config.Paginator) (Paginator, error) {
	switch cfg.Type {
	case numeric.Type:
		return numeric.New(cfg), nil
	case list.Type:
		return list.New(cfg), nil
	default:
		return nil, errors.ErrPaginatorDoesNotExist
	}
}
