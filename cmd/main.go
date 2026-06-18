package main

import (
	"log"
	"net"
	"time"
	"xcode/cache"
	configs "xcode/config"
	"xcode/mongoconn"
	"xcode/natsclient"
	"xcode/repository"
	"xcode/service"

	problemService "github.com/lijuuu/GlobalProtoXcode/ProblemsService"
	redisboard "github.com/lijuuu/RedisBoard"
	"go.uber.org/zap"

	"xcode/logutil"

	"google.golang.org/grpc"
)

func mustConnect(name string, fn func() error) {
	for {
		err := fn()
		if err == nil {
			log.Printf("Connected to %s", name)
			return
		}
		log.Printf("Retrying connection to %s: %v", name, err)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	config := configs.LoadConfig()

	var nc *natsclient.Client
	mustConnect("NATS", func() error {
		var err error
		nc, err = natsclient.NewClient(config.NatsURL)
		return err
	})

	// Initialize Zap logger based on environment
	var logger *zap.Logger
	var err error
	if config.Environment == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("Failed to initialize Zap logger: " + err.Error())
	}
	defer logger.Sync()

	// Initialize Logger
	logShipper := logutil.New("problem-service")

	redisCacheClient := cache.NewRedisCache(config.RedisURL, "", 0)

	mongoclientInstance := mongoconn.ConnectDB(config.MongoDBURL)

	// Initialize RedisBoard Leaderboard
	lbConfig := redisboard.Config{
		Namespace:   "user_Leaderboard_Unique",
		K:           10,
		MaxUsers:    1_000_000,
		MaxEntities: 200,
		FloatScores: true,
		RedisAddr:   config.RedisURL,
	}
	var lb *redisboard.Leaderboard
	mustConnect("RedisBoard", func() error {
		var lbErr error
		lb, lbErr = redisboard.New(lbConfig)
		return lbErr
	})
	defer lb.Close()

	repoInstance := repository.NewRepository(mongoclientInstance, lb, logShipper)

	serviceInstance := service.NewService(*repoInstance, nc, *redisCacheClient, lb, logShipper)

	serviceInstance.StartCronJob()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+config.ProblemService)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", config.ProblemService, err)
	}

	grpcServer := grpc.NewServer()
	problemService.RegisterProblemsServiceServer(grpcServer, serviceInstance)

	log.Printf("ProblemService gRPC server running on port %s", config.ProblemService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
