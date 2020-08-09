package command

import (
	"ellis.io/data"
	"ellis.io/utils"
	"regexp"
	"strings"
)

type ListSecretsArgs struct {
	Args []string
}

/* list -k {{file}} */
func (args ListSecretsArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("list\\s-k\\s[\\w](.|\\/)*", input)
	utils.CheckError(err)

	return ok
}

func (args ListSecretsArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]
		data.FetchAllSecretKeys(param[2])
	}
}
