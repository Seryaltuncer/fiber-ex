package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type User struct {
	Name string `json:"name"`
	Surname string`json:"surname"`
	Age string`json:"age"`


}

func main(  ) {

	router := fiber.New()

	router.Get(" /" , func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello there")

	})
	router.Post("/create/:name /:surname /:age" , func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		surname := ctx.Params("surname")
		age := ctx.Params("age")
		tempUser := &User{
			name,surname, age,
		}
		marshal, err := json.Marshal(tempUser)
		if err != nil {
			return ctx.JSON(nil)
		}
		return ctx.JSON(marshal)
	})


    router.Get("/dash",monitor.New())
	err := router.Listen(":3000")
	if err != nil{
		return
	 }
}

