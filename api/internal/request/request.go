package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoadAndValidateBody(body any, c *fiber.Ctx) error {
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	return validator.New().Struct(body)
}

func LoadAndValidateQuery(query any, c *fiber.Ctx) error {
	if err := c.QueryParser(query); err != nil {
		fmt.Println(err)
		return err
	}
	return validator.New().Struct(query)
}
