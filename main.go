package main

import (
	"os"

	"github.com/Focinfi/gosqs/node"
)

func main() {
	addr := os.Getenv("SQS_NODE_ADDR")
	masterAddr := os.Getenv("SQS_MASTER_ADDR")
	if addr == "" || masterAddr == "" {
		panic("Need SQS_NODE_ADDR and SQS_MASTER_ADDR environment parameters")
	}

	node.New(addr, 8081, masterAddr).Start()
}
