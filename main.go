package main

import (
	"fmt"
	"time"
	"todo-list/db"
	"todo-list/models"
	"todo-list/repo"
)

func main(){
	db := db.Connect()

	taskRepo := repo.NewTaskRepo(db)

	dueDate := time.Now().Add(24*time.Hour)
	task := models.Task{
		Title: "Learn Go", 
		Description: "Finish Repo Layer",
		Completed: false, 
		DueDate: &dueDate,
		Priority: "High",
	}
	taskRepo.CreateTaskRepo(task)
	fmt.Println("Task Created")

	found := taskRepo.FindTaskRepo(1)
	fmt.Println("Found Task:", found)

	tasks := taskRepo.ListTaskRepo()
	fmt.Println("All Tasks:", tasks)

	found.Completed = true
	taskRepo.UpdateTaskRepo(found)
	fmt.Println("Task Updated", found)

	taskRepo.DeleteTaskRepo(1)
	fmt.Println("Task Deleted")

	tasks1 := taskRepo.ListTaskRepo()
	fmt.Println("All Tasks:", tasks1)
}

/* 
todo
- repo -> service -> handler
todo
- whole CRUD in repo
*/