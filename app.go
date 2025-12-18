package main

import (
  "context"
  "ws-mux-gui/mux"
)

type App struct {
  ctx context.Context
}

func NewApp() *App { return &App{} }

func (a *App) Startup(ctx context.Context) {
  a.ctx = ctx
}

func (a *App) Start(server, token, secret, listen string) string {
  err := mux.Start(mux.Config{
    Server: server,
    Token: token,
    Secret: secret,
    Listen: listen,
  })
  if err != nil {
    return err.Error()
  }
  return "启动成功"
}
