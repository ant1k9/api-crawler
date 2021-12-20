package list

import (
	"strings"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/errors"
)

const Type = "list"

type (
	paginator struct {
		config.Paginator
		next    int
		current string
	}
)

func New(cfg config.Paginator) *paginator {
	return &paginator{
		Paginator: cfg,
		next:      0,
	}
}

func (p *paginator) NextPage(payload string) (string, error) {
	if p.next >= len(p.Items) {
		return "", errors.ErrPaginationEnd
	}
	defer func() { p.current = p.Items[p.next]; p.next++ }()
	return strings.ReplaceAll(payload, p.Key, p.Items[p.next]), nil
}

func (p *paginator) AdjustPlugin(plugin string) string {
	return plugin + "_" + p.current
}
