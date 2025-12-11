package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    userpb "twirp/rpc/user"
)

// In-memory database
var users = make(map[int32]*userpb.GetUserResponse)
var idCounter int32 = 1

// Implement UserService
type userServer struct{}

func (s *userServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
    id := idCounter
    idCounter++
    users[id] = &userpb.GetUserResponse{
        Name:  req.Name,
        Email: req.Email,
    }
    return &userpb.CreateUserResponse{Id: id}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
    if user, ok := users[req.Id]; ok {
        return user, nil
    }
    return nil, fmt.Errorf("user not found")
}

func main() {
    server := &userServer{}
    twirpHandler := userpb.NewUserServiceServer(server)

    log.Println("Starting Twirp server on :8080")
    http.ListenAndServe(":8080", twirpHandler)
}