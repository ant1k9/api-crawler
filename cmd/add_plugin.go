package cmd

import (
	"errors"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/db"
	"github.com/ant1k9/api-crawler/internal/pkg/log"
	"github.com/spf13/cobra"
)

var addPlugin = &cobra.Command{
	Use:   "add-plugin",
	Short: "adds new partition to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.FatalIfErr(errors.New("provide only one argument to the command"))
		}

		db, err := db.New(config.Config.Database)
		log.FatalIfErr(err)
		log.FatalIfErr(db.CreatePartition(args[0]))
	},
}
