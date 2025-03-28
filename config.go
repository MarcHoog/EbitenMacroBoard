package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type KeyConfig struct {
	Label     string `yaml:"label"`
	RuneValue string `yaml:"runeValue"`
}

type Config struct {
	KeyBoardOptions `yaml:"KeyBoardOptions"`
	Keys            []KeyConfig `yaml:"Keys"`
}

const DefaultConfigFileName = ".macroboard_config.yaml"

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := path + "\\" + DefaultConfigFileName
	return configFilePath, err
}

func NewConfigFromFile(path string) (*Config, error) {

	var config Config

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Unmarshal the YAML data into the Config struct
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	if len(config.Keys) == 0 {
		config.Keys = append(config.Keys, KeyConfig{
			Label:     "A",
			RuneValue: "A",
		})
	}

	return &config, nil
}
