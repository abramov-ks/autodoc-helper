package checker

import (
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"github.com/abramov-ks/autodoc-helper/pkg/db"
	"github.com/abramov-ks/autodoc-helper/pkg/telegram"
)

type Config struct {
	Autodoc     autodoc.AutodocConfig   `yaml:"autodoc"`
	Telegram    telegram.TelegramConfig `yaml:"telegram"`
	VersionFile string                  `yaml:"version_file"`
	DataBase    db.DatabaseConfig       `yaml:"database"`
}

type AppAction struct {
	Action string
	Value  string
}
