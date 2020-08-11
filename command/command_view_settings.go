package command

import (
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
	ok, err := regexp.MatchString("view\\s-k\\s[\\w](.|/)*", input)
	if err != nil {
		println(err.Error())
		return false
	}

	return ok
}

func (args ViewSecretsArgs) Execute() (bool, error) {
	if !args.CanExecute() {
		return false, nil
	}

	param := args.Args[1:]
	e := data.FetchSecretKeysHolder(param[2], false)
	if e != nil {
		println(e.Error())
		return false, e
	}
	return true, nil
}

func (args ViewSecretsArgs) AfterExecute() {
	panic("implement me")
}
