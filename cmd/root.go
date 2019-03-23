package cmd

import (
  "github.com/spf13/cobra"
  "os"
  "log"
  "github.com/garigari-kun/perica/client"
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
      github_client := client.NewGitHubClient(input_url)
      divided_url := client.NewUrlDivider(input_url)
      log.Fatal(*divided_url)
      title, body := github_client.GetIssueTitleAndBody(input_url)
      client := client.NewTrelloClient()
      client.CreateCard(title, body)
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
