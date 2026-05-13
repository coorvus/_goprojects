package todo

import (
	"coorvus/scheduler/storage"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type Todo struct {
	Id          int
	Description string
	IsCompleted bool
	CreatedAt   string
}

func New(desc string) Todo {
	return Todo{
		Id:          1,
		Description: desc,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (t Todo) Save() {

	newTodo := []string{strconv.FormatInt(int64(t.Id), 10), t.Description, t.CreatedAt, strconv.FormatBool(t.IsCompleted)}

	err := storage.Save(newTodo)
	if err != nil {
		fmt.Printf("Error: `%s`\n", err.Error())
		return
	}

	fmt.Printf("Todo `%s` saved !\n", t.Description)
}

func Complete(strId string) {
	intID, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println("Error: Invalid ID, except int value")
		return
	}
	err = storage.Complete(intID)
	if err != nil {
		fmt.Printf("Error: `%s`\n", err.Error())
		return
	}
	fmt.Printf("Todo `%s` completed!\n", strId)

}

func Delete(strId string) {
	intID, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println("Error: Invalid ID, except int value")
		return
	}
	err = storage.Delete(intID)
	if err != nil {
		fmt.Printf("Error: `%s`\n", err.Error())
		return
	}

	fmt.Printf("Todo `%s` deleted!\n", strId)

}

func PrintAll(flag string) {
	tasks, err := storage.GetAll()
	if err != nil {
		fmt.Printf("Error: `%s`\n", err.Error())
		return
	}

	toSendTasks := [][]string{}

	switch flag {
	case "-a":
		toSendTasks = tasks
	case "":
		for _, t := range tasks {
			if t[3] == "true" {
				continue
			}
			toSendTasks = append(toSendTasks, t)
		}
	}

	displayTodo(toSendTasks, flag == "-a")

}

func displayTodo(todos [][]string, all bool) {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	header := []string{"ID", "Description", "Created"}
	if all {
		header = append(header, "Done")
	}

	if len(todos) > 1 {
		todos = todos[1:]
	}

	fmt.Fprintln(w, strings.Join(header, "\t"))

	for _, line := range todos {

		if !all {
			line = line[:3]
		}

		rawDate := line[2]
		if len(rawDate) > 19 {
			rawDate = rawDate[:19]
		}
		todoDate, err := time.Parse("2006-01-02 15:04:05", rawDate)
		if err == nil {
			line[2] = timediff.TimeDiff(todoDate)
		}
		fmt.Fprintln(w, strings.Join(line, "\t"))
	}
	w.Flush()
	fmt.Println()

}
