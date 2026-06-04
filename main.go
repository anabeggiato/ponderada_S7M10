package main

import (
	"log"

	"github.com/anabeggiato/pond3_M10S07/config"
	_ "github.com/anabeggiato/pond3_M10S07/docs"
	"github.com/anabeggiato/pond3_M10S07/handler"
	"github.com/anabeggiato/pond3_M10S07/repository"
	"github.com/anabeggiato/pond3_M10S07/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Figurinhas
// @version 1.0
// @description API simples para gerenciar figurinhas.
// @host localhost:8080
// @BasePath /
func main() {
	db := config.ConnectDatabase()

	figurinhaRepo := repository.NewFigurinhaRepository(db)
	figurinhaSvc := service.NewFigurinhaService(figurinhaRepo)
	figurinhaHandler := handler.NewFigurinhaHandler(figurinhaSvc)

	r := setupRouter(figurinhaHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRouter(figurinhaHandler *handler.FigurinhaHandler) *gin.Engine {
	r := gin.Default()
	figurinhas := r.Group("/figurinhas")
	{
		figurinhas.POST("", figurinhaHandler.Create)
		figurinhas.GET("", figurinhaHandler.List)
		figurinhas.GET("/:id", figurinhaHandler.GetByID)
		figurinhas.PUT("/:id", figurinhaHandler.Update)
		figurinhas.DELETE("/:id", figurinhaHandler.Delete)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
