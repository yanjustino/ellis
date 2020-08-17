package command

import (
	"ellis.com/application"
	"ellis.com/data"
	"fmt"
	"regexp"
	"strings"
)

type SetSecretArgs struct {
	Args []string
}

/* set -k {{jwt path}} {{key}} {{value}}*/
func (args SetSecretArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("(set\\s-k)(\\s[\\w+/.]*)(\\s[\\w|\".*?\"]*)\\s((\\w+)|(\".*?\"))", input)
	application.HandleError(err)

	return ok && len(param) == 5
}

func (args SetSecretArgs) Execute() bool {
	param := args.Args[1:]
	data.StoreSecretKey(param[2], param[3], param[4])
	args.AfterExecute()
	return true
}

func (args SetSecretArgs) AfterExecute() {
	param := args.Args[1:]
	println(fmt.Sprintf("The key %v was registered in %v!", param[3], param[2]))
}
