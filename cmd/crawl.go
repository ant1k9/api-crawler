package cmd

import (
	"errors"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/crawler"
	"github.com/ant1k9/api-crawler/internal/pkg/db"
	"github.com/ant1k9/api-crawler/internal/pkg/log"
	"github.com/spf13/cobra"
)

var crawl = &cobra.Command{
	Use:   "crawl",
	Short: "crawls the site with given crawler",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.FatalIfErr(errors.New("provide only one argument to the command"))
		}

		var crawlerConfig config.Crawler
		for _, cfg := range cfg.Crawlers {
			if cfg.Type == args[0] {
				crawlerConfig = cfg
				break
			}
		}

		db, err := db.New(cfg.Database)
		log.FatalIfErr(err)
		cr, err := crawler.New(crawlerConfig, db)
		log.FatalIfErr(err)
		log.FatalIfErr(cr.Crawl())
	},
}
