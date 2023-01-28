package main

import (
	"errors"
	"fmt"
)

type node struct {
	el   interface{}
	next *node
}

type linkedList struct {
	head   *node
	length int
}



func newLinkedList() *linkedList {
	return &linkedList{head: &node{el: "head", next: nil}, length: 1}
}

func (l *linkedList) find(p int) *node {
  //start at the head and traverse through until you find the position
	c := l.head
	for i := 1; i < p-1; i++ {
		c = c.next
	}
	return c
}

func (l *linkedList) insert(n interface{}, p int) (int, error) {
  //handle incorrect input
	if p > l.length + 1 {
		return l.length, errors.New("Position exceeds current length")
	}
  //create new node pointer
	nn := &node{el: n}
  // if the user inserts it at the beginning replace the head
	if p == 1 {
		nn.next = l.head
		l.head = nn
		l.length++
		return l.length, nil
	}
  //find the position before we want to enter
	c := l.find(p)
  //set the next value on the new node as the address of the current next node
	nn.next = c.next
  //set the next node of the current node to ur new node
	c.next = nn
	l.length++
	return l.length, nil
}

func (l *linkedList) remove(p int) (int, error) {
	if p > l.length {
		e := errors.New("Invalid position")
		return l.length, e
	}
  //find the node before the node we are deleting
	c := l.find(p)
  //the node to remove is the next node in the list
	r := c.next
  //set the new next node as removed node's next val
	c.next = r.next
  //set the next value of our removed node to nil to free up mem
	r.next = nil
	l.length--
	return l.length, nil
}

func (l *linkedList) print() {
	c := l.head
	for {
		fmt.Printf("%v", c.el)
		if c.next == nil {
			fmt.Println()
			break
		}
		c = c.next

		fmt.Print(" -> ")
	}
	fmt.Printf("Current Length: %v", l.length)
	fmt.Println()

}

func main() {

	ll := newLinkedList()
	ll.print()
	_, err := ll.insert("p-2", 2)
	if err != nil {
		fmt.Println(err)
	}
	ll.print()
	_, err = ll.insert("p-3", 3)

	if err != nil {
		fmt.Println(err)
	}
	ll.print()
	_, err = ll.insert("p-4", 4)

	if err != nil {
		fmt.Println(err)
	}
	ll.print()
	_, err = ll.insert("new-head", 1)

	if err != nil {
		fmt.Println(err)
	}
	ll.print()
	_, err = ll.insert("new-p-2", 2)

	if err != nil {
		fmt.Println(err)
	}
	ll.print()
	_, err = ll.remove(3)
	if err != nil {
		fmt.Println(err)
	}

	ll.print()
}
