package main

import (
	"context"
	"fmt"
	"net/http"

	userpb "twirp/example.com/myapp/rpc/user"
)

func main() {
	client := userpb.NewUserServiceProtobufClient("http://localhost:8080", &http.Client{})

	// Create a user
	resp, _ := client.CreateUser(context.Background(), &userpb.CreateUserRequest{
		Name:  "Alice",
		Email: "alice@example.com",
	})
	fmt.Println("Created user with ID:", resp.Id)

	// Get the user
	user, _ := client.GetUser(context.Background(), &userpb.GetUserRequest{Id: resp.Id})
	fmt.Println("Fetched user:", user.Name, user.Email)
}
