package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yerinadler/go-ddd/config"
	"github.com/yerinadler/go-ddd/internal/api"
	"github.com/yerinadler/go-ddd/internal/application"
	"github.com/yerinadler/go-ddd/internal/infrastructure"
)

func main() {

	config, err := config.GetConfig()

	if err != nil {
		log.Fatal("Can not load environment variable")
	}

	router := gin.Default()

	// Initialise services & domain objects
	db_client := infrastructure.GetMongoConnection(config.GetString("MONGODB_URI"))
	order_repo := infrastructure.NewMongoOrderRepository(db_client.Database(config.GetString("MONGODB_DATABASE")))
	order_service := application.NewOrderApplicationService(order_repo)

	// Register API Handlers
	api.NewOrderHandler(router, order_service)
	api.NewCommonHandler(router)

	fmt.Println(order_repo)
	router.Run(":8888")
}
