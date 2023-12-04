package control

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	AddNewTask() fiber.Handler
	GetInfoTask() fiber.Handler
	UpdateTask() fiber.Handler
	DeleteTask() fiber.Handler
	GetAllTask() fiber.Handler
}