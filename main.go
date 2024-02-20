package main

import (
	"fmt"

	"github.com/jyz0309/loglite/cmd"
	"github.com/jyz0309/loglite/sqlparser"
)

func main() {
	cmd.StdOut()
	err := sqlparser.Parse("show files from xx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
