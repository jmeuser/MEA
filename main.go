package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// riba gives a Random Index Bounded Array of length n.
func riba(n int) []int {
	A := make([]int, n)
	for i := 1; i < n; i++ {
		A[i] = rand.Intn(i)
	}
	return A
}

type Employee struct{ ID, MID, Pos, Dep string }

// ret gives a Random Employee Tree of order n in dep(artment) as a map[string]*Employee.
func ret(n int, dep string) map[string]*Employee {
	M := make(map[string]*Employee)
	A := riba(n)
	for i := n - 1; 0 < i; i-- {
		iid := strconv.Itoa(i)
		if M[iid] == nil {
			imid := strconv.Itoa(A[i])
			ipos := []string{"Developer", "QA Tester"}[rand.Intn(2)]
			M[iid] = &Employee{iid, imid, ipos, dep}
			for j := A[i]; (0 < j) && M[strconv.Itoa(j)] == nil; j = A[j] {
				jid := strconv.Itoa(j)
				jmid := strconv.Itoa(A[j])
				M[jid] = &Employee{jid, jmid, "Manager", dep}
			}
		}
	}
	M["0"] = &Employee{"0", "0", "Manager", dep}
	return M
}
func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(riba(1 + rand.Intn(1+10)))
	}
	for i := 0; i < 5; i++ {
		fmt.Println()
		M := ret(1+rand.Intn(10), "a")
		for i := 0; i < len(M); i++ {
			fmt.Println(*M[strconv.Itoa(i)])
		}
	}
}
