package usecase

import (
	"go-clean-architecture/entity"
	"go-clean-architecture/repository"
	"go-clean-architecture/validator"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]entity.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (entity.TaskResponse, error)
	CreateTask(task entity.Task) (entity.TaskResponse, error)
	UpdateTask(task entity.Task, userId uint, taskId uint) (entity.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUseCase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUseCase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUseCase {
	return &taskUseCase{tr, tv}
}

func (tu *taskUseCase) GetAllTasks(userId uint) ([]entity.TaskResponse, error) {
	tasks := []entity.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []entity.TaskResponse{}
	for _, v := range tasks {
		t := entity.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUseCase) GetTaskById(userId uint, taskId uint) (entity.TaskResponse, error) {
	task := entity.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) CreateTask(t entity.Task) (entity.TaskResponse, error) {
	if err := tu.tv.TaskValidate(t); err != nil {
		return entity.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&t); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) UpdateTask(t entity.Task, userId uint, taskId uint) (entity.TaskResponse, error) {
	if err := tu.tv.TaskValidate(t); err != nil {
		return entity.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&t, userId, taskId); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
