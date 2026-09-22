package main

import (
	"fmt"
	"seckill/cmd/server"
	"uuid"
)

func main() {
	server.Cmd()
	fmt.Printf(uuid.New().String() + "\n")
}
