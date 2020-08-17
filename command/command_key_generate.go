package command

import (
	"ellis.com/application"
	"ellis.com/crypto/keys"
	"regexp"
	"strings"
)

type KeyGenerateArgs struct {
	Args []string
}

/* keys -g {{id}} */
func (args KeyGenerateArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("(keys\\s-g)(\\s[\\w+/.|\\w+\\.]*)", input)
	application.HandleError(err)

	return ok
}

func (args KeyGenerateArgs) Execute() bool {

	param := args.Args[1:]
	fileName := strings.Replace(param[2], " ", "_", -1)
	keys.WriteKeys(fileName)
	args.AfterExecute()
	return true
}

func (args KeyGenerateArgs) AfterExecute() {
	println("The keys were created!")
}
