package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}
type admin struct {
	person user
	level  string
}
type Duration int64

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

	// var arr [2]int
	// x := &arr
	// fmt.Println(x)
	// x[0] = 100
	// x[1] = 200
	// fmt.Println(arr)

	// slice := make([]int, 10)
	// slice[0] = 188

	// fmt.Println(&foo(slice)[0])
	// fmt.Println(foo(slice))
	// fmt.Println(&slice[0])

	// dic := make(map[[]string]string)
	// fmt.Println(dic)

	// v, exists := color["blue"]

	// if exists {
	// 	fmt.Println(v)
	// }

	// color := map[int]int{}

	// for i := 0; i <= 10; i++ {
	// 	color[i] = i
	// }
	// x := []int{1, 2}

	// var bill user
	// fmt.Println(bill)

	// Lisa := user{
	// 	name:       "Lisa",
	// 	email:      "lisa@mail.ru",
	// 	ext:        123,
	// 	privileged: true,
	// }
	// fred := admin{
	// 	person: user{
	// 		name:       "Mazha",
	// 		email:      "mazhitov.t@list.ru",
	// 		ext:        777,
	// 		privileged: true,
	// 	},
	// 	level: "admin",
	// }
	// fmt.Println(Lisa)
	// fmt.Println(fred)
	// var d Duration
	// d = int64(1000)
	// fmt.Println(d)

}

// func foo(a []int) []int {

// 	a = append(a, 500)
// 	a[0] = 100
// 	return a
// }
