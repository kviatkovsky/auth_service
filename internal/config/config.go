package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string        `yaml:"env" env-default:"local"`
	MySQL 	StorageConfig `yaml:"mysql"`
	Service Service       `yaml:"service" env-default:"8090"`
}

type Service struct {
	Port  string
	Host  string `yaml:"host" env-default:"0.0.0.0"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Attempts int    `yaml:"attempts"`
}

func MustLoad() *Config {
	args := fetchArgs()
	if args["config_path"] == "" {
		panic("config path is empty")
	}

	if args["service_port"] == "" {
		panic("config path is empty")
	}

	return MustLoadPath(args["config_path"], args["service_port"])
}

func MustLoadPath(configPath, servicePort string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	cfg.Service.Port = servicePort

	return &cfg
}

func fetchArgs() map[string]string {
	var configPath string
	var servicePort string


	flag.StringVar(&configPath, "config", "", "config path")
	flag.StringVar(&servicePort, "port", "", "config path")
	flag.Parse()

	return map[string]string {
		"config_path": configPath,
		"service_port": servicePort,
	}
}
