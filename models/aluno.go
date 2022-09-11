package models

import (
	"context"
	"log"
	ce "new-alunos/custom_errors"
	pb "new-alunos/protoimp"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/paemuri/brdoc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validate = validator.New()
)

type NewAluno struct {
	Name           string `json:"name,omitempty" validate:"required,min=3,max=32"`
	Serie          string `json:"serie,omitempty" validate:"required"`
	Cpf            string `json:"cpf,omitempty" validate:"required"`
	Email          string `json:"email,omitempty" validate:"required,email"`
	Datanascimento string `json:"datanascimento,omitempty" validate:"required"`
	Cep            string `json:"cep,omitempty" validate:"required"`
	Coin           string `json:"coin,omitempty" validate:"required,min=3,max=4"`
}

type Aluno struct {
	Id             string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Serie          string `json:"serie,omitempty"`
	Cpf            string `json:"cpf,omitempty"`
	Email          string `json:"email,omitempty"`
	Datanascimento string `json:"datanascimento,omitempty"`
	Idade          int32  `json:"idade,omitempty"`
	Endereco       Address
	CoinName       string `json:"coin,omitempty"`
	CoinResponse   CoinResponse
}

func (a NewAluno) IsValid() error {
	err := validate.Struct(a)
	if err != nil {
		return err
	}

	if !brdoc.IsCPF(a.Cpf) {
		return ce.ErrCPFInvalido
	}

	return nil
}

func (a NewAluno) QueryFiber(ctx *fiber.Ctx) (*Aluno, error) {
	query := &NewAluno{}
	dados_aluno := &Aluno{}

	err := ctx.BodyParser(query)
	if err != nil {
		return nil, err
	}

	dados_aluno.Cpf = query.Cpf
	dados_aluno.Name = query.Name

	return dados_aluno, nil
}

func (a NewAluno) FromFiber(ctx *fiber.Ctx) (*Aluno, error) {
	payload := &NewAluno{}

	err := ctx.BodyParser(payload)
	if err != nil {
		return nil, err
	}

	err = payload.IsValid()
	if err != nil {
		return nil, err
	}

	endereco, _, err := ToAdress(payload.Cep)
	if err != nil {
		log.Printf("%v", ce.ErrCEPnaoRecuperado)
		return nil, ce.ErrCEPnaoRecuperado
	}

	aluno_criado := &Aluno{}
	aluno_criado.Id = primitive.NewObjectID().String()
	aluno_criado.Name = payload.Name
	aluno_criado.Serie = payload.Serie
	aluno_criado.Cpf = payload.Cpf
	aluno_criado.Email = payload.Email
	aluno_criado.Datanascimento = payload.Datanascimento
	aluno_criado.CoinName = payload.Coin
	aluno_criado.Endereco = *endereco

	return aluno_criado, nil
}

func (a NewAluno) FromFiberUpdate(ctx *fiber.Ctx) (*Aluno, error) {
	payload := &NewAluno{}

	err := ctx.BodyParser(payload)
	if err != nil {
		return nil, err
	}

	err = payload.IsValid()
	if err != nil {
		return nil, err
	}

	endereco, _, err := ToAdress(payload.Cep)
	if err != nil {
		log.Printf("%v", ce.ErrCEPnaoRecuperado)
		return nil, ce.ErrCEPnaoRecuperado
	}

	aluno_criado := &Aluno{}
	aluno_criado.Name = payload.Name
	aluno_criado.Serie = payload.Serie
	aluno_criado.Cpf = payload.Cpf
	aluno_criado.Email = payload.Email
	aluno_criado.Datanascimento = payload.Datanascimento
	aluno_criado.Endereco = *endereco

	return aluno_criado, nil
}

func (a *NewAluno) FromProto(ctx context.Context, in *pb.NewAluno) (*pb.Aluno, error) {

	a = &NewAluno{
		Name:           in.GetName(),
		Email:          in.GetEmail(),
		Cpf:            in.GetCpf(),
		Serie:          in.GetSerie(),
		Datanascimento: in.GetDatanascimento(),
		Cep:            in.GetCep(),
	}

	err := a.IsValid()
	if err != nil {
		return nil, err
	}

	created_aluno := &pb.Aluno{

		Name:           a.Name,
		Email:          a.Email,
		Cpf:            a.Cpf,
		Serie:          a.Serie,
		Datanascimento: a.Datanascimento,
		Id:             primitive.NewObjectID().String(),
	}

	_, endereco, err := ToAdress(a.Cep)
	if err != nil {
		log.Printf("%v", ce.ErrCEPnaoRecuperado)
		return nil, ce.ErrCEPnaoRecuperado
	}

	log.Printf("endereço: %v", endereco)
	created_aluno.Endereco = endereco

	return created_aluno, nil
}
