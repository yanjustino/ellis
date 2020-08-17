package command

import (
	"ellis.com/application"
	"ellis.com/data"
	"regexp"
	"strings"
)

type ListSecretsArgs struct {
	Args []string
}

func (args ListSecretsArgs) AfterExecute() {
	panic("implement me")
}

/* list -k {{jwt path}} */
func (args ListSecretsArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("(list\\s-k)(\\s[\\w+/.]*)", input)
	application.HandleError(err)

	return ok
}

func (args ListSecretsArgs) Execute() bool {
	param := args.Args[1:]
	data.FetchAllSecretKeys(param[2])
	return true
}
