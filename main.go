package main

import (
	"github.com/hnamzian/goblock/internal/config"
	"github.com/hnamzian/goblock/internal/node"
)

func main() {
	cfg, _ := config.GetConfigs()

	nodeCfg := &node.Config{
		Addr: cfg.Addr,
	}
	n := node.New(nodeCfg)

	staticNodes := []string{
		"localhost:8080",
		"localhost:8081",
	}

	n.Start(staticNodes)
}
