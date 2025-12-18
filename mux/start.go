package mux

import "net"

type Config struct {
  Server string
  Token string
  Secret string
  Listen string
}

func Start(cfg Config) error {
  _, err := net.Listen("tcp", cfg.Listen)
  return err
}
