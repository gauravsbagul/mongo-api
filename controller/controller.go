package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gauravsbagul/mongo-api/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURL string = "mongodb+srv://gauravsbagul:gauravsbagul@cluster0.hfrot.mongodb.net"

const dbName string = "netflix"

const colName string = "watchlist"

//!IMPORTANT: MOST

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo db connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reference is ready")

}

// MONGODB HELPERS - file

// INSERT 1 RECORD
func insertOneMovie(movie model.Netflix) *mongo.InsertOneResult {
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("insertion res is", res)
	return res
}

// UPDATE 1 RECORD
func updateOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("id res is", id)

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count is", result.ModifiedCount)

}

// DELETE 1 RECORD
func deleteOneMovie(movieId string) *mongo.DeleteResult {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("id res is", id)

	filter := bson.M{"_id": id}

	deletedCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deletedCount is", deletedCount)

	return deletedCount

}

// DELETE ALL RECORD
func deleteAllMovies() *mongo.DeleteResult {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleteResult is", deleteResult)

	return deleteResult
}

// GET ALL RECORD
func getAllMovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cursor is", cursor)

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
		fmt.Println("cursor is", movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// ACTUAL CONTROLLERS - file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertedMovie := insertOneMovie(movie)

	json.NewEncoder(w).Encode(insertedMovie)
}

func MarkAsReadMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode("Movie updated")
}

func DeleteOneMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deletedCount := deleteOneMovie(params["id"])

	fmt.Println("deletedCount", deletedCount)
	json.NewEncoder(w).Encode("Movie deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedCount := deleteAllMovies()

	fmt.Println("deletedCount", deletedCount)
	json.NewEncoder(w).Encode("All Movie deleted")
}
