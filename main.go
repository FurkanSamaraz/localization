package main

import (
	"localization/endpoints"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// @title          Maple Modules API
// @version        1.0
// @description    This is a sample server for a Maple Localization API.
// @contact.name   API Support
// @contact.email  team@workmaple.com
// @host           localhost:8081
// @BasePath       /api/v1
// @schemes        http https
// @Accept         json
// @Produce        json

func main() {
	// ...

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	// Set Middlewares
	app.Use(cors.New())
	app.Use(filesystem.New(filesystem.Config{
		Root:         http.Dir("./locales"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))
	endpoints.SetupPoints(app)

	// listent to port 9981
	app.Listen(":8081")
}
