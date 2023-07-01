package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) Create(c *fiber.Ctx) error {

	return c.Status(http.StatusCreated).JSON("")
}

func (ctrl *ProductController) List(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("")
}

func (ctrl *ProductController) Read(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("")
}

func (ctrl *ProductController) Update(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("")
}

func (ctrl *ProductController) Delete(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("")
}
