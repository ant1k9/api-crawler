package numeric

import (
	"strconv"
	"strings"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/errors"
)

const Type = "numeric"

type (
	paginator struct {
		config.Paginator
		next int
	}
)

func New(cfg config.Paginator) *paginator {
	return &paginator{
		Paginator: cfg,
		next:      cfg.Start,
	}
}

func (p *paginator) NextPage(payload string) (string, error) {
	if p.next > p.End {
		return "", errors.ErrPaginationEnd
	}
	defer func() { p.next++ }()
	return strings.Replace(payload, p.Key, strconv.Itoa(p.next), 1), nil
}
