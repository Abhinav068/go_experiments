package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	postIDs := []string{"1755857757358", "1755857773971", "1755857782927"}
	totalRequests := 500

	type Counter struct {
		Like    int
		Dislike int
	}

	counts := make(map[string]*Counter)
	for _, id := range postIDs {
		counts[id] = &Counter{}
	}

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, postID := range postIDs {
		for i := 0; i < totalRequests; i++ {
			wg.Add(1)
			go func(id string) {
				defer wg.Done()
				if rand.Intn(2) == 0 {
					_, err := http.Post("http://localhost:8088/like/"+id, "application/json", nil)
					if err == nil {
						mu.Lock()
						counts[id].Like++
						mu.Unlock()
					}
				} else {
					_, err := http.Post("http://localhost:8088/dislike/"+id, "application/json", nil)
					if err == nil {
						mu.Lock()
						counts[id].Dislike++
						mu.Unlock()
					}
				}
			}(postID)
		}
	}

	wg.Wait()

	for _, id := range postIDs {
		fmt.Printf("Post ID %s: Like API hits: %d, Dislike API hits: %d\n", id, counts[id].Like, counts[id].Dislike)
	}
}
