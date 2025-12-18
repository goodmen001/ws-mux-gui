package main

import (
  "github.com/wailsapp/wails/v2"
  "github.com/wailsapp/wails/v2/pkg/options"
)

func main() {
  app := NewApp()
  wails.Run(&options.App{
    Title: "WS-MUX GUI",
    Width: 420,
    Height: 300,
    Bind: []interface{}{app},
  })
}
