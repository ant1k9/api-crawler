package config

import (
	"io/ioutil"

	"github.com/ant1k9/api-crawler/internal/pkg/log"
	"gopkg.in/yaml.v3"
)

type (
	Yaml2Go struct {
		Crawlers []Crawler `yaml:"crawlers"`
		Database Database  `yaml:"database"`
	}

	Crawler struct {
		Name      string    `yaml:"name"`
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
		Start int    `yaml:"start"`
		End   int    `yaml:"end"`
		Type  string `yaml:"type"`
		Sleep Sleep  `yaml:"sleep"`
		Key   string `yaml:"key"`
	}

	Sleep struct {
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	}

	Iterator struct {
		Type              string `yaml:"type"`
		CollectionPath    string `yaml:"collection_path"`
		IdentificatorPath string `yaml:"identificator_path"`
	}

	Header struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	}
)

var Config Yaml2Go

func init() {
	config, err := ioutil.ReadFile("config.yml")
	log.FatalIfErr(err)
	log.FatalIfErr(yaml.Unmarshal(config, &Config))
}