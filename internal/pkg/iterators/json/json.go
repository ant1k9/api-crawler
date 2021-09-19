package json

import (
	"errors"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
	"github.com/tidwall/gjson"
)

const Type = "json"

type (
	iterator struct {
		config.Iterator
		itemType string
	}
)

func New(cfg config.Iterator, itemType string) *iterator {
	return &iterator{
		Iterator: cfg,
		itemType: itemType,
	}
}

func (i *iterator) GetCollection(payload string) ([]dto.Item, error) {
	values := gjson.Get(payload, i.CollectionPath).Array()
	items := make([]dto.Item, 0, len(values))
	for _, v := range values {
		items = append(items, dto.Item{
			ID:      gjson.Get(v.Raw, i.IdentificatorPath).Int(),
			Plugin:  i.itemType,
			Payload: v.Raw,
		})
	}

	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return items, nil
}
