package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/LeoAntunesBrombilla/readspacev2/docs"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/handler"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/middleware"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/migration"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/dbrepo/postgres"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/dbrepo/redis"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/external"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase"
	"github.com/LeoAntunesBrombilla/readspacev2/pkg/config"
	"github.com/LeoAntunesBrombilla/readspacev2/pkg/database"
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

	bookListRepo := postgres.NewBookListRepository(db)
	bookListUseCase := usecase.NewBookListUseCase(bookListRepo)
	bookListHandler := handler.NewBookListHandler(bookListUseCase)

	booksRepo := postgres.NewBooksRepository(db)
	booksUseCase := usecase.NewBooksUseCase(booksRepo)
	booksHandler := handler.NewBooksHandler(booksUseCase)

	readSessionRepo := postgres.NewReadSessionsRepository(db)
	readSessionUseCase := usecase.NewReadingSessionUseCase(readSessionRepo)
	readSessionHandler := handler.NewReadSessionHandler(readSessionUseCase)

	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.POST("/logout", authHandler.Logout)

	userGroup := r.Group("/user", middleware.AuthenticationMiddleware(store))
	{
		userGroup.GET("/", userHandler.ListAllUsers)
		userGroup.DELETE("/", userHandler.DeleteUserById)
		userGroup.PATCH("/", userHandler.UpdateUser)
		userGroup.PATCH("/password", userHandler.UpdateUserPassword)
	}

	r.POST("/user", userHandler.CreateUser)

	bookListGroup := r.Group("/bookList", middleware.AuthenticationMiddleware(store))
	{
		bookListGroup.POST("/", bookListHandler.Create)
		bookListGroup.GET("/", bookListHandler.ListAllBookList)
		bookListGroup.DELETE("/", bookListHandler.DeleteBookListById)
		bookListGroup.PATCH("/", bookListHandler.UpdateBookList)
	}

	bookGroup := r.Group("/book", middleware.AuthenticationMiddleware(store))
	{
		bookGroup.POST("/", booksHandler.Create)
		bookGroup.DELETE("/", booksHandler.Delete)
	}

	readSessionGroup := r.Group("/readSession", middleware.AuthenticationMiddleware(store))
	{
		readSessionGroup.POST("/", readSessionHandler.CreateReadSession)
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
	// 0-. Alguns endpoints estao funfando sem o token auth
	// 1. Add CRUD for a read session
	// 3. Add userMe endpoint
	// 4. Add documentation
	// 5. Melhorar organizacao de pasta
	// 6. Adicionar testes unitarios e coverage

	r.Run()
}
