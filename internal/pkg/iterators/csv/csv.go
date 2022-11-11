package csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
)

const Type = "csv"
const defaultSeparator = ","

type (
	iterator struct {
		config.Iterator
		itemType string
	}
)

func New(cfg config.Iterator, itemType string) *iterator {
	if cfg.Separator == "" {
		cfg.Separator = defaultSeparator
	}

	return &iterator{
		Iterator: cfg,
		itemType: itemType,
	}
}

func (i *iterator) GetCollection(payload string) (items []dto.Item, err error) {
	r := csv.NewReader(strings.NewReader(payload))
	r.Comma = []rune(i.Separator)[0]
	values, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(values) < 1 {
		return nil, errors.New("empty collection")
	}

	identificatorColumn := getIdentificatorColumn(i.IdentificatorPath, values[0])
	if identificatorColumn == -1 {
		return nil, errors.New("incorrect identificator path")
	}

	raw := make(map[string]string)
	for _, row := range values[1:] {
		id, err := strconv.ParseInt(row[identificatorColumn], 10, 64)
		if err != nil {
			return nil, err
		}
		for idx, v := range row {
			raw[values[0][idx]] = v
		}

		payload, err := json.Marshal(raw)
		if err != nil {
			return nil, err
		}

		items = append(items, dto.Item{
			ID:      id,
			Plugin:  i.itemType,
			Payload: string(payload),
		})
	}

	return items, nil
}

func getIdentificatorColumn(columnName string, values []string) int {
	for idx, name := range values {
		if name == columnName {
			return idx
		}
	}
	return -1
}
