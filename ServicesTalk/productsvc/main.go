package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Product represents product data
type Product struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID int     `json:"user_id"`
}

// User represents minimal user data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Mock database of products
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99, UserID: 1},
	{ID: 2, Name: "Phone", Price: 699.99, UserID: 1},
	{ID: 3, Name: "Tablet", Price: 399.99, UserID: 2},
	{ID: 4, Name: "Headphones", Price: 149.99, UserID: 3},
	{ID: 5, Name: "Monitor", Price: 299.99, UserID: 2},
}

func main() {
	// Define routes
	http.HandleFunc("/products", getAllProducts)
	http.HandleFunc("/products/", getProductByID)
	http.HandleFunc("/products/user/", getProductsByUserID)

	// Get port from environment variable or use default
	port := getEnv("PRODUCT_SERVICE_PORT", "8082")
	userServicePort := getEnv("USER_SERVICE_PORT", "8081")
	userServiceHost := getEnv("USER_SERVICE_HOST", "localhost")

	// Set global user service URL for use in requests
	userServiceURL = fmt.Sprintf("http://%s:%s", userServiceHost, userServicePort)

	// Start server
	fmt.Printf("Product service starting on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// getAllProducts returns all products
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Enrich products with user info
	enrichedProducts := make([]Product, len(products))
	copy(enrichedProducts, products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrichedProducts)
}

// getProductByID returns a product by ID
func getProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Skip /products/ prefix to get the ID
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Find product
	var foundProduct *Product
	for i := range products {
		if products[i].ID == id {
			foundProduct = &products[i]
			break
		}
	}

	if foundProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Get user info for this product
	user, err := getUserInfo(foundProduct.UserID)
	if err != nil {
		log.Printf("Failed to get user info for product %d: %v", id, err)
		// Continue without user info
	}

	// Create response with product and user info
	type ProductResponse struct {
		Product
		OwnerName string `json:"owner_name,omitempty"`
	}

	response := ProductResponse{
		Product: *foundProduct,
	}

	if user != nil {
		response.OwnerName = user.Name
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getProductsByUserID returns all products for a specific user
func getProductsByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Skip /products/user/ prefix to get the ID
	idStr := r.URL.Path[len("/products/user/"):]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find all products for this user
	var userProducts []Product
	for _, product := range products {
		if product.UserID == userID {
			userProducts = append(userProducts, product)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProducts)
}

// Global variable to store user service URL
var userServiceURL string

// getUserInfo fetches user information from the user service
func getUserInfo(userID int) (*User, error) {
	usrsvcUrl := fmt.Sprintf("%s/users/%d", userServiceURL, userID)
	resp, err := http.Get(usrsvcUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// getEnv returns the value of an environment variable or a default value if not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}


