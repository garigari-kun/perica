package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"log"
)

type GithubConfig struct {
	Token string
}

func NewGithubConfig() *GithubConfig {
	home, err := homedir.Dir()
	if err != nil {
		log.Print(err)
	}
	config, err := toml.LoadFile(home + "/.perica")
	if err != nil {
		log.Print(err)
	}
	token := config.Get("github_token").(string)

	return &GithubConfig{Token: token}
}
