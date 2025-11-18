package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func InitMongo() {
	_ = godotenv.Load(".env") // load if available

	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")

	if uri == "" {
		log.Fatal("MONGO_URI not set")
	}
	if db == "" {
		log.Fatal("MONGO_DB not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("mongo connect error:", err)
	}

	MongoDB = client.Database(db)
	log.Println("MongoDB connected:", db)
}
