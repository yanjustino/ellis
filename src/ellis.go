/*
MIT License
Data: 05/08/2020
Copyright (c) 2020 Yan Justino

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"ellis.io/command"
	"os"
)

func main() {
	var e error
	_, e = command.KeyGenerateArgs{Args: os.Args}.Execute()
	_, e = command.SetSecretArgs{Args: os.Args}.Execute()
	_, e = command.ListSecretsArgs{Args: os.Args}.Execute()
	_, e = command.ViewSecretsArgs{Args: os.Args}.Execute()
	_, e = command.EjectSecretsArgs{Args: os.Args}.Execute()
	_, e = command.HelpArgs{Args: os.Args}.Execute()

	if e != nil {
		println(e.Error())
	}
}