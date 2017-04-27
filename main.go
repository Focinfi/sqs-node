package main

import (
	"os"

	"github.com/Focinfi/sqs/node"
)

func main() {
	masterAddr := os.Getenv("SQS_MASTER_ADDR")
	if masterAddr == "" {
		panic("Need SQS_MASTER_ADDR environment parameters")
	}

	node.New(":8080", masterAddr).Start()
}
