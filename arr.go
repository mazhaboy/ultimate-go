package main

import "fmt"

func main() {
	// array := [5]*int{0: new(int), 1: new(int)}
	// fmt.Println(array)
	// *array[0] = 10
	// *array[1] = 20
	// fmt.Println(*array[0])

	// a := [5]string{"a", "b", "c", "d", "e"}
	// b := [5]string{}
	// b = a
	// b[0] = "X"
	// fmt.Print(*a[0])
	// fmt.Println(&a[0])
	// fmt.Println(&b[0])

	// c := [3]*string{0: new(string), 1: new(string), 2: new(string)}
	// *c[0] = "Pira"
	// *c[1] = "Nurbek"
	// *c[2] = "Mazha"
	// d := [3]*string{}
	// d = c

	// a := [2][2]int{0: {0: 10, 1: 20}, 1: {0: 35, 1: 45}}
	// a[0][0] = 100
	// a[0][1] = 200
	// a[1][0] = 300
	// a[1][1] = 400
	// b := [2][2]int{}
	// b = a
	// fmt.Println(a[0][0])
	// fmt.Println(b[0][0])
	// var arr [2]int = a[0]
	// fmt.Println(arr)
	// var c int = a[0][0]
	// fmt.Println(c)

	var arr [2]int
	x := &arr
	x[0] = 100
	x[1] = 200
	fmt.Println(arr)

}

// func foo(arr *[2]int){
// 	for i := range *arr {
// 		arr[i] = i
// 	}
// }
