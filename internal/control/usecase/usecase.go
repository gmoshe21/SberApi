package usecase

import (
	"SberApi/config"
	"SberApi/internal/control"
	"SberApi/internal/models"
	"context"
)

type controlUC struct {
	cfg                    *config.Config
	controlRepo            control.Repository
}

func NewControlUseCase( cfg *config.Config, controlRepo control.Repository) control.UseCase {
	return &controlUC{cfg: cfg, controlRepo: controlRepo}
}

func (c *controlUC) AddNewTask(ctx context.Context, params models.NewTask) error {
	return c.controlRepo.AddNewTask(ctx, params)
}

func (c *controlUC) GetInfoTask(ctx context.Context, params models.Task) (result []byte, err error) {
	return c.controlRepo.GetInfoTask(ctx, params)
}

func (c *controlUC) UpdateTask(ctx context.Context, params models.Task) error {
	return c.controlRepo.UpdateTask(ctx, params)
}

func (c *controlUC) DeleteTask(ctx context.Context, params models.Task) error {
	return c.controlRepo.DeleteTask(ctx, params)
}

func (c *controlUC) GetAllTask(ctx context.Context) (result []byte, err error) {
	return c.controlRepo.GetAllTask(ctx)
}
