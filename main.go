package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/set", cacheFn)
	// http.HandleFunc("/get", getFn)
	http.ListenAndServe(":8888", nil)
}
func cacheFn(res http.ResponseWriter, req *http.Request) {
	var kv struct {
		Key string `json:"key"`
		Val string `json:"val"`
		Exp int64  `json:"exp"`
	}
	err := json.NewDecoder(req.Body).Decode(&kv)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(kv)
	val, _ := json.Marshal(kv)
	res.Write(val)
}
