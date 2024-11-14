package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Task        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(task string) {
	todo := Todo{
		Task:        task,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
	fmt.Println("---Task added Successfully---")
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	fmt.Println("---Task deleted Successfully---")
	return nil
}

func (todos *Todos) ToggleStatus(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	fmt.Println("---Task toggled Successfully---")
	return nil
}

func (todos *Todos) EditList(index int, task string) error {
	t := *todos

	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	t[index].Task = task

	fmt.Println("---Task edited Successfully---")
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetHeaders("No", "Task", "Completed", "Created At", "Completed At")
	table.SetRowLines(false)

	for index, t := range *todos{
		completed := "❌"
		completedAt := ""

		if t.Completed{
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Task, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
