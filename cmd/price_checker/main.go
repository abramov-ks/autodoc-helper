package main

import (
	"fmt"
	"github.com/abramov-ks/autodoc-helper/pkg/checker"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Options struct {
	Config         string `long:"config" short:"c" default:"./config.yml" description:"Set config path"`
	CleanupAction  bool   `long:"cleanup" description:"Do system cleanup"`
	CheckAllAction bool   `long:"check-all" description:"Check all listing spare parts"`
	AddAction      bool   `long:"add" description:"Check all listing spare parts"`
	CheckAction    bool   `long:"check" description:"Check spare part by ID"`
	Version        bool   `long:"version" short:"v" description:"Get app version"`
	PartnumberId   string `long:"partnumber" short:"p" description:"ID of part number"`
	ManufacterId   string `long:"manufacter" short:"m" description:"ID of manufacter"`
}

var AppVersion = "1.1"

// ValidateConfigPath Валидация конфига
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

// NewConfig Загрузка конфига
func NewConfig(configPath string) (*checker.Config, error) {
	config := &checker.Config{}
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

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	if opts.Version == true {
		fmt.Printf("Price Checker v%s\n", AppVersion)
		os.Exit(0)
	}

	if err := ValidateConfigPath(opts.Config); err != nil {
		fmt.Println("No valid config file provided")
		os.Exit(0)
	}

	cfg, err := NewConfig(opts.Config)
	if err != nil {
		log.Fatal(err)
	}

	var action = new(checker.AppAction)
	if opts.CheckAction == true {
		action.Action = "check"
		action.Value = append(action.Value, opts.PartnumberId, opts.ManufacterId)
	} else if opts.AddAction == true {
		action.Action = "add"
		action.Value = append(action.Value, opts.PartnumberId, opts.ManufacterId)
	} else if opts.CheckAllAction == true {
		action.Action = "check-all"
	} else if opts.CleanupAction == true {
		action.Action = "cleanup"
	}

	os.Exit(cfg.Run(action))

}
