package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Divak-ar/mongo_go_api/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// MONGODB HELPERs (goes into separate files) , helpers starts with lowercase letter cause we expect them not to be transported/imported 

func insertOneMovie(movie model.Netflix){
	// context has to be added when working with db 
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id: ", inserted.InsertedID)
}

// updating one record
func updateOneMovie(movieId string){

	// get the movie id :- convert the string into something(an object _id) that mongoDB understand(uses bson) -> in js , mongoose madde it easy to work with mongoDB without directly dealing with mongoDB .... mongoose are odm (object data modelling) , prisma is orm (object relational mapping). ORM are used with relational db (like Mysql, Postgres) and ODM are used with non-relational db (like MOngoDb, Cassandra, Redis)
	id, _ := primitive.ObjectIDFromHex(movieId)

	// this id will be used to run through all queries in mongoDB
	// find the movie in the record with _id same as id (movieId provided by the user)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// takes context, condition for update, new update value 
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Modified Records : ", result.ModifiedCount)
}

// delete one record 
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted Records : ", result.DeletedCount)
}

// delete all records
func deleteAllMovie() int64{
	// no providing any values so everything will be deleted
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	// collection.DeleteMany(context.Background(), bson.D{{}})


	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted Records : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from mongo

func getAllMovies() []primitive.M{
	// we will get an cursor (not the data directly) we have to loop over cursor to get the answer (uses .Next() property for traversing - Linkedlist)...........get all movies and append it to our return array/slice , use bson.D{{}} or bson.M{{}} -> differ with ordering of data and case sensitive data (ABc then ABc only)

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil{
		log.Fatal(err)
	}

	// looping over the cursor ;- we have to use slice of primitive type to pass as the reference data strcuture and type for decode of the cursor data (which is just slice/array of objects/struct)

	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movies)
		if err != nil {
			log.Fatal(err)
		}
		// appending all the movies
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// Now Controllers will uses these helpers to perform operation and return response to the request

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){

	// setting headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	// calling the helper func
	allMovies := getAllMovies()

	// returning a json response (by writer)
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){

	// setting headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	
}



params := mux.Vars(r)
var data model.Netflix
_ = json.Decoder(r.Body).Decode(&data)

