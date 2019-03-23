package client

import (
  "fmt"
  "log"
  "golang.org/x/oauth2"
  "github.com/google/go-github/github"
  "github.com/garigari-kun/perica/config"
  "context"
  "net/url"
  "strings"
  "strconv"
)

type UrlDivider struct {
  Domain string
  Org string
  Repo string
  Type string
  Id int
}

type GitHubClient struct {
  Client *github.Client
}

func NewUrlDivider(input_url string) *UrlDivider {
  u, err := url.Parse(input_url)
  if err != nil {
    log.Fatal(err)
  }
  parts := strings.Split(u.RequestURI(), "/")
  i, _ := strconv.Atoi(parts[4])
  return &UrlDivider{Domain: u.Hostname(), Org: parts[1], Repo: parts[2], Type: parts[3], Id: i}
}

func NewGitHubClient(input_url string) *GitHubClient {
  divided_url := NewUrlDivider(input_url)
  config := config.NewGithubConfig()
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token})
  tc := oauth2.NewClient(oauth2.NoContext,ts)
  client := github.NewClient(tc)
  if divided_url.Domain != "github.com" {
    enterpriseURL, err := url.Parse("https://" + divided_url.Domain + "/api/v3/")
    if err != nil {
      fmt.Println(err)
    }
    client.BaseURL = enterpriseURL
  }
  return &GitHubClient{Client: client}
}

func (self *GitHubClient) GetIssueTitleAndBody(input_url string) (string, string) {
  divided_url := NewUrlDivider(input_url)
  issue, _, err := self.Client.Issues.Get(context.Background(), divided_url.Org, divided_url.Repo, divided_url.Id)
  if err != nil {
    log.Print(nil)
  }

  return *issue.Title, *issue.Body
}

func (self *GitHubClient) GetPrTitleAndBody(input_url string) (string, string) {
  divided_url := NewUrlDivider(input_url)
  pr, _, err := self.Client.PullRequests.Get(context.Background(), divided_url.Org, divided_url.Repo, divided_url.Id)
  if err != nil {
    log.Print(nil)
  }
  return *pr.Title, *pr.Body
}
