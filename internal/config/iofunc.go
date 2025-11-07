package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read() (*Config, error) {
	var cfg Config

	path, err := GetConfigFilePath()
	if err != nil {
		return &cfg, err
	}

	f, err := os.Open(path)
	if err != nil {
		return &cfg, err
	}

	defer f.Close()

	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return &cfg, err
	}
	return &cfg, nil

}

func Write(cfg Config) error {
	path, err := GetConfigFilePath()

	//debug
	fmt.Println(path)
	fmt.Println(cfg)

	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	return Write(*c)
}

func GetConfigFilePath() (string, error) {
	wd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, CONFIGFILENAME)
	return path, nil
}
