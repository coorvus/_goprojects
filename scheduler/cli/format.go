package cli

import (
	"fmt"
	"slices"
)

var allCmds []string = []string{"complete", "list", "add", "delete"}


func formatArgs(args []string) (*Cmd, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("\nYou must provide some cmd\n")
	}

	exeptCmd := args[0]

	if !slices.Contains(allCmds, exeptCmd) {
		return nil, fmt.Errorf("\nError: Unknown cmd\n")
	}

	var value string
	if len(args) > 1 {
		value = args[1]
	}


	return &Cmd{
		Entry: exeptCmd,
		Value: value,
	},  nil

}