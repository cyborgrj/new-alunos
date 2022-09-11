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

	// Gerar moeda para exibir no contexto, a moeda e seus valores não são gravados no banco.
	coinFiber, _, errCoin := models.ToCoin(aluno.CoinName)
	if errCoin != nil {
		return errCoin
	}
	aluno.CoinResponse = *coinFiber

	// Gerar a data para exibir no contexto mas não gravar no banco.
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", aluno.Datanascimento)
	aluno.Idade = int32(Age(datanasc, hoje))

	return ctx.Status(http.StatusBadRequest).JSON(aluno)

}

func GetAlunosFiber(ctx *fiber.Ctx) error {
	cntx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var singleAluno *models.Aluno
	var alunos []models.Aluno
	a := &models.NewAluno{}

	query, err := a.QueryFiber(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoNaoEncontrado.Error())
	}

	if query.Cpf != "" {
		cpf_aluno := query.Cpf
		err := alunoCollection.FindOne(cntx, bson.M{"cpf": cpf_aluno}).Decode(&singleAluno)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoNaoEncontrado.Error())
		}
		// Calcular data de nascimento
		hoje := time.Now()
		datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
		singleAluno.Idade = int32(Age(datanasc, hoje))
	} else if query.Name != "" {
		nome_aluno := query.Name
		filter := bson.M{"name": bson.M{"$regex": nome_aluno, "$options": "im"}}
		results, err := alunoCollection.Find(cntx, filter)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoNaoEncontrado.Error())
		}
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
	} else {
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

	return ctx.Status(http.StatusOK).JSON(singleAluno)
}

func UpdateAlunoFiber(ctx *fiber.Ctx) error {
	cntx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a := &models.NewAluno{}

	updated_aluno, err := a.FromFiberUpdate(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = alunoCollection.FindOneAndUpdate(cntx, bson.M{"cpf": updated_aluno.Cpf}, bson.M{"$set": updated_aluno}).Decode(&updated_aluno)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoInvalido.Error())
	}

	// recuperar o aluno atualizado no banco para exibir no contexto, em vez dos dados do aluno anteriores à atualização
	err = alunoCollection.FindOne(cntx, bson.M{"cpf": updated_aluno.Cpf}).Decode(&updated_aluno)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoNaoEncontrado.Error())
	}

	// Gerar a data para exibir no contexto mas não gravar no banco.
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", updated_aluno.Datanascimento)
	updated_aluno.Idade = int32(Age(datanasc, hoje))

	return ctx.Status(http.StatusBadRequest).JSON(updated_aluno)
}

func DeleteAlunoFiber(ctx *fiber.Ctx) error {
	cntx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a := &models.NewAluno{}

	query, err := a.QueryFiber(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = alunoCollection.FindOneAndDelete(cntx, bson.M{"cpf": query.Cpf}).Decode(&query)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(custom_errors.ErrAlunoInvalido.Error())
	}

	return ctx.Status(http.StatusBadRequest).JSON(query)
}
