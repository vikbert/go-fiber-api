package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-fiber-api/database"
	"github.com/vikbert/go-fiber-api/model"
)

type ResponseUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserNameData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func createResponseUser(user model.User) ResponseUser {
	return ResponseUser{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

func (ctrl *UserController) List(c *fiber.Ctx) error {
	var (
		users         []model.User
		responseUsers []ResponseUser
	)
	database.Database.Db.Find(&users)
	for _, user := range users {
		responseUser := createResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *model.User) error {
	database.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user not found by given :id")
	}

	return nil
}

func (ctrl *UserController) Read(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user model.User

	if err != nil {
		return c.Status(400).JSON("Invalid :id given")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createResponseUser(user))
}

func (ctrl *UserController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Invalid :id given")
	}

	var userName UserNameData
	if err := c.BodyParser(&userName); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user model.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = userName.FirstName
	user.LastName = userName.LastName
	database.Database.Db.Save(&user)

	return c.Status(200).JSON(createResponseUser(user))
}

func (ctrl *UserController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(":id must be an integer!")
	}

	var user model.User
	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user, uint(id)).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).SendString("Deleted")
}
