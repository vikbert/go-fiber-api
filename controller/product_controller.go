package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) Create(c *fiber.Ctx) error {
	return c.Status(204).JSON("")
}

func (ctrl *ProductController) List(c *fiber.Ctx) error {
	return c.Status(204).JSON("")
}

func (ctrl *ProductController) Read(c *fiber.Ctx) error {
	return c.Status(204).JSON("")
}

func (ctrl *ProductController) Update(c *fiber.Ctx) error {
	return c.Status(204).JSON("")
}

func (ctrl *ProductController) Delete(c *fiber.Ctx) error {
	return c.Status(204).JSON("")
}
