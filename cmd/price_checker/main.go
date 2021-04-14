package main

import (
	"flag"
	"fmt"
	"github.com/abramov-ks/autodoc-helper/pkg/autodoc"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type TelegramConfig struct {
	Token  string `yaml:"token"`
	ChatId int    `yaml:"chat_id"`
}

type Config struct {
	Autodoc     autodoc.AutodocConfig `yaml:"autodoc"`
	Telegram    TelegramConfig        `yaml:"telegram"`
	VersionFile string                `yaml:"version_file"`
}

var APP_VERSION = "0.1"

/** Валидация конфига */
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}

	return nil
}

/** Загрузка конфига */
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

/** Запуск */
func (config Config) Run(partnumber string) {
	log.Println("run app for user", config.Autodoc.Username)
	var autodocSession autodoc.AutodocSession
	autodocSession.FillFromConfig(&config.Autodoc)
	ok := autodocSession.Auth()
	if !ok {
		log.Println("Cannot create autodoc session")
	}

	if partnumber == "" {
		log.Println("No partnumber to check")
		return
	}

	partNumberInfo, partNumberInfoErr := autodocSession.CheckPartnumber(partnumber)
	if partNumberInfoErr != nil {
		log.Println("Cannot check partnumber price: %s", partNumberInfoErr)
		return
	}

	log.Printf("Деталь: %s цена: %.2f руб.", partNumberInfo.PartNumber, partNumberInfo.MinimalPrice)
}

func main() {
	var cfgPath string
	var partNumber string
	var appVersion bool
	flag.StringVar(&cfgPath, "config", "./config.yml", "path to config file")
	flag.BoolVar(&appVersion, "version", false, "show application version")
	flag.StringVar(&partNumber, "partnumber", "", "Partnumber to check")
	flag.Parse()
	if appVersion == true {
		fmt.Printf("Price Checker v%s\n", APP_VERSION)
		return
	}
	if err := ValidateConfigPath(cfgPath); err != nil {
		fmt.Println("No config file provieded")
		return
	}

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Run(partNumber)

}
