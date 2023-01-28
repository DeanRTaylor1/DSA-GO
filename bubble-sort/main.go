package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 50; i++ {
		r := rand.Intn(1000)
		a = append(a, r)
	}

	fmt.Println(a)
	fmt.Println(bubbleSort(a))
}

func bubbleSort(a []int) []int {
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a)-1-i; j++{
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
	}
	return a
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}
