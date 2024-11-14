package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task by index & specify a new task. id:new_task")
	flag.IntVar(&cf.Del, "del", -1, "specify todo with index no to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.Add(cf.Add)
		todos.print()
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid Format for edit. Please use index:new_task")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)
		}
		todos.EditList(index, parts[1])
		todos.print()
	case cf.Toggle != -1:
		todos.ToggleStatus(cf.Toggle)
		todos.print()
	case cf.Del != -1:
		todos.Delete(cf.Del)
		todos.print()
	default:
		fmt.Println("Invalid Command")
	}

}

func main() {

	todos := Todos{}

	data := NewStorage[Todos]("todos.json")
	data.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	data.Save(todos)
}
