package json

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
	"github.com/tidwall/gjson"
)

const Type = "json"

type (
	iterator struct {
		config.Iterator
		itemType string
		re       *regexp.Regexp
	}
)

func New(cfg config.Iterator, itemType string) *iterator {
	var re *regexp.Regexp
	if cfg.Regexp != "" {
		re = regexp.MustCompile(cfg.Regexp)
	}

	return &iterator{
		Iterator: cfg,
		itemType: itemType,
		re:       re,
	}
}

func (i *iterator) GetCollection(payload string) (items []dto.Item, err error) {
	values := gjson.Get(payload, i.CollectionPath).Array()

	items = make([]dto.Item, 0, len(values))
	for _, v := range values {
		id := gjson.Get(v.Raw, i.IdentificatorPath).Int()

		if i.re != nil {
			m := i.re.FindStringSubmatch(gjson.Get(v.Raw, i.IdentificatorPath).String())
			if len(m) < 2 {
				return nil, errors.New("improper regex for id")
			}
			id, err = strconv.ParseInt(m[1], 10, 64)
			if err != nil {
				return nil, err
			}
		}

		items = append(items, dto.Item{
			ID:      id,
			Plugin:  i.itemType,
			Payload: v.Raw,
		})
	}

	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return items, nil
}
