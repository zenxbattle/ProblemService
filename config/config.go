package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIGATEWAYPORT string
	UserGRPCPort   string
	MongoDBURL     string
	ProblemService string
	NatsURL        string
	RedisURL       string
	Environment    string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults:", err)
	}

	mongoUser := getEnv("MONGO_USER", "admin")
	mongoPassword := getEnv("MONGO_PASSWORD", "admin")
	mongoHost := getEnv("MONGO_HOST", "mongodb.default.svc.cluster.local")
	mongoPort := getEnv("MONGO_PORT", "27017")
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPassword, mongoHost, mongoPort)

	config := Config{
		APIGATEWAYPORT: getEnv("APIGATEWAYPORT", "7000"),
		UserGRPCPort:   getEnv("USERGRPCPORT", "50051"),
		MongoDBURL:     mongoURL,
		ProblemService: getEnv("PROBLEMSERVICE", "50055"),
		NatsURL:        getEnv("NATSURL", "nats://localhost:4222"),
		RedisURL:       getEnv("REDISURL", "localhost:6379"),
		Environment:    getEnv("ENVIRONMENT", "development"),
	}

	// fmt.Println(config)
	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
