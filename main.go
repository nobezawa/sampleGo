package main

import (
	"fmt"
	"sample/hello"
)

type Task struct {
	ID     int
	Detail string
	done   bool
	*User // Userを埋め込む
}

type User struct {
	FileName string
	LastName string
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

func (u *User) FullName() string  {
	fullname := fmt.Sprintf("%s $s", u.FileName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User  {
	return &User{
		FileName: firstName,
		LastName: lastName,
	}
}


func Newtask(id int, detail, firstName, lastName string)  *Task {
	task := &Task{
		ID: id,
		Detail: detail,
		done: bool,
		User: NewUser(firstName, lastName),
	}
	return task
}