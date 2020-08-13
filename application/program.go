package application

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
