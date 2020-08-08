package command

import (
	keys "ellis.io/crypto/keys"
	"strings"
)

type KeyGenerateArgs struct {
	Args []string
}

func (args KeyGenerateArgs) CanExecute() bool {
	param := args.Args[1:]
	return len(param) == 3 && param[0] == "keys" && param[1] == "-g" && param[2] != ""
}

func (args KeyGenerateArgs) Execute() {
	param := args.Args[1:]
	if args.CanExecute() {
		fileName := strings.Replace(param[2], " ", "_", -1)
		keys.WriteKeys(fileName)
		args.AfterExecute()
	}
}

func (args KeyGenerateArgs) AfterExecute() {
	println("The files were created!")
}
