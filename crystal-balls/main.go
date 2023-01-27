package main

import (
	"fmt"
	"math"
)

func main() {
	b := []bool{}

  for i := 0; i < 101; i++{
    if(i < 99){
      b = append(b, false)
    } else {
      b = append(b, true)
    }
  }

	fmt.Println(CrystalBalls(b))
}

func CrystalBalls(b []bool) int {

	j := int(math.Sqrt(float64(len(b))))

	fmt.Println(j)

	i := j

	for i < len(b) {
		fmt.Println(b[i], i)
		if b[i] {
			break
		}
		i += j
	}
	fmt.Println(i)
	i -= int(j)

	fmt.Println(i)

	for k := 0; k < j && i < len(b); {
		k += 1
		i += 1
		fmt.Println(b[i], i, k)
		if b[i] {
			return i
		}

	}

	return -1

}
