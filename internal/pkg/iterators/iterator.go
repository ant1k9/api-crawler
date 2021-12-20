package iterator

import (
	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/errors"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/csv"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/json"
)

type (
	Iterator interface {
		GetCollection(payload string) ([]dto.Item, error)
	}
)

func New(cfg config.Iterator, itemType string) (Iterator, error) {
	switch cfg.Type {
	case json.Type:
		return json.New(cfg, itemType), nil
	case csv.Type:
		return csv.New(cfg, itemType), nil
	default:
		return nil, errors.ErrIteratorDoesNotExist
	}
}
