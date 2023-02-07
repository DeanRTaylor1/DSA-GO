package main

import "fmt"

/*func GetGrade(a, b, c int) rune {*/
	/*grades := []rune("ABCDF")*/
  /*fmt.Println(grades[0]) */
	/*switch avg := (a + b + c) / 3; {*/
	/*case avg > 90:*/
		/*return grades[0]*/
	/*case avg > 80:*/
		/*return grades[1]*/
	/*case avg > 70:*/
		/*return grades[2]*/
	/*case avg > 60:*/
		/*return grades[3]*/
	/*default:*/
		/*return grades[4]*/
	/*}*/
/*}*/

/*func main() {*/

/*//	testA := [...][3]int{{95, 90, 93}, {100, 85, 96}, {92, 93, 94}}*/
	/*fmt.Println(GetGrade(95, 90, 93))*/
/*}*/


func ToAlternatingCase(str string) string {
  
  for _, v := range str {
    
    switch ; {
      case v == 32:
        continue
      case v > 90:
        fmt.Println(string(v))
      default:
        fmt.Println(string(v))
    }
  }
  return str
}
