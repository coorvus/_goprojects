package main

import (
	"coorvus/scheduler/cli"
	"os"
)

func main() {
	args := os.Args[1:]
	cli.Run(args)
}
