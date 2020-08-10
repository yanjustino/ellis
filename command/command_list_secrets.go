package command

import (
	"ellis.com/data"
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
	if err != nil {
		println(err.Error())
		return false
	}

	return ok
}

func (args ListSecretsArgs) Execute() (bool, error) {
	if !args.CanExecute() {
		return false, nil
	}

	param := args.Args[1:]
	e := data.FetchAllSecretKeys(param[2])
	if e != nil {
		println(e.Error())
		return false, e
	}
	return true, nil
}
