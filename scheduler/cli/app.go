package cli

import "fmt"



func Run(args []string) {
	
	cmd, err := formatArgs(args)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	err = handleCmd(*cmd)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
}