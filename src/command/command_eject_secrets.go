package command

import (
	"ellis.io/data"
	"regexp"
	"strings"
)

type EjectSecretsArgs struct {
	Args []string
}

/* eject -k {{file}} */
func (args EjectSecretsArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("eject\\s-k\\s[\\w](.|\\/)*", input)
	if err != nil {
		println(err.Error())
		return false
	}

	return ok && len(param) == 3
}

func (args EjectSecretsArgs) Execute() (bool, error) {
	if !args.CanExecute() {
		return false, nil
	}

	param := args.Args[1:]
	e := data.FetchSecretKeysHolder(param[2], true)
	if e != nil {
		println(e.Error())
		return false, e
	}
	args.AfterExecute()
	return true, nil
}

func (args EjectSecretsArgs) AfterExecute() {
	println("The keys were ejected!")
}
