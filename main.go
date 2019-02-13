package main

import (
  // "os"
  "fmt"
  "github.com/adlio/trello"
  "github.com/mitchellh/go-homedir"
  "github.com/pelletier/go-toml"
)

func main() {
  home, _ := homedir.Dir()
  fmt.Println(home)
  config, _ := toml.LoadFile(home + "/.perica")
  app_key := config.Get("app_key").(string)
  token := config.Get("token").(string)
  username := config.Get("username").(string)

  client := trello.NewClient(app_key, token)

  member, _ := client.GetMember(username, trello.Defaults())
  boards, _ := member.GetBoards(trello.Defaults())

}
