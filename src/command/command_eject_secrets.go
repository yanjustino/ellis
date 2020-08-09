package command


import (
	"ellis.io/data"
	"ellis.io/utils"
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
	utils.CheckError(err)

	return ok && len(param) == 3
}

func (args EjectSecretsArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]
		data.FetchSecretKeysHolder(param[2], true)
		args.AfterExecute()
	}
}

func (args EjectSecretsArgs) AfterExecute()  {
	println("The keys were ejected!")
}
