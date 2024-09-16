package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port int
	}
	Paths struct {
		XSDPath   string
		XMLInput  string
		JSONInput string
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, err
	}
	cfg.Server.Port = port

	cfg.Paths.XSDPath = getEnv("XSD_PATH", "./schemas/ACCC470.xsd")
	cfg.Paths.XMLInput = getEnv("XML_INPUT_PATH", "./output.xml")
	cfg.Paths.JSONInput = getEnv("JSON_INPUT_PATH", "./output.json")

	return cfg, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
