package command

import (
	"ellis.io/data"
)

type ListSecretsArgs struct {
	Args []string
}

func (args ListSecretsArgs) CanExecute() bool {
	param := args.Args[1:]
	return len(param) == 3 && param[0] == "list" && param[1] == "-k" && param[2] != ""
}

func (args ListSecretsArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]

		data.FetchAllSecretKeys(param[2])
	}
}
