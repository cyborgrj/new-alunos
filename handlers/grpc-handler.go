package handlers

import (
	"context"
	"log"
	"new-alunos/configs"
	ce "new-alunos/custom_errors"
	"new-alunos/models"
	"time"

	pb "new-alunos/protoimp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var alunoCollection *mongo.Collection = configs.GetCollection(configs.DB, "alunos")

func Age(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

type AlunoManagementServer struct {
	pb.UnimplementedAlunoManegementServer
}

func (server *AlunoManagementServer) CreateAluno(ctx context.Context, in *pb.NewAluno) (*pb.Aluno, error) {
	log.Printf("Received: %v", in.GetName())
	a := &models.NewAluno{}

	aluno_criado, err := a.FromProto(ctx, in)
	if err != nil {
		log.Fatalf("error parsing new aluno %v", ce.ErrAlunoInvalido)
	}

	err = alunoCollection.FindOne(ctx, bson.M{"cpf": aluno_criado.Cpf}).Decode(&aluno_criado)
	if err != mongo.ErrNoDocuments {
		return nil, ce.ErrCPFJaCadastrado
	}

	_, err = alunoCollection.InsertOne(ctx, aluno_criado)
	if err != nil {
		return nil, err
	}

	// Gerar a data para exibir no contexto mas não gravar no banco.
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", aluno_criado.Datanascimento)
	aluno_criado.Idade = int32(Age(datanasc, hoje))

	return aluno_criado, nil
}

func (server *AlunoManagementServer) GetAlunos(ctx context.Context, in *pb.GetAlunosParams) (*pb.AlunosList, error) {
	alunos_list := &pb.AlunosList{}
	singleAluno := pb.Aluno{}

	if len(in.GetName()) < 1 && len(in.GetCpf()) < 1 {
		results, err := alunoCollection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		defer results.Close(ctx)

		//iterando os dados do banco
		for results.Next(ctx) {
			singleAluno := pb.Aluno{} // criar um novo aluno para cada loop para receber os dados e adicionar.
			if err = results.Decode(&singleAluno); err != nil {
				return nil, err
			}
			// Calcular data de nascimento
			hoje := time.Now()
			datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
			singleAluno.Idade = int32(Age(datanasc, hoje))
			alunos_list.Alunos = append(alunos_list.Alunos, &singleAluno)
		}
		return alunos_list, nil
	} else if len(in.GetName()) > 1 {
		nome_aluno := in.GetName()
		log.Printf("Received Name: %v", nome_aluno)

		filter := bson.M{"name": bson.M{"$regex": nome_aluno, "$options": "im"}}
		results, err := alunoCollection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}

		// iterando os dados do banco
		for results.Next(ctx) {
			singleAluno := pb.Aluno{} // criar um novo aluno para cada loop para receber os dados e adicionar.
			if err = results.Decode(&singleAluno); err != nil {
				return nil, err
			}
			// Calcular data de nascimento
			hoje := time.Now()
			datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
			singleAluno.Idade = int32(Age(datanasc, hoje))
			alunos_list.Alunos = append(alunos_list.Alunos, &singleAluno)
		}
		return alunos_list, nil
	} else {
		cpf_aluno := in.GetCpf()
		log.Printf("Received CPF: %v", cpf_aluno)

		err := alunoCollection.FindOne(ctx, bson.M{"cpf": cpf_aluno}).Decode(&singleAluno)
		if err != nil {
			return nil, err
		}

		// Calcular data de nascimento
		hoje := time.Now()
		datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
		singleAluno.Idade = int32(Age(datanasc, hoje))
		alunos_list.Alunos = append(alunos_list.Alunos, &singleAluno)

		return alunos_list, nil
	}
}

func (server *AlunoManagementServer) UpdateAluno(ctx context.Context, in *pb.NewAluno) (*pb.Aluno, error) {
	singleAluno := &pb.Aluno{}
	a := &models.NewAluno{}

	cpf_aluno := in.GetCpf()
	log.Printf("Received CPF: %v", cpf_aluno)

	err := alunoCollection.FindOne(ctx, bson.M{"cpf": cpf_aluno}).Decode(&singleAluno)
	if err != nil {
		return nil, ce.ErrCPFNaoEncontrado
	}

	aluno_criado, err := a.FromProto(ctx, in)
	if err != nil {
		log.Fatalf("error parsing new aluno %v", ce.ErrAlunoInvalido)
	}

	err = alunoCollection.FindOneAndUpdate(ctx, bson.M{"cpf": aluno_criado.Cpf}, bson.M{"$set": aluno_criado}).Decode(&aluno_criado)
	if err != nil {
		return nil, err
	}

	// recuperar o aluno atualizado no banco para exibir no contexto, em vez dos dados do aluno anteriores à atualização
	err = alunoCollection.FindOne(ctx, bson.M{"cpf": aluno_criado.Cpf}).Decode(&aluno_criado)
	if err != nil {
		return nil, err
	}

	// Gerar a data para exibir no contexto mas não gravar no banco.
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", aluno_criado.Datanascimento)
	aluno_criado.Idade = int32(Age(datanasc, hoje))

	return aluno_criado, nil

}

func (server *AlunoManagementServer) DeleteAluno(ctx context.Context, in *pb.GetAlunosParams) (*pb.Aluno, error) {
	singleAluno := &pb.Aluno{}

	cpf_aluno := in.GetCpf()
	log.Printf("Received CPF: %v", cpf_aluno)

	err := alunoCollection.FindOneAndDelete(ctx, bson.M{"cpf": cpf_aluno}).Decode(&singleAluno)
	if err != nil {
		return nil, err
	}

	// Calcular data de nascimento
	hoje := time.Now()
	datanasc, _ := time.Parse("02/01/2006", singleAluno.Datanascimento)
	singleAluno.Idade = int32(Age(datanasc, hoje))

	return singleAluno, nil
}
