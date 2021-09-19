package crawler

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/db"
	errs "github.com/ant1k9/api-crawler/internal/pkg/errors"
	iterator "github.com/ant1k9/api-crawler/internal/pkg/iterators"
	"github.com/ant1k9/api-crawler/internal/pkg/log"
	"github.com/ant1k9/api-crawler/internal/pkg/paginator"
)

type (
	crawler struct {
		cfg       config.Crawler
		iterator  iterator.Iterator
		paginator paginator.Paginator
		store     db.Store

		minSleepDuration, maxSleepDuration time.Duration
	}

	Crawler interface {
		Crawl() error
	}
)

func New(cfg config.Crawler, store db.Store) (Crawler, error) {
	it, err := iterator.New(cfg.Iterator, cfg.Type)
	if err != nil {
		return nil, err
	}

	pager, err := paginator.New(cfg.Paginator)
	if err != nil {
		return nil, err
	}

	minSleepDuration, err := time.ParseDuration(cfg.Paginator.Sleep.Min)
	if err != nil {
		return nil, err
	}

	maxSleepDuration, err := time.ParseDuration(cfg.Paginator.Sleep.Max)
	if err != nil {
		return nil, err
	}

	return &crawler{
		cfg:              cfg,
		iterator:         it,
		paginator:        pager,
		store:            store,
		minSleepDuration: minSleepDuration,
		maxSleepDuration: maxSleepDuration,
	}, nil
}

func (c *crawler) Crawl() error {
	getNextPage := c.paginator.NextPage
	for next, err := getNextPage(c.cfg.Payload); ; next, err = getNextPage(c.cfg.Payload) {
		if err != nil {
			return handlePaginatorErr(err)
		}

		log.Info("start crawling " + next)
		data, err := c.getData(next)
		if err != nil {
			return err
		}

		items, err := c.iterator.GetCollection(string(data))
		if err != nil {
			return err
		}

		for _, item := range items {
			err = c.store.InsertItem(item)
			if err != nil {
				return err
			}
		}
		c.sleep()
	}
}

func (c *crawler) getData(next string) ([]byte, error) {
	switch c.cfg.Method {
	case "post":
		reader := strings.NewReader(next)
		req, err := http.NewRequest(http.MethodPost, c.cfg.Link, reader)
		if err != nil {
			return nil, err
		}
		for _, header := range c.cfg.Headers {
			req.Header.Add(header.Key, header.Value)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(resp.Body)

	default:
		return nil, errors.New("TODO: add get method, currently not implemented")
	}
}

func (c *crawler) sleep() {
	jitter := int64(c.maxSleepDuration - c.minSleepDuration)
	time.Sleep(time.Duration(rand.Int63n(jitter)) + c.minSleepDuration)
}

func handlePaginatorErr(err error) error {
	if errors.Is(err, errs.ErrPaginationEnd) {
		return nil
	}
	return err
}
