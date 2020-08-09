package command

import (
	"ellis.io/data"
	"ellis.io/utils"
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
	utils.CheckError(err)

	return ok && len(param) == 5
}

func (args SetSecretArgs) Execute() {
	if args.CanExecute() {
		param := args.Args[1:]
		data.StoreSecretKey(param[2], param[3], param[4])
		args.AfterExecute()
	}
}

func (args SetSecretArgs) AfterExecute()  {
	param := args.Args[1:]
	println(fmt.Sprintf("The key %v was registered in %v!", param[3], param[2]))
}
