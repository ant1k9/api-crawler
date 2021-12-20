package config

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ant1k9/api-crawler/internal/pkg/log"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Crawlers []Crawler `yaml:"crawlers"`
		Database Database  `yaml:"database"`
	}

	Crawler struct {
		Name      string    `yaml:"name"`
		OnError   string    `yaml:"on_error"`
		Type      string    `yaml:"type"`
		Method    string    `yaml:"method"`
		Link      string    `yaml:"link"`
		Payload   string    `yaml:"payload"`
		Paginator Paginator `yaml:"paginator"`
		Iterator  Iterator  `yaml:"iterator"`
		Headers   []Header  `yaml:"headers"`
	}

	Database struct {
		Sslmode  string `yaml:"sslmode"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	Paginator struct {
		Start int      `yaml:"start"`
		End   int      `yaml:"end"`
		Type  string   `yaml:"type"`
		Sleep Sleep    `yaml:"sleep"`
		Key   string   `yaml:"key"`
		Items []string `yaml:"items"`
	}

	Sleep struct {
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	}

	Iterator struct {
		Type              string `yaml:"type"`
		Regexp            string `yaml:"regex"`
		CollectionPath    string `yaml:"collection_path"`
		IdentificatorPath string `yaml:"identificator_path"`
		Separator         string `yaml:"separator"`
	}

	Header struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}
)

func (c *Crawler) GetPaginatorOrigin() string {
	switch strings.ToUpper(c.Method) {
	case http.MethodPost:
		return c.Payload
	default:
		return c.Link
	}
}

func Init() (cfg Config) {
	config, err := ioutil.ReadFile("config.yml")
	log.FatalIfErr(err)
	log.FatalIfErr(yaml.Unmarshal(config, &cfg))
	return cfg
}
