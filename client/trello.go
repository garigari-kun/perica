package client

import (
  "log"
  "github.com/adlio/trello"
  "github.com/garigari-kun/perica/config"
)

type TrelloClient struct {
  Config *config.TrelloConfig
  Client *trello.Client
}

func NewTrelloClient() *TrelloClient {
  config := config.NewTrelloConfig()
  client := trello.NewClient(config.AppKey, config.Token)
  return &TrelloClient{Config: config, Client: client}
}

func (t *TrelloClient) CreateCard(title string, desc string) {
  list, err := t.Client.GetList(t.Config.ListId, trello.Defaults())
  if err != nil {
    log.Print(err)
  }

  err = list.AddCard(&trello.Card{ Name: title, Desc: desc }, trello.Defaults())
  if err != nil {
    log.Print(err)
    log.Print("Card could not be created.")
  } else {
    log.Print("Card successfully created.")
  }
}
