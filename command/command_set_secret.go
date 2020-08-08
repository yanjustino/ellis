package command

type SetSecretArgs struct {
	Args []string
}

func (args SetSecretArgs) Execute() {
	println("Set Secrets")
}

func (args SetSecretArgs) CanExecute() bool {
	return false
}
