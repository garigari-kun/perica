package main

import (
  // "os"
  "fmt"

  "github.com/mitchellh/go-homedir"
  "github.com/pelletier/go-toml"
)

func main() {
  home, _ := homedir.Dir()
  fmt.Println(home)
  config, _ := toml.LoadFile(home + "/.perica")
  fmt.Println(config.Get("username").(string))
}
