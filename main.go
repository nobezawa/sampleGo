package main

import (
	"fmt"
	"sample/hello"
)

type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}

	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
	hello.Hello()
}
