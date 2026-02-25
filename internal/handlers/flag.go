package handler

import (
	"flagd/internal/service"

	"github.com/gofiber/fiber/v3"
)

type FlagHandler struct {
	flagService *service.FlagService
}

func NewFlagHandler(flagService *service.FlagService) *FlagHandler {
	return &FlagHandler{flagService: flagService}
}

func (h *FlagHandler) GetById(c fiber.Ctx) error {
	id := c.Params("id")
	flag, err := h.flagService.GetById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(flag)
}

func (h *FlagHandler) Create(c fiber.Ctx) error {
	var input struct {
		Key         string `json:"key"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	flag, err := h.flagService.CreateFlag(c.Context(), input.Key, input.Name, input.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(flag)
}
