package command

type Command interface {
	Execute()
	CanExecute() bool
	AfterExecute()
}
