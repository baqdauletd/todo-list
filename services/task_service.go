package services

import (
	"todo-list/models"
	"todo-list/repo"
)

type TaskService struct{
	taskRepo *repo.TaskRepo
}

func NewTaskService(taskRepo *repo.TaskRepo) *TaskService{
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTaskService(task *models.Task){
	s.taskRepo.CreateTaskRepo(task)
}

func (s *TaskService) FindTaskService(id uint){
	s.taskRepo.FindTaskRepo(id)
}

func (s *TaskService) ListTaskService(){
	s.taskRepo.ListTaskRepo()
}

func (s *TaskService) UpdateTaskService(task *models.Task){
	s.taskRepo.UpdateTaskRepo(task)
}

func (s *TaskService) DelteTaskService(id uint){
	s.taskRepo.DeleteTaskRepo(id)
}