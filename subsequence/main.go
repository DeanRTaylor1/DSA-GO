package main

import "fmt"

func isSubsequence(s string, t string) bool {

	k := 0
	for i := 0; i < len(t); i++ {
		if k < len(s) && s[k] == t[i] {
			k++
		}

	}
	return len(s) == k
}


func main() {
  fmt.Println(isSubsequence("b", "c"))
}
