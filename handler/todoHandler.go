package handler

import (
	"bookingApp/constant"
	"bookingApp/helper"
	"bookingApp/todo"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *todoHandler {
	return &todoHandler{todoService}
}

var failed = constant.AcknowledgeEnum.Failed
var success = constant.AcknowledgeEnum.Success

func (h *todoHandler) GetAllTodoHandler(c *fiber.Ctx) error {
	allTodos, err := h.todoService.FindAllTodo()

	if err != nil {
		c.JSON(fiber.Map{
			"Message":     "Failed to load data",
			"Status":      400,
			"Data":        nil,
			"Acknowledge": 0,
		})
		return err
	}

	format := todo.FormatGetAllTodos(allTodos)

	c.JSON(fiber.Map{
		"Message":     "Success to load data",
		"Status":      200,
		"Data":        format,
		"Acknowledge": 1,
	})

	return nil
}

func (h *todoHandler) CreateTodoHandler(c *fiber.Ctx) error {
	var input todo.TodoInput
	c.BodyParser(&input)

	if err := input.ValidateTodoInput(); err != nil {
		response := helper.APIResponse(failed, http.StatusUnprocessableEntity, nil, "Failed to save data")
		return c.Status(422).JSON(response)
	}

	newTodo, err := h.todoService.SaveTodo(input)

	if err != nil {
		response := helper.APIResponse(failed, http.StatusBadRequest, nil, "Failed to save data")
		return c.Status(400).JSON(response)
	}

	format := todo.FormatCreateTodos(newTodo)

	response := helper.APIResponse(success, http.StatusOK, format, "Success to save data")
	return c.Status(http.StatusOK).JSON(response)
}
