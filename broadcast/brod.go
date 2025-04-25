package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up  1
	FlagBroadcast                      // supports broadcast access capability 2(10)
	FlagLoopback                       // is a loopback interface 4 (100)
	FlagPointToPoint                   // belongs to a point-to-point link 8 (1000)
	FlagMulticast                      // supports multicast access capability 16 (10000)
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 } // 10|10000 

func main() {
	var v Flags =  FlagUp 
	fmt.Printf("%b %t\n", v, 1<<1==0.0) 
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true" 
	
}
