package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db_name string

func connectDB() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb+srv://abhinav:Abhinav@cluster0.9q0f4.mongodb.net/CatTest?retryWrites=true&w=majority&appName=Cluster0"
	}

	db_name = os.Getenv("DB_NAME")
	if db_name == "" {
		db_name = "test"
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	log.Println("Connected to MongoDB")
}

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := client.Database(db_name).Collection("users")
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	user.ID = res.InsertedID.(string)
	json.NewEncoder(w).Encode(user)
}

func add(a, b int) int {
	return a + b
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := 10
	y := 20
	result := add(x, y)
	fmt.Println("Result:", result)
	json.NewEncoder(w).Encode(map[string]int{"result": result})
}

func main() {
	connectDB()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", getUser)
		r.Post("/users", createUser)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not set
	}

	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
