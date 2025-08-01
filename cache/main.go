package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
// anurag version
type Cache struct {
	Data   map[string]string
	Expiry map[string]int64
}

func (ch *Cache) Set(inp, val string, exp int64) {
	if inp != "" && val != "" {
		ch.Data[inp] = val
	}
	if exp > 0 {
		ch.Expiry[inp] = exp
	}
}

func (ch *Cache) Get(inp string) string {
	timeNow := time.Now().Unix()
	if ch.Data[inp] == "" || ch.Expiry[inp] < timeNow {
		delete(ch.Data, inp)
		delete(ch.Expiry, inp)
		return ""
	}
	return ch.Data[inp]
}

var cache Cache
var CurrTime int64


func main2() {
	cache = Cache{
		Data:   map[string]string{},
		Expiry: map[string]int64{},
	}

	http.HandleFunc("/set", cacheFn)
	http.HandleFunc("/get", getFn)
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
	var expTime int64
	if kv.Exp > 0 {
		duration := kv.Exp * int64(time.Second)
		expTime = time.Now().Add(time.Duration(duration)).Unix()
		cache.Set(kv.Key, kv.Val, expTime)
	}

	cache.Set(kv.Key, kv.Val, 0)
	// fmt.Println(cache.Data)
	// fmt.Println(cache.Expiry)
	fmt.Println(kv)
}

func getFn(res http.ResponseWriter, req *http.Request) {
	val := req.URL.Query()["key"][0]
	// fmt.Println(val)
	// fmt.Println(cache.Get(val))
	res.Write([]byte(cache.Get(val)))
}
