package command

import (
	"ellis.io/data"
)

type SetSecretArgs struct {
	Args []string
}

func (args SetSecretArgs) CanExecute() bool {
	param := args.Args[1:]
	return len(param) == 5 && param[0] == "set" && param[1] == "-k" && param[2] != "" && param[3] != "" && param[4] != ""
}

func (args SetSecretArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]

		data.StoreSecretKey(param[2], param[3], param[4])
	}
}
