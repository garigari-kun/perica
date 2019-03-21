package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"log"
)

type TrelloConfig struct {
	AppKey string
	Token  string
	ListId string
}

func NewTrelloConfig() *TrelloConfig {
	home, err := homedir.Dir()
	if err != nil {
		log.Print(err)
	}
	config, err := toml.LoadFile(home + "/.perica")
	if err != nil {
		log.Print(err)
	}
	app_key := config.Get("app_key").(string)
	token := config.Get("token").(string)
	list_id := config.Get("list_id").(string)

	return &TrelloConfig{AppKey: app_key, Token: token, ListId: list_id}
}
