package main

import (
	"new-alunos/servers"
)

func main() {
	go servers.GrpcStart()
	servers.HttpStart()
}
