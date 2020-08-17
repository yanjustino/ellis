package command

import (
	"ellis.com/application"
	"ellis.com/data"
	"fmt"
	"regexp"
	"strings"
)

type GetSecretArgs struct {
	Args []string
}

/* get -k {{jwt path}} "{{key}}" */
func (args GetSecretArgs) CanExecute() bool {
	param := args.Args[1:]

	input := strings.Join(param, " ")
	ok, err := regexp.MatchString("(get\\s-k)(\\s[\\w+/.|\\w+\\.]*)(\\s[\\w|\".*?\"]*)", input)
	application.HandleError(err)

	return ok && len(param) == 4
}

func (args GetSecretArgs) Execute() bool {
	param := args.Args[1:]
	ok, secret := data.FetchDecryptedSecret(param[2], param[3])
	if ok {
		println(secret.ToJson())
		return true
	}

	return false
}

func (args GetSecretArgs) AfterExecute() {
	println(fmt.Sprintf("No Implemented"))
}
