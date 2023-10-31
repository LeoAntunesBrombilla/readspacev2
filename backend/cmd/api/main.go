package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"readspacev2/docs"
	"readspacev2/internal/entity"
	"readspacev2/internal/handler"
	"readspacev2/internal/middleware"
	"readspacev2/internal/migration"
	"readspacev2/internal/repository/dbrepo/postgres"
	"readspacev2/internal/repository/dbrepo/redis"
	"readspacev2/internal/repository/external"
	"readspacev2/internal/usecase"
	"readspacev2/pkg/config"
	"readspacev2/pkg/database"

	"github.com/boj/redistore"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

//@title Readspace API
//@version 1.0
//@description This is a sample server celler server.

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.api_key Bearer
// @in header
// @name Authorization
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
	realBcryptWrapper := entity.RealBcryptWrapper{}
	userHandler := handler.NewUserHandler(userUseCase, realBcryptWrapper)

	authRepo, _ := redis.NewRedisAuthRepository(ctx)
	authUseCase := usecase.NewAuthenticationUseCase(authRepo, userRepo)
	authHandler := handler.NewAuthenticationHandler(authUseCase, store)

	externalBookServiceRepo := external.NewExternalBookRepository()
	externalBookServiceUseCase := usecase.NewExternalBookServiceUseCase(externalBookServiceRepo)
	externalBookServiceHandler := handler.NewExternalBookServiceHandler(externalBookServiceUseCase)

	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.POST("/logout", authHandler.Logout)

	userGroup := r.Group("/user", middleware.AuthenticationMiddleware(store))
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.ListAllUsers)
		userGroup.DELETE("/", userHandler.DeleteUserById)
		userGroup.PATCH("/", userHandler.UpdateUser)
		userGroup.PATCH("/password", userHandler.UpdateUserPassword)
	}

	externalBookServiceGroup := r.Group("/searchBook", middleware.AuthenticationMiddleware(store))
	{
		externalBookServiceGroup.GET("/", externalBookServiceHandler.SearchBooks)
	}

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//TODO
	// 1. Add CRUD for a read session
	// 2. Add CRUD for books
	// 3. Add userMe endpoint
	// 3. Add documentation

	r.Run()
}
