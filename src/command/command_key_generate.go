package command

import (
	"ellis.io/crypto/keys"
	"ellis.io/utils"
	"regexp"
	"strings"
)

type KeyGenerateArgs struct {
	Args []string
}

/* keys -g {{identification}} */
func (args KeyGenerateArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("keys\\s-g\\s[\\w](.|\\/)*", input)
	utils.CheckError(err)

	return ok
}

func (args KeyGenerateArgs) Execute() {
	param := args.Args[1:]
	if args.CanExecute() {
		fileName := strings.Replace(param[2], " ", "_", -1)
		keys.WriteKeys(fileName)
		args.AfterExecute()
	}
}

func (args KeyGenerateArgs) AfterExecute() {
	println("The keys were created!")
}
