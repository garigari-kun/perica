package main

import (
  // "fmt"
  // "log"
  "github.com/garigari-kun/perica/client"
  // inspect "github.com/vigo/go-inspect"
)

func main() {
  client := client.NewTrelloClient()
  fmt.Printf("%T\n", client)
  client.CreateCard()


}

// You can get app-key and token from here https://trello.com/app-key
