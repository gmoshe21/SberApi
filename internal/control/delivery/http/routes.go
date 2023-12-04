package http

import(
	"SberApi/internal/control"

	"github.com/gofiber/fiber/v2"
)

func MapAPIRoutes(group fiber.Router, h control.Handlers) {
	group.Post("/task/add_new_task", h.AddNewTask())
	group.Get("/task/get_info_task", h.GetInfoTask())
	group.Put("/task/update_task", h.UpdateTask())
	group.Delete("/task/delete_task", h.DeleteTask())
	group.Get("/task/get_all_task", h.GetAllTask())
}