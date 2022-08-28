package servers

import (
	"log"
	"net"
	"new-alunos/configs"
	hdl "new-alunos/handlers"
	pb "new-alunos/protoimp"

	"google.golang.org/grpc"
)

func GrpcStart() {
	port := configs.EnvGrpcPort()
	if port == "" {
		port = ":50051"
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	aluno_mgmt_server := &hdl.AlunoManagementServer{}
	s := grpc.NewServer()
	pb.RegisterAlunoManegementServer(s, aluno_mgmt_server)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
