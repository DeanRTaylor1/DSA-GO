package main


import "fmt"

func main() {
  s := "foo"
  t := "bar"

  fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
      return false
    }

    sm := make(map[byte]int)
    tm := make(map[byte]int)
    
    for i := range s {
      if sm[s[i]] != tm[t[i]] {
        return false
      } else {
        //save the last index plus one
        //if we check each time we find it and we have the last index
        //we can ensure that each index matches
        sm[s[i]] = i + 1
        tm[t[i]] = i + 1
      }
    }

    return true

}
