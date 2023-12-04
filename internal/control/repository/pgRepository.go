package repository

import (
	"SberApi/internal/control"
	"SberApi/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type controlRepo struct {
	db     *sqlx.DB
}

func NewControlRepository(db *sqlx.DB) control.Repository {
	return &controlRepo{db: db}
}

func (c *controlRepo) AddNewTask(ctx context.Context, params models.NewTask) error {
	_, err := c.db.ExecContext(ctx, queryAddNewTask, params.UUID, params.Title, params.Description, params.Data)
	if err != nil {
		return errors.Wrapf(err, "controlRepo.AddNewTask.ExecContext()")
	}
	return nil
}

func (c *controlRepo) GetInfoTask(ctx context.Context, params models.Task) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetInfoTask, params.UUID)
	if err != nil {
		return result, errors.Wrapf(err, "controlRepo.GetInfoTask.ExecContext(UUID:%s)", params.UUID)
	}
	return result, err
}

func (c *controlRepo) UpdateTask(ctx context.Context, params models.Task) error {
	_, err := c.db.ExecContext(ctx, queryUpdateTask, params.UUID)
	if err != nil {
		return errors.Wrapf(err, "controlRepo.UpdateTask.ExecContext()")
	}
	return nil
}

func (c *controlRepo) DeleteTask(ctx context.Context, params models.Task) error {
	_, err := c.db.ExecContext(ctx, queryDeleteTask, params.UUID)
	if err != nil {
		return errors.Wrapf(err, "controlRepo.DeleteTask.ExecContext()")
	}
	return nil
}

func (c *controlRepo) GetAllTask(ctx context.Context) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetAllTask)
	if err != nil {
		return result, errors.Wrapf(err, "controlRepo.GetInfoTask.ExecContext(UUID:%s)")
	}
	return result, err
}
