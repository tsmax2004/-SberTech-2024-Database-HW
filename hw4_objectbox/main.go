package main

import (
	"fmt"
	"github.com/objectbox/objectbox-go/objectbox"
	"hw4_objectbox/model"
)

func initObjectBox() (*objectbox.ObjectBox, error) {
	return objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
}

func main() {
	// load objectbox
	ob, err := initObjectBox()
	if err != nil {
		panic(err)
	}

	defer ob.Close() // In a server app, you would just keep ob and close on shutdown

	box := model.BoxForTask(ob)

	// Create
	id, err := box.Put(&model.Task{
		Title: "Shop list",
		Text:  "Buy milk",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID: %d\n", id)

	task, err := box.Get(id) // Read
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved task's text: %s\n", task.Text)

	// Update
	task.Text += " & some bread"
	if _, err = box.Put(task); err != nil {
		panic(err)
	}

	task, err = box.Get(id) // Read
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated task's text: %s\n", task.Text)

	// Delete
	if err = box.Remove(task); err != nil {
		panic(err)
	}
}
