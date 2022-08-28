package handlers

import (
	"context"
	"net/http"
	"new-alunos/custom_errors"
	"new-alunos/models"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAlunoFiber(ctx *fiber.Ctx) error {
	a := &models.NewAluno{}
	aluno, err := a.FromFiber(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	// var alunoCollection já declarada no grpc-handler
	err = alunoCollection.FindOne(ctx.UserContext(), bson.M{"cpf": aluno.Cpf}).Decode(&aluno)
	if err != mongo.ErrNoDocuments {
		return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrCPFJaCadastrado.Error())
	}

	_, err = alunoCollection.InsertOne(ctx.UserContext(), aluno)
	if err != nil {
		return err
	}

	// Gerar a data para exibir no contexto mas não gravar no banco.
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", aluno.Datanascimento)
	aluno.Idade = int32(Age(datanasc, hoje))

	return ctx.Status(http.StatusBadRequest).JSON(aluno)

}

func GetAlunosFiber(ctx *fiber.Ctx) error {
	cntx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var alunos []models.Aluno
	defer cancel()

	results, err := alunoCollection.Find(cntx, bson.M{})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	//iterando os dados do banco
	defer results.Close(cntx)
	for results.Next(cntx) {
		var singleAluno models.Aluno
		if err = results.Decode(&singleAluno); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		// Calcular data de nascimento
		hoje := time.Now()
		datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
		singleAluno.Idade = int32(Age(datanasc, hoje))
		alunos = append(alunos, singleAluno)
	}
	return ctx.Status(http.StatusOK).JSON(alunos)
}
