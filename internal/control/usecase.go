package control

import (
	"SberApi/internal/models"
	"context"
)

type UseCase interface {
	AddNewTask(ctx context.Context, params models.NewTask) error
	GetInfoTask(ctx context.Context, params models.Task) (result []byte, err error)
	UpdateTask(ctx context.Context, params models.Task) error
	DeleteTask(ctx context.Context, params models.Task) error
	GetAllTask(ctx context.Context) (result []byte, err error)
}