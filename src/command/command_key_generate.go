package command

import (
	"ellis.io/crypto/keys"
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
	if err != nil {
		println(err.Error())
		return false
	}

	return ok
}

func (args KeyGenerateArgs) Execute() (bool, error) {
	if !args.CanExecute() {
		return false, nil
	}

	param := args.Args[1:]
	fileName := strings.Replace(param[2], " ", "_", -1)
	e := keys.WriteKeys(fileName)
	if e != nil {
		return false, e
	}
	args.AfterExecute()
	return true, nil
}

func (args KeyGenerateArgs) AfterExecute() {
	println("The keys were created!")
}
