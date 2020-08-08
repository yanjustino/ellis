package main

import (
	"ellis.io/command"
	"os"
)

func main() {
	command.KeyGenerateArgs{Args: os.Args}.Execute()
}
