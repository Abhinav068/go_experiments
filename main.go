package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Cache struct {
	Data   map[string]string
	Expiry map[int64][]string
}

func (ch *Cache) Set(inp, val string, exp int64) {
	if inp != "" && val != "" {
		ch.Data[inp] = val
	}
	if exp > 0 {
		ch.Expiry[exp] = append(ch.Expiry[exp], inp)
	}
}

func (ch *Cache) Get(inp string) string {
	return ch.Data[inp]
}

var cache Cache
var CurrTime int64

func main() {
	cache = Cache{
		Data:   map[string]string{},
		Expiry: map[int64][]string{},
	}
	CurrTime = time.Now().Unix()
	go func() {
		for {
			<-time.Tick(time.Second)
			CurrTime += int64(time.Second)
			cacheDataKeys := cache.Expiry[CurrTime]
			fmt.Println("CurrTime", CurrTime)
			if len(cacheDataKeys) > 0 {
				fmt.Println(cacheDataKeys)
				for i := range cacheDataKeys {
					delete(cache.Data, cacheDataKeys[i])
				}
				delete(cache.Expiry, CurrTime)
			}
		}
	}()
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
		expTime = CurrTime + int64(time.Duration(duration))
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
