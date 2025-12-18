package main

import (
  "github.com/wailsapp/wails/v2/pkg/menu"
  "github.com/wailsapp/wails/v2/pkg/runtime"
)

func setupTray(app *App) *menu.TrayMenu {
  tray := menu.NewTrayMenu()
  tray.AddText("WS-MUX", nil)
  tray.AddText("显示窗口", func(_ *menu.CallbackData) {
    runtime.Show(app.ctx)
  })
  tray.AddText("退出", func(_ *menu.CallbackData) {
    runtime.Quit(app.ctx)
  })
  return tray
}
