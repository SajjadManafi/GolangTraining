package main

import (
	"fmt"
)

var y = 42

func main() {

	fmt.Println(y)
	fmt.Printf("%v", y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("this : %#x\n", y)
	fmt.Println(s)

	// general printing to standard out
	//func Print(a ...interface{}) (n int, err error)
	//func Printf(format string, a ...interface{}) (n int, err error)
	//func Println(a ...interface{}) (n int, err error)t

	//printing to a string which you can assign to a variable
	//func Sprint(a ...interface{}) string
	//func Sprintf(format string, a ...interface{}) string
	//func Sprintln(a ...interface{}) string

	//printing to a file or a web server’s response
	//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

	//for more : https://pkg.go.dev/fmt

}
