package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// mongoUri := os.Getenv("MONGOURI")
	// if mongoUri == ""{

	// }
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	//ping no banco de dados
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connectado ao MongoDB.")
	return client
}

//Instância de Client
var DB *mongo.Client = ConnectDB()

//Pegado as coleções do DB
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// collection := client.Database("golangAPI").Collection((collectionName))
	collection := client.Database("alunos").Collection((collectionName))
	return collection
}
