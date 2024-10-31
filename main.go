package main 

import "fmt" 

var p int = 3
var q =& p
func main()  {
	fmt.Println(*q) // 3
	fmt.Println(p) // 3
}