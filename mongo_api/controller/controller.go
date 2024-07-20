package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://divakar:WEVktyAqtwgC0jLY@cluster0.evxlsbr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// Most important : 
var collection *mongo.Collection
// A collection is a grouping of MongoDB documents . Basically db=collections and tables=documents (sql/mongodb)

// Connect with mongoDB
// a init (initialization method) func runs only one time to initialize the the program
func init(){

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	// Whenever we connect with any other service we use context to describe the lifetime of the service with respect the app. In Go, context is used to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. It's part of the context package and is crucial for managing the lifecycle of processes and ensuring efficient resource utilization.
	// context.Background() : used when you want to keep it running in the background not dependent on other resources
	// context.TODO() : use when unsure which context to use or are planning to replace it later. It's a placeholder for future context decisions.
	// context.WithTimeout() and context.WithCancel()

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success!")

	//  creating collection/db
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}
