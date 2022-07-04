package main

import (
	"context"
	"fmt"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	// "net/url"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Sleeplessss/gin_mongo_driver/controllers"
	"github.com/Sleeplessss/gin_mongo_driver/routes"
)

var (
	server      *gin.Engine
	ui          controllers.UserInterface
	ur          routes.UserRoutes
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func initDatabase() {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017/?maxPoolSize=20&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!!")

	userc = mongoclient.Database("go_gin_db").Collection("users")
	ui = controllers.NewUserService(userc, ctx)
	ur = routes.New(ui)
	server = gin.Default()

}

func main() {
	initDatabase()
	defer mongoclient.Disconnect(context.TODO())

	basepath := server.Group("/v1")
	ur.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":3000"))
}
