package main

import (
	"log"
	"time"
)

func main() {
	var t interface{}
	t = time.Now().UnixMilli()
	t2, ok := t.(int64)
	t3 := time.Duration(t2)

	t4 := time.UnixMilli(int64(t3))
		formattedDate := t4.Format("2006-Jan-02")
	if !ok {
		log.Print("some error")
	}
	log.Println(formattedDate)
}
