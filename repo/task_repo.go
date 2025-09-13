package repo

import (
	"todo-list/models"

	"gorm.io/gorm"
)

type TaskRepo struct{
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo{
	return &TaskRepo{db:db}
}

func (r *TaskRepo) CreateTaskRepo(tasks *models.Task) error{
	return r.db.Create(&tasks).Error
}

func (r *TaskRepo) FindTaskRepo(id uint) (*models.Task){
	var task models.Task
	r.db.First(&task, "id=?", id)

	return &task
}

func (r *TaskRepo) ListTaskRepo() (*[]models.Task){
	var tasks []models.Task
	r.db.Find(&tasks)
	return &tasks
}

func (r *TaskRepo) UpdateTaskRepo(task *models.Task){
	r.db.Save(&task)
}

func (r *TaskRepo) DeleteTaskRepo(id uint){
	var task models.Task
	r.db.First(&task, "id=?", id)
	
	r.db.Delete(&task)
}