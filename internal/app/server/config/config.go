package config

import (
	yaml "gopkg.in/yaml.v2"
)

type configFile struct {
	Postgresql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"postgresql"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := configFile{}

	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}
	c.DB.Host = cf.Postgresql.Host
	c.DB.Port = cf.Postgresql.Port
	c.DB.User = cf.Postgresql.User
	c.DB.Password = cf.Postgresql.Password
	c.DB.DBName = cf.Postgresql.DBName
	c.DB.Sslmode = cf.Postgresql.Sslmode
	return &c, nil
}
