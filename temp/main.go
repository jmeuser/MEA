// Package main is a temperary program for experimenting with code that will be used in the tests of the MEA package. It generates a random tree (not "uniformly random", but good enough for the tests) that will eventually be used to generate random department/manager structures to test the efficiency of MEA.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ParentList gives an []int each item of which gives the index of its parent node.
func ParentList(order int) []int {
	var p []int
	for i := 0; i < order; i++ {
		p = append(p, rand.Intn(1+i))
	}
	return p
}

type Node struct {
	Label, ParentLabel int
	HasChild           bool
}

func (n *Node) String() string {
	return fmt.Sprintf("Label: %v\nParent Label: %v\nHas Child? %v", n.Label, n.ParentLabel, n.HasChild)
}

// makeTree takes a []int p, each of whose entries is less than its index, and returns a map that takes an index to its corresponding tree Node under the assumption that p[i] is the index of node i's parent.
func makeTree(p []int) map[int]*Node {
	m := make(map[int]*Node)
	for i := len(p) - 1; 0 < i; i-- {
		if m[i] == nil {
			m[i] = &Node{i, p[i], false}
			for x := p[i]; (0 < x) && m[x] == nil; x = p[x] {
				m[x] = &Node{x, p[x], true}
			}
		}
	}
	m[0] = &Node{0, 0, true}
	return m
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		p := ParentList(1 + rand.Intn(10))
		fmt.Println(p)
		for k, v := range makeTree(p) {
			fmt.Println(k)
			fmt.Println(v)
			fmt.Println()
		}
	}
	m := make(map[string][]string)
	fmt.Println("m is", m)
	fmt.Println("Is m nil?", m == nil)
	var x map[string][]string
	fmt.Println("x is",x)
	fmt.Println("Is x nil?", x == nil)
	y := make(map[string][]string)
	fmt.Println("y[\"Hello\"] is", y["Hello"])
	fmt.Println("Is y[\"Hello\"] nil?", y["Hello"] == nil)
	fmt.Println(make(map[string][]string))
}
