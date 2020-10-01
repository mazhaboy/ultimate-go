package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	numb int
}

// type sum struct {
// 	a int
// 	b int
// }
var x notifier

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
func (u *admin) notify() {
	fmt.Printf("Sending Admin Email To <%d>\n", u.numb)
}

//func (s user) sum() {
//fmt.Println(s.name, "and", s.email)
//}

// func (u *user) changeEmail(email string) {
// 	u.email = email
// }

func main() {

	// sm := sum{5, 6}
	x := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	mazha := admin{
		user: user{
			name:  "Bill",
			email: "bill@email.com",
		},

		numb: 1998,
	}

	send(&x)
	send(&mazha)

}
func send(n notifier) {
	n.notify()
}

// func sendSum(n notifier) {
// 	n.sum()
// }

// package main

// import "fmt"

// main is the entry point for the application. 13
// func main() {
// var b bytes.Bufer

// Write a string to the buffer.
// b.Write([]byte("Hello"))

// Use Fprintf to concatenate a string to the Buffer.
// fmt.Fprintf(&b, " World!")

// Write the content of the Buffer to stdout.
// io.Copy(os.Stdout, &b)
// 	a := Local
// 	fmt.Println(a)
// }
