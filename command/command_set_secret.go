package command

import (
	"ellis.com/data"
	"fmt"
	"regexp"
	"strings"
)

type SetSecretArgs struct {
	Args []string
}

/* set -k {{file}} {{key}} {{value}}*/
func (args SetSecretArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("set\\s-k\\s[\\w](.|\\/)*", input)
	if err != nil {
		println(err.Error())
		return false
	}

	return ok && len(param) == 5
}

func (args SetSecretArgs) Execute() (bool, error) {
	if !args.CanExecute() {
		return false, nil
	}

	param := args.Args[1:]
	e := data.StoreSecretKey(param[2], param[3], param[4])
	if e != nil {
		println(e.Error())
		return false, e
	}
	args.AfterExecute()
	return true, nil
}

func (args SetSecretArgs) AfterExecute() {
	param := args.Args[1:]
	println(fmt.Sprintf("The key %v was registered in %v!", param[3], param[2]))
}
