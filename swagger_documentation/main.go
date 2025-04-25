package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "documentation/docs" // This will be replaced with your own docs

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User API
// @version 1.0
// @description This is a simple API for user management
// @host localhost:8080
// @BasePath /api/v1

// UserType represents the type of user
// @Description Type of user account
type UserType string

const (
	CustomerType UserType = "customer"
	AdminType    UserType = "admin"
)

// BaseUser represents the base user model
// @Description Base user account information
type BaseUser struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	UserType UserType `json:"user_type"`
	Password string   `json:"-"` // Password will not be shown in the response
}

// CustomerResponse represents the customer user response
// @Description Customer user account information for API responses
type CustomerResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	UserType    UserType `json:"user_type"`
	LoyaltyTier string   `json:"loyalty_tier,omitempty"`
}

// AdminResponse represents the admin user response
// @Description Admin user account information for API responses
type AdminResponse struct {
	ID          int      `json:"id" example:"1"`
	Name        string   `json:"name" example:"Admin User"`
	Email       string   `json:"email" example:"admin@example.com"`
	UserType    UserType `json:"user_type" example:"admin"`
	AccessLevel string   `json:"access_level,omitempty" example:"Super Admin"`
	Department  string   `json:"department,omitempty" example:"IT"`
}

// CreateUserRequest represents the request to create a user
// @Description Request body for creating a new user
type CreateUserRequest struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	UserType    UserType `json:"user_type"`
	LoyaltyTier string   `json:"loyalty_tier,omitempty"` // Only for customers
	AccessLevel string   `json:"access_level,omitempty"` // Only for admins
	Department  string   `json:"department,omitempty"`   // Only for admins
}

// UsersDB is a simple in-memory database
var UsersDB = []BaseUser{
	{ID: 1, Name: "John Doe", Email: "john@example.com", UserType: CustomerType, Password: "secret"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", UserType: AdminType, Password: "secret"},
}

// @Summary Get a list of all users
// @Description Get all users from the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} BaseUser
// @Router /users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UsersDB)
}

// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} CustomerResponse "When user type is customer"
// @Success 200 {object} AdminResponse "When user type is admin"
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get ID from URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
		return
	}

	// Find user
	for _, user := range UsersDB {
		if user.ID == id {
			// Return different response based on user type
			if user.UserType == CustomerType {
				customerResp := CustomerResponse{
					ID:          user.ID,
					Name:        user.Name,
					Email:       user.Email,
					UserType:    user.UserType,
					LoyaltyTier: "Gold", // Hardcoded for example
				}
				json.NewEncoder(w).Encode(customerResp)
			} else {
				adminResp := AdminResponse{
					ID:          user.ID,
					Name:        user.Name,
					Email:       user.Email,
					UserType:    user.UserType,
					AccessLevel: "Super Admin", // Hardcoded for example
					Department:  "IT",          // Hardcoded for example
				}
				json.NewEncoder(w).Encode(adminResp)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User object"
// @Success 201 {object} CustomerResponse "When user type is customer"
// @Success 201 {object} AdminResponse "When user type is admin"
// @Failure 400 {object} map[string]string
// @Router /users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userRequest CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate user type
	if userRequest.UserType != CustomerType && userRequest.UserType != AdminType {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user type. Must be 'customer' or 'admin'"})
		return
	}

	// Create base user
	var baseUser BaseUser
	baseUser.Name = userRequest.Name
	baseUser.Email = userRequest.Email
	baseUser.Password = userRequest.Password
	baseUser.UserType = userRequest.UserType

	// Assign a new ID (simple implementation)
	if len(UsersDB) > 0 {
		baseUser.ID = UsersDB[len(UsersDB)-1].ID + 1
	} else {
		baseUser.ID = 1
	}

	// Add to DB
	UsersDB = append(UsersDB, baseUser)

	// Return different response based on user type
	w.WriteHeader(http.StatusCreated)

	if userRequest.UserType == CustomerType {
		customerResp := CustomerResponse{
			ID:          baseUser.ID,
			Name:        baseUser.Name,
			Email:       baseUser.Email,
			UserType:    baseUser.UserType,
			LoyaltyTier: userRequest.LoyaltyTier,
		}
		json.NewEncoder(w).Encode(customerResp)
	} else {
		adminResp := AdminResponse{
			ID:          baseUser.ID,
			Name:        baseUser.Name,
			Email:       baseUser.Email,
			UserType:    baseUser.UserType,
			AccessLevel: userRequest.AccessLevel,
			Department:  userRequest.Department,
		}
		json.NewEncoder(w).Encode(adminResp)
	}
}

// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CreateUserRequest true "User object"
// @Success 200 {object} CustomerResponse "When user type is customer"
// @Success 200 {object} AdminResponse "When user type is admin"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get ID from URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
		return
	}

	// Decode request body
	var userRequest CreateUserRequest
	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate user type
	if userRequest.UserType != CustomerType && userRequest.UserType != AdminType {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user type. Must be 'customer' or 'admin'"})
		return
	}

	// Find user index
	userIndex := -1
	for i, user := range UsersDB {
		if user.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	// Update user
	UsersDB[userIndex].Name = userRequest.Name
	UsersDB[userIndex].Email = userRequest.Email
	if userRequest.Password != "" {
		UsersDB[userIndex].Password = userRequest.Password
	}
	UsersDB[userIndex].UserType = userRequest.UserType

	// Return different response based on user type
	if userRequest.UserType == CustomerType {
		customerResp := CustomerResponse{
			ID:          UsersDB[userIndex].ID,
			Name:        UsersDB[userIndex].Name,
			Email:       UsersDB[userIndex].Email,
			UserType:    UsersDB[userIndex].UserType,
			LoyaltyTier: userRequest.LoyaltyTier,
		}
		json.NewEncoder(w).Encode(customerResp)
	} else {
		adminResp := AdminResponse{
			ID:          UsersDB[userIndex].ID,
			Name:        UsersDB[userIndex].Name,
			Email:       UsersDB[userIndex].Email,
			UserType:    UsersDB[userIndex].UserType,
			AccessLevel: userRequest.AccessLevel,
			Department:  userRequest.Department,
		}
		json.NewEncoder(w).Encode(adminResp)
	}
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", getUsers)
		r.Get("/users/{id}", getUser)
		r.Post("/users", createUser)
		r.Put("/users/{id}", updateUser)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
