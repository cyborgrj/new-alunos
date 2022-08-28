package servers

import (
	"new-alunos/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func HttpStart() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{"data": "Fiber server listening on :6000"})
	})

	routes.AlunoRoute(app)

	app.Listen(":6000")

}
