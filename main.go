package main

import (
  "github.com/garigari-kun/perica/client"
  // inspect "github.com/vigo/go-inspect"
  "github.com/garigari-kun/perica/cmd"
)

func main() {
  cmd.Execute()
}

func _main() {
  github_client := client.NewGitHubClient("https://github.com/garigari-kun/shutto/issues/5")
  title, body := github_client.GetIssueTitleAndBody("https://github.com/garigari-kun/shutto/issues/5")
  client := client.NewTrelloClient()
  client.CreateCard(title, body)
}

// You can get app-key and token from here https://trello.com/app-key
