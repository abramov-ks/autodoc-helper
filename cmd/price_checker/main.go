package main

import (
	"flag"
	"fmt"
	"github.com/abramov-ks/autodoc-helper/pkg/checker"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var APP_VERSION = "0.1"

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
	var cfgPath string
	var partNumber string
	var addPartNumber string
	var appVersion bool
	var checkAll bool
	var manufacterId string

	flag.StringVar(&cfgPath, "config", "./config.yml", "path to config file")
	flag.BoolVar(&appVersion, "version", false, "show application version")
	flag.StringVar(&partNumber, "check", "", "Partnumber to check")
	flag.StringVar(&manufacterId, "manufacter", "", "Manufacturer ID")
	flag.BoolVar(&checkAll, "check-all", false, "Partnumber to check")
	flag.StringVar(&addPartNumber, "add", "", "Partnumber to check")
	flag.Parse()
	if appVersion == true {
		fmt.Printf("Price Checker v%s\n", APP_VERSION)
		return
	}
	if err := ValidateConfigPath(cfgPath); err != nil {
		fmt.Println("No config file provided")
		return
	}

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var action = new(checker.AppAction)
	if partNumber != "" {
		action.Action = "check"
		action.Value = append(action.Value, partNumber, manufacterId)
	} else if addPartNumber != "" {
		action.Action = "add"
		action.Value = append(action.Value, addPartNumber, manufacterId)
	} else if checkAll == true {
		action.Action = "check-all"
	}

	cfg.Run(action)

}
