package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	{
		//ex1:
		n, err := fmt.Println("Hey")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(n)
	}

	{
		//If we enter a string instead of a number, we will get an error
		var num1, num2 int

		fmt.Print("num: ")
		_, err := fmt.Scan(&num1)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("num2: ")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("num1: ", num1, " num2: ", num2)

	}

	{
		f, err := os.Create("names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		r := strings.NewReader("Harry Potter")
		io.Copy(f, r)

	}

	{
		//If the file does not exist, we will get an error.
		f, err := os.Open("names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(bs))
	}

}
