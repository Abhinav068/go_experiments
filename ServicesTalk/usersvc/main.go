package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


// User represents user data
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Products []Product `json:"products,omitempty"`
	Age      int       `json:"age,omitempty"`
} 
// Product represents product data
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Mock database of users
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

func main() {
	// Define routes
	http.HandleFunc("/users", getAllUsers)
	http.HandleFunc("/users/", getUserByID)

	// Get port from environment variable or use default
	port := getEnv("USER_SERVICE_PORT", "8081")
	productServicePort := getEnv("PRODUCT_SERVICE_PORT", "8082")
	productServiceHost := getEnv("PRODUCT_SERVICE_HOST", "localhost")

	// Set global product service URL for use in requests
	productServiceURL = fmt.Sprintf("http://%s:%s", productServiceHost, productServicePort)

	// Start server
	fmt.Printf("User service starting on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// getAllUsers returns all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Add products to users by fetching from product service
	enrichedUsers := make([]User, len(users))
	copy(enrichedUsers, users)

	for i := range enrichedUsers {
		products, err := getProductsForUser(enrichedUsers[i].ID)
		if err != nil {
			log.Printf("Failed to get products for user %d: %v", enrichedUsers[i].ID, err)
			// Continue even if we can't get products for a user
		} else {
			enrichedUsers[i].Products = products
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrichedUsers)
}

// getUserByID returns a user by ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	hex := "67da74ef62f77d54f9a73a37"
	hexobj,_ := primitive.ObjectIDFromHex(hex)
	log.Println(hexobj)

	// Find user
	var foundUser *User
	var age int
	for i := range users {
		if users[i].ID == id {
			foundUser = &users[i]
			break
		}
		age += i
	}

	if foundUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get products for this user
	products, err := getProductsForUser(id)
	if err != nil {
		log.Printf("Failed to get products for user %d: %v", id, err)
		// Continue without products
	} else {
		userCopy := *foundUser
		userCopy.Products = products
		userCopy.Age = age
		foundUser = &userCopy
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundUser)
}

// Global variable to store product service URL
var productServiceURL string

// getProductsForUser fetches products for a specific user from the product service
func getProductsForUser(userID int) ([]Product, error) {
	getProductUrl := fmt.Sprintf("%s/products/user/%d", productServiceURL, userID)
	log.Print(getProductUrl)
	resp, err := http.Get(getProductUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// getEnv returns the value of an environment variable or a default value if not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value

}