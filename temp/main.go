// Package main is a temperary program for experimenting with code that will be used in the tests of the MEA package. It generates a random tree (not "uniformly random", but good enough for the tests) that will eventually be used to generate random department/manager structures to test the efficiency of MEA.
package main

import (
	"fmt"
	"time"
	"math/rand"
)

// ParentList gives an []int each item of which gives the index of its parent node.
func ParentList(order int) []int {
	var p []int
	for i:=0;i<order;i++{
		p = append(p, rand.Intn(1+i))
	}
	return p
}

type Node struct {
	Label, ParentLabel int
	HasChild bool
}

func (n *Node) String() string {
	return	fmt.Sprintf("Index: %v\nParent Index: %v\nHas Child? %v",n.Label, n.ParentLabel, n.HasChild)
}

func makeTree(p []int) map[int]*Node {
	m := make(map[int]*Node)
	for i := len(p)-1; (0 < i) && m[i] == nil; i-- {
		m[i] = &Node{i,p[i],false}
		for x := p[i]; (0 < x) && m[x] == nil; x = p[x] {
			m[x] = &Node{x,p[x],true}
		}
	}
	m[0] = &Node{0,0,true}
	return m
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i<2; i++ {
		p := ParentList(1+rand.Intn(25))
		for k,v:= range makeTree(p) {
			fmt.Println(k)
			fmt.Println(v)
			fmt.Println()
		}
	}
}
