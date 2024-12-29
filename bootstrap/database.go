package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
)

func NewMongoDBConnection(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBHost
	dbUser := env.DBUser
	dbPass := env.DBPass
	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPass, dbHost)
	if dbUser == "" || dbPass == "" {
		log.Fatal("You must set your 'MONGODB_USER' and 'MONGODB_PASS' environment variables. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
