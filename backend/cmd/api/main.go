package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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

	//Vamos aplicar o migrate aqui
	if err := migration.ApplyMigrations(cfg.DatabaseConnStr); err != nil {
		log.Fatalf("Fatal! Failed to apply migrations: %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	authRepo, _ := redis.NewRedisAuthRepository(ctx)
	authUseCase := usecase.NewAuthenticationUseCase(authRepo, userRepo)

	authHandler := handler.NewAuthenticationHandler(authUseCase)

	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.Use(middleware.AuthenticationMiddleware())
	r.POST("/user", func(c *gin.Context) { userHandler.CreateUser(c.Writer, c.Request) })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
