package main

import (
	"context"
	"github.com/boj/redistore"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"readspacev2/internal/handler"
	"readspacev2/internal/middleware"
	"readspacev2/internal/migration"
	"readspacev2/internal/repository/dbrepo/postgres"
	"readspacev2/internal/repository/dbrepo/redis"
	"readspacev2/internal/usecase"
	"readspacev2/pkg/config"
	"readspacev2/pkg/database"

	"github.com/gin-gonic/gin"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	defer db.Close()

	cfg := config.New()
	ctx := context.Background()
	addr := os.Getenv("REDIS_ADDR")

	if addr == "" {
		addr = "localhost:6379"
	}

	secretKey := os.Getenv("SECRET_KEY")

	store, err := redistore.NewRediStore(10, "tcp", addr, "", []byte(secretKey))
	if err != nil {
		log.Fatalf("Failed to create RediStore: %v", err)
	}

	if err := migration.ApplyMigrations(cfg.DatabaseConnStr); err != nil {
		log.Fatalf("Fatal! Failed to apply migrations: %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	authRepo, _ := redis.NewRedisAuthRepository(ctx)
	authUseCase := usecase.NewAuthenticationUseCase(authRepo, userRepo)
	authHandler := handler.NewAuthenticationHandler(authUseCase, store)

	r := gin.Default()

	r.POST("/login", authHandler.Login)

	r.Use(middleware.AuthenticationMiddleware(store))

	userGroup := r.Group("/user")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.ListAllUsers)
		userGroup.DELETE("/", userHandler.DeleteUserById)
		userGroup.PATCH("/", userHandler.UpdateUser)
		userGroup.PATCH("/password", userHandler.UpdateUserPassword)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
