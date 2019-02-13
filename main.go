package main

import (
  // "os"
  "fmt"
  "github.com/adlio/trello"
  "github.com/mitchellh/go-homedir"
  "github.com/pelletier/go-toml"
  // inspect "github.com/vigo/go-inspect"
)

func main() {
  home, _ := homedir.Dir()
  fmt.Println(home)
  config, _ := toml.LoadFile(home + "/.perica")
  app_key := config.Get("app_key").(string)
  token := config.Get("token").(string)
  // username := config.Get("username").(string)

  client := trello.NewClient(app_key, token)

  // member, _ := client.GetMember(username, trello.Defaults())

  // Print Trelli lists
  // board, _ := client.GetBoard("rGH9kMUZ", trello.Defaults())
  // lists, _ := board.GetLists(trello.Defaults())
  // for _, element := range(lists) {
  //   fmt.Println(inspect.Element(element))
  // }

  card := trello.Card{
    Name: "Card Name",
    Desc: "Card description",
  }
  list, _ := client.GetList("5c6450f088729a297d8dacfd", trello.Defaults())
  list.AddCard(&card, trello.Defaults())

}
