package handler

import (
	"errors"
	"flagd/internal/domain"
	"flagd/internal/service"

	"github.com/gofiber/fiber/v3"
)

type FlagHandler struct {
	flagService *service.FlagService
}

func NewFlagHandler(flagService *service.FlagService) *FlagHandler {
	return &FlagHandler{flagService: flagService}
}

// mapError translates domain sentinel errors to the correct HTTP status code.
func mapError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrFlagNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, domain.ErrFlagKeyExists):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, domain.ErrFlagArchived):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, domain.ErrEnvNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
}

func (h *FlagHandler) GetAll(c fiber.Ctx) error {
	flags, err := h.flagService.GetAll(c.Context())
	if err != nil {
		return mapError(c, err)
	}
	return c.JSON(flags)
}

func (h *FlagHandler) GetById(c fiber.Ctx) error {
	id := c.Params("id")
	flag, err := h.flagService.GetById(c.Context(), id)
	if err != nil {
		return mapError(c, err)
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
		return mapError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(flag)
}

func (h *FlagHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.flagService.DeleteFlag(c.Context(), id)
	if err != nil {
		return mapError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *FlagHandler) Toggle(c fiber.Ctx) error {
	id := c.Params("id")
	envSlug := c.Params("envSlug")

	fe, err := h.flagService.ToggleFlag(c.Context(), id, envSlug)
	if err != nil {
		return mapError(c, err)
	}

	return c.JSON(fe)
}
