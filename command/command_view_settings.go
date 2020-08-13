package command

import (
	"ellis.com/application"
	"ellis.com/data"
	"regexp"
	"strings"
)

type ViewSecretsArgs struct {
	Args []string
}

func (args ViewSecretsArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("(view\\s-k)(\\s[\\w+/.]*)", input)
	application.HandleError(err)

	return ok
}

func (args ViewSecretsArgs) Execute() bool {
	param := args.Args[1:]
	data.FetchSecretKeysHolder(param[2], false)
	return true
}

func (args ViewSecretsArgs) AfterExecute() {
	panic("implement me")
}
