package main

import (
  // "fmt"
  // "log"
  "github.com/garigari-kun/perica/client"
  "github.com/garigari-kun/perica/scraper"
  // inspect "github.com/vigo/go-inspect"
)

func main() {
  scr := scraper.NewScraper("https://github.com/garigari-kun/shutto/issues/5")
  title := scr.GetTitleFromIssue()
  client := client.NewTrelloClient()
  desc := "ref " + scr.Url
  client.CreateCard(title, desc)
}

// You can get app-key and token from here https://trello.com/app-key
