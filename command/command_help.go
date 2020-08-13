package command

type HelpArgs struct {
	Args []string
}

/* list -k {{file}} */
func (args HelpArgs) CanExecute() bool {
	param := args.Args[1:]

	return len(args.Args) == 1 ||
		param[0] == "help" || param[0] == "-h" || param[0] == "--h"
}

func (args HelpArgs) Execute() bool {
	println(AsciiBanner())
	println(Helptext())
	return true
}

func AsciiBanner() string {
	return " ______ _      _      _____  _____\n" +
		"|  ____| |    | |    |_   _|/ ____|\n" +
		"| |__  | |    | |      | | | (___\n" +
		"|  __| | |    | |      | |  \\___ \\\n" +
		"| |____| |____| |____ _| |_ ____) |\n" +
		"|______|______|______|_____|_____/\n\n"
}

func Helptext() string {
	return " Usage: ellis <command> [options] [path-to-jwk]\n" +
		" Usage: ellis [path-to-jwk]\n" +
		" \n" +
		" command:\n" +
		"   help  Display help \n" +
		"   keys  Generate a pair of RSA Keys (The public key is generated in JWK format) \n" +
		"   list  List all secrets for JWK file \n" +
		"   view  Preview the settings file \n" +
		"   set   Store a key (only A-Z and \"standard\" digits) and value (in quotes) \n" +
		"   eject Generate a settings file \n" +
		" \n" +
		" command [options]:\n" +
		"   help  [-h] \n" +
		"   keys  [-g] [jwk-id] \n" +
		"   list  [-k] [path-to-jwk] \n" +
		"   view  [-k] [path-to-jwk] \n" +
		"   set   [-k] [path-to-jwk] \"[key]\" \"[value]\" (eg. ellis set -k key.json \"MY_KEY\" \"$4564%*&3@#\")\n" +
		"   eject [-k] [path-to-jwk] \n" +
		" \n" +
		"path-to-jwk:\n" +
		" 	The path to JWK file to execute.\n" +
		" \n" +
		"about:\n" +
		"	https://github.com/yanjustino/ellis\n" +
		"	MIT License\n" +
		"	Data: 05/08/2020\n" +
		"	Copyright (c) 2020 Yan Justino\n"
}

func (args HelpArgs) AfterExecute() {
	panic("implement me")
}
