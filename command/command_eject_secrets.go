package command

import (
	"ellis.com/data"
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
	ok, err := regexp.MatchString("(eject\\s-k)(\\s[\\w+/.]*)", input)
	if err != nil {
		println(err.Error())
		return false
	}

	return ok && len(param) == 3
}

func (args EjectSecretsArgs) Execute() bool {
	param := args.Args[1:]
	data.FetchSecretKeysHolder(param[2], true)
	args.AfterExecute()
	return true
}

func (args EjectSecretsArgs) AfterExecute() {
	println("The keys were ejected!")
}
