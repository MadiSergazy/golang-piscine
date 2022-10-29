package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	glas := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = glas
	} else {
		glas.Next = l.Head
		l.Head = glas
	}
}

// Working code for CHECKPOINT////////////////////////////////////////////////////////////////////////
func ListSize(l *List) int {
	golova := l.Head
	count := 0
	for golova != nil {
		count++
		golova = golova.Next
	}
	return count
}
