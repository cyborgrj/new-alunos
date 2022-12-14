package routes

import (
	hdl "new-alunos/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func AlunoRoute(app *fiber.App) {
	app.Post("/aluno", hdl.CreateAlunoFiber)
	app.Get("/alunos", hdl.GetAlunosFiber)
	app.Put("/aluno", hdl.UpdateAlunoFiber)
	app.Delete("/aluno", hdl.DeleteAlunoFiber)
}
