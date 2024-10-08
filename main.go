package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BMS/config"
	"github.com/BMS/database"
	"github.com/BMS/routes"
	"github.com/BMS/indexes"
)

func main() {
	//trying to setup connection
	err := config.LoadConfig()
	if err != nil{
		log.Fatal("Error in loading config",err.Error());
	}
	os.Setenv("GIN_MODE", "release");
	err = database.ConnectRedis();
	if err !=nil {
		log.Fatalf("Error in connecting to Redis: %v",err.Error());
	}
	err = database.ConnectMongoDB();
	if err != nil {
		log.Fatalf("Error in connecting to MongoDB: %v",err.Error());
	}
	err = indexes.SetUpIndexes()
	if err != nil {
		log.Fatalf("Error while creating indexes: %v",err.Error());
	}
	err = database.ConnectMySQLDB();
	if err != nil {
		log.Fatalf("Error in connecting to MySQLDB: %v",err.Error());
	}
	router := routes.SetRoutes();
	// start the server in a separate goroutine
	go func(){
		log.Println("Listenning and Serving on port: 8080")
		err = router.Run(":8080")
		if err != nil {
			fmt.Println("Shutting down the server due to error");
		}
	}()
	
	quit := make(chan os.Signal,1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	<-quit

	log.Println("Shutting down gracefully...")
	database.DisconnectRedis();
	database.DisconnectMongoDB();
	database.DisconnectMySQLDB();
}
