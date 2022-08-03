package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println((err.Error()))
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

//connect to mongodb

func Db() {
	db_uri := GoDotEnvVariable("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(("all good"))
	}
}
