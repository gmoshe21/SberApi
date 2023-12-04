package http

import (
	"SberApi/config"
	"SberApi/internal/control"
	"SberApi/internal/models"
	"SberApi/pkg/reqvalidator"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type controlHandlers struct {
	cfg       *config.Config
	controlUC control.UseCase
}

func NewControlHandlers(cfg *config.Config, controlUC control.UseCase) control.Handlers {
	return &controlHandlers{cfg: cfg, controlUC: controlUC}
}

func (ctrl *controlHandlers)AddNewTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.NewTask

		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.AddNewTask.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.AddNewTask(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers)GetInfoTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.Task
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.GetInfoTask.ReadRequest", err)
			return err
		}

		result, err := ctrl.controlUC.GetInfoTask(context.Background(), params)
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}

func (ctrl *controlHandlers)UpdateTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.Task
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.UpdateTask.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.UpdateTask(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers)DeleteTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.Task
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.DeleteTask.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.DeleteTask(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers)GetAllTask() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := ctrl.controlUC.GetAllTask(context.Background())
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}