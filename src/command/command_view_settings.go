package command

import (
	"ellis.io/data"
	"ellis.io/utils"
	"regexp"
	"strings"
)

type ViewSecretsArgs struct {
	Args []string
}

func (args ViewSecretsArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("view\\s-k\\s[\\w](.|\\/)*", input)
	utils.CheckError(err)

	return ok
}

func (args ViewSecretsArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]
		data.FetchSecretKeysHolder(param[2], false)
	}
}
