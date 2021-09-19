package db

import (
	"fmt"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	migrateSQL    = `CREATE TABLE data_%[1]s PARTITION OF data FOR VALUES IN ('%[1]s')`
	insertItemSQL = `
		INSERT INTO data (id, plugin, payload)
		VALUES ($1, $2, $3)
		ON CONFLICT (id, plugin) DO UPDATE
			SET payload = excluded.payload`
)

type (
	store struct {
		conn *sqlx.DB
	}

	Store interface {
		CreatePartition(name string) error
		InsertItem(item dto.Item) error
	}
)

func New(cfg config.Database) (Store, error) {
	conn, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Sslmode,
		),
	)
	return &store{conn: conn}, err
}

func (s *store) CreatePartition(name string) error {
	_, err := s.conn.Exec(fmt.Sprintf(migrateSQL, name))
	return err
}

func (s *store) InsertItem(item dto.Item) error {
	_, err := s.conn.Exec(insertItemSQL, item.ID, item.Plugin, item.Payload)
	return err
}
