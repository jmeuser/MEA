package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"io/ioutil"
	"encoding/json"
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

// randomEmployeeTree gives []*Employee with length order and IDS from firstNewIID to firstNewID + order in departmetn dep.
func randomEmployeeTree(order int, firstNewID, dep string) ([]*Employee, error) {
	if order < 1 {
		return nil, fmt.Errorf("The order of an Employee Tree must be greater than or equal to one: order = %v", order)
	}
	N, err := strconv.Atoi(firstNewID)
	if err != nil {
		return nil, err
	}
	A := riba(order)
	M := make(map[string]*Employee)
	for i := order - 1; 0 < i; i-- {
		iid := strconv.Itoa(i + N)
		if M[iid] == nil {
			imid := strconv.Itoa(A[i] + N)
			ipos := []string{"Developer", "QA Tester"}[rand.Intn(2)]
			M[iid] = &Employee{iid,imid,ipos,dep}
			for j := A[i] ; (0 < j) && M[strconv.Itoa(j+N)] == nil; j=A[j] {
				jid := strconv.Itoa(j+N)
				jmid:= strconv.Itoa(A[j]+N)
				M[jid] = &Employee{jid,jmid,"Manager",dep}
			}
		}
	}
	// Assumes the root of every employee tree in a department is a manager: refactor to deal with order == 1.
	M[firstNewID] = &Employee{firstNewID, firstNewID, "Manager", dep}
	E := make([]*Employee, 0, order)
	for _, e := range M {
		E = append(E, e)
	}
	return E, nil
}

func makeRBPF(n int, fileName string) error {
	M := ret(n, "a")
	employees := make([]*Employee, 0, len(M))
	for _, e := range M {
		employees = append(employees, e)
	}
	data, err := json.Marshal(employees)
	if err != nil {
		return fmt.Errorf("Marhsal fail: %v", err)
	}
	err = ioutil.WriteFile(fileName + ".json", data, 0600)
	if err != nil {
		return err
	}
	return nil
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
	makeRBPF(10, "testRBPF")
}
