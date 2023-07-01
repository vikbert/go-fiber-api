package controller

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-fiber-api/database"
	"github.com/vikbert/go-fiber-api/model"
	"net/http"
)

type ResponseProduct struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type ProductData struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type ProductController struct {
}

func createResponse(p model.Product) ResponseProduct {
	return ResponseProduct{
		ID:           p.ID,
		Name:         p.Name,
		SerialNumber: p.SerialNumber,
	}
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) Create(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	return c.Status(http.StatusCreated).JSON(createResponse(product))
}

func (ctrl *ProductController) List(c *fiber.Ctx) error {
	var (
		products         []model.Product
		responseProducts []ResponseProduct
	)

	database.Database.Db.Find(&products)
	for _, product := range products {
		responseProducts = append(responseProducts, createResponse(product))
	}

	return c.Status(http.StatusOK).JSON(responseProducts)
}

func findById(id int, product *model.Product) error {
	database.Database.Db.Find(&product, "id=?", id)

	fmt.Println("found product with ID: ", product.ID)
	if product.ID == 0 {
		return errors.New("product not found by given ID")
	}

	return nil
}

func (ctrl *ProductController) Read(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("ID is not an integer!")
	}
	var product model.Product

	if err := findById(id, &product); err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(createResponse(product))
}

func (ctrl *ProductController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("ID is not an integer!")
	}

	var postData ProductData
	if err := c.BodyParser(&postData); err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var product model.Product
	if err := findById(id, &product); err != nil {
		c.Status(http.StatusNotFound).SendString(err.Error())
	}

	product.Name = postData.Name
	product.SerialNumber = postData.SerialNumber

	database.Database.Db.Save(&product)
	return c.Status(http.StatusOK).JSON(createResponse(product))
}

func (ctrl *ProductController) Delete(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("")
}
