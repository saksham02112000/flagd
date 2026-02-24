package handler

import (
	"github.com/gofiber/fiber/v3"
	"flagd/internal/service"
	"flagd/internal/domain"
)
type FlagHandler struct{
	flagService service.FlagService
}

func NewFlagHandler(flagService *service.FlagService) *FlagHandler{
	return &FlagHandler{flagService: *flagService}
}

func (h *FlagHandler) GetById(c *fiber.Ctx) (flag *domain.Flag, err error){
	return h.flagService.GetById(("id"))
}