package checker

import (
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"github.com/abramov-ks/autodoc-helper/pkg/db"
)

type TelegramConfig struct {
	Token  string `yaml:"token"`
	ChatId int    `yaml:"chat_id"`
}

type Config struct {
	Autodoc     autodoc.AutodocConfig `yaml:"autodoc"`
	Telegram    TelegramConfig        `yaml:"telegram"`
	VersionFile string                `yaml:"version_file"`
	DataBase    db.DatabaseConfig     `yaml:"database"`
}

type AppAction struct {
	Action string
	Value  string
}
