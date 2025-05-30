package main

import (
	"log"
	"reflect"
)

// in this code block, the weird part is:
//  see in main func, dh is a pointer to int and is nill.
//  but when we pass it to the function someFunc, it is not nil.

type Abhinav interface {
}

func main() {
	// var dh interface{}
	var dh *int
	// dh = nil
	if dh == nil {
		log.Print("dh is nil inside main")
	}
	someFunc(dh)
}

func someFunc(abh Abhinav) {
	log.Print("inside some func", reflect.TypeOf(abh))

	if abh == nil {
		log.Print("abh is nil")
	} else {
		log.Print("abh is not nil. Value is :", abh)
	}
	//* this happens because abh is interface and the value that was passed was pointer type. So this make the type of 'abh' interface as *int and value in nil.
	//* And in golang, for an interface to be nil, both its types and value should be nil. 
}

// uncomment the following function; ie. pointer to interface

// func someFunc(abh *Abhinav) {
// 	log.Print("inside some func")
// 	if abh == nil {
// 		log.Print("abh is nil")
// 	} else {
// 		log.Print("abh is not nil")
// 	}
// }



// from Go Tour 
// package main

// import "fmt"

// type I interface {
// 	M()
// }

// func main() {
// 	var i I
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }
