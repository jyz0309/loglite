package cmd

import "fmt"

func StdOut() {
	fmt.Println("+-------------------+---------------+")
	fmt.Println("|         ID        |      Size     |")
	fmt.Println("+-------------------+---------------+")
	fmt.Println("| 123123            | key           |")
	fmt.Println("+-------------------+---------------+")
}

type StdPrinter struct {

}

func (std *StdPrinter) 