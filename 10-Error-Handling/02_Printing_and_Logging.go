package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	{
		fmt.Println("fmt.Println")
		_, err := os.Open("No-File.txt")

		if err != nil {
			fmt.Println("err : ", err)
		}
		// resualt : err :  open No-File.txt: no such file or directory
	}

	{
		fmt.Println("log.Println")
		_, err := os.Open("No-File.txt")

		if err != nil {
			log.Println("err : ", err)
		}
		// resualt : 2021/08/31 19:25:37 err :  open No-File.txt: no such file or directory
	}

	{
		fmt.Println("setOutPut and log.Println")
		f, err := os.Create("log.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		// set out put for log printing
		log.SetOutput(f)

		f2, err2 := os.Open("No-File.txt")

		if err2 != nil {
			log.Println("err : ", err2)
		}
		defer f2.Close()

		// resualt in log.txt file : 2021/08/31 19:28:32 err :  open No-File.txt: no such file or directory
	}

	{
		fmt.Println("log.Fatalln -> os.Exit()")
		_, err := os.Open("No-File.txt")

		if err != nil {
			//log.Fatalln("err2 : ", err)
			//To run, uncomment the above line
		}

	}

	{
		fmt.Println("log.Panicln")
		_, err := os.Open("No-File.txt")

		if err != nil {
			//log.Panicln("err : ", err)
			//To run, uncomment the above line
		}

		// resualt :
		// goroutine 1 [running]:
		// log.Panicln(0xc000087f38, 0x2, 0x2)
		// /snap/go/current/src/log/log.go:368 +0xae
		// main.main()
		// /home/sajjad/go/src/GolangTraining/10-Error-Handling/02_Printing_and_Logging.go:56 +0x52f
		// exit status 2

		//	deferred functions dont run
	}

	{
		fmt.Println("panic(err)")
		_, err := os.Open("No-File.txt")

		if err != nil {
			panic(err)
			//To run, uncomment the above line
		}

		// resualt :
		// panic: open No-File.txt: no such file or directory

		// goroutine 1 [running]:
		// main.main()
		// /home/sajjad/go/src/GolangTraining/10-Error-Handling/02_Printing_and_Logging.go:85 +0x62c
		// exit status 2
	}

}
