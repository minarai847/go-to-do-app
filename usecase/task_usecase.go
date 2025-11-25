package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.Task, error)
	GetTaskById(userId uint, taskId uint) (model.Task, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := make([]model.TaskResponse, len(tasks))
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:         v.ID,
			Title:      v.Title,
			Created_At: v.Created_At,
			Updated_At: v.Updated_At,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}
