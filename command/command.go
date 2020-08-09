package command

type Command interface {
	Execute() (bool, error)
	CanExecute() bool
	AfterExecute()
}
