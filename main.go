package main

import (
  "fmt"
  // "log"
  "github.com/garigari-kun/perica/client"
  // "github.com/garigari-kun/perica/scraper"
  // inspect "github.com/vigo/go-inspect"
)

func main() {
  // scr := scraper.NewScraper("https://github.com/garigari-kun/shutto/issues/5")
  // title := scr.GetTitleFromIssue()
  // client := client.NewTrelloClient()
  // desc := "ref " + scr.Url
  // client.CreateCard(title, desc)
  // client.NewGitHubClient()
  github_client := client.NewGitHubClient("https://github.com/garigari-kun/shutto/issues/5")
  title, body := github_client.GetIssueTitleAndBody("https://github.com/garigari-kun/shutto/issues/5")
  fmt.Println(title)
  fmt.Println(body)
}

// You can get app-key and token from here https://trello.com/app-key
