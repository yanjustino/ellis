package command

type Command interface {
	Execute() bool
	CanExecute() bool
	AfterExecute()
}

func Builder(args []string) Command {
	commands := []Command{
		KeyGenerateArgs{args},
		ListSecretsArgs{args},
		SetSecretArgs{args},
		GetSecretArgs{args},
		ViewSecretsArgs{args},
		EjectSecretsArgs{args},
		HelpArgs{args},
	}

	for _, cmd := range commands {
		if cmd.CanExecute() {
			return cmd
		}
	}

	return HelpArgs{args}
}
