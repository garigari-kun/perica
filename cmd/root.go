package cmd

import (
  "github.com/spf13/cobra"
  "os"
  "log"
  "fmt"
  "strings"
  "net/url"
  "github.com/garigari-kun/perica/client"
  "github.com/adlio/trello"
)

func RootCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "perica",
    Short: "Issue to Trello card.",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
        log.Fatal("You need to parse issue url!")
      }
      input_url := args[0]
      u, err := url.Parse(input_url)
      if err != nil {
        log.Fatal(err)
      }
      if u.Hostname() == "trello.com" {
        parts := strings.Split(u.RequestURI(), "/")
        trello_client := client.NewTrelloClient()
        board, err := trello_client.Client.GetBoard(parts[2], trello.Defaults())
        if err != nil {
          log.Fatal("Board Not Found.")
        }
        lists, err := board.GetLists(trello.Defaults())
        if err != nil {
          log.Fatal("List not found.")
        }
        log.Print(lists)
        for _, list := range lists {
          fmt.Println("List ID: " + list.ID + "  " + "[" + list.Name + "]")
        }
        log.Fatal("Lists are above.")
      }
      github_client := client.NewGitHubClient(input_url)
      divided_url := client.NewUrlDivider(input_url)
      if divided_url.Type == "pull" {
        title, body := github_client.GetPrTitleAndBody(input_url)
        body = body + "\n" + "ref: " + input_url
        client := client.NewTrelloClient()
        client.CreateCard(title, body)
      } else if divided_url.Type == "issues" {
        title, body := github_client.GetIssueTitleAndBody(input_url)
        body = body + "\n" + "ref: " + input_url
        client := client.NewTrelloClient()
        client.CreateCard(title, body)
      }
    },
  }
  return cmd
}

func Execute() {
  cmd := RootCmd()
  if err := cmd.Execute(); err != nil {
    log.Print(err)
    os.Exit(1)
  }
}
