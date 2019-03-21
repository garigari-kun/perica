package scraper

import (
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)

type Scraper struct {
  Url string
}

func NewScraper(url string) *Scraper {
  return &Scraper{Url: url}
}

func (self *Scraper) GetTitleFromIssue() string {
    res, err := http.Get(self.Url)
    if err != nil {
      log.Fatal(err)
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
      log.Fatal("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
      log.Fatal(err)
    }

    selection := doc.Find(".gh-header-title")
    title := selection.Find(".js-issue-title").Text()
    return title
}
