package command

import (
	"ellis.io/data"
)

type ViewSecretsArgs struct {
	Args []string
}

func (args ViewSecretsArgs) CanExecute() bool {
	param := args.Args[1:]
	return len(param) == 3 && param[0] == "view" && param[1] == "-k" && param[2] != ""
}

func (args ViewSecretsArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]

		data.FetchSecretKeysHolder(param[2])
	}
}
