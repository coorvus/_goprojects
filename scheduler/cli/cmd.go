package cli

import "coorvus/scheduler/todo"

type Cmd struct {
	Entry string
	Value string
}

func handleCmd(cmd Cmd) error {
	entry := cmd.Entry

	switch entry {

	case "list":
		todo.PrintAll(cmd.Value)

	case "add":
		t := todo.New(cmd.Value)
		t.Save()

	case "complete":
		todo.Complete(cmd.Value)

	case "delete":
		todo.Delete(cmd.Value)

	default:
		// Nothing
	}

	return nil
}
