package main

import (
	"fmt"
_	"strconv"
	"math/rand"
	"time"
)

// riba gives a Random Index Bounded Array of length n.
func riba(n int) []int {
	A := make([]int, n)
	for i := range A {
		A[i] = rand.Intn(1 + i)
	}
	return A
}
type Employee struct{ ID, MID, Pos, Dep string }

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(riba(1 + rand.Intn(1 + 10)))
	}
}
