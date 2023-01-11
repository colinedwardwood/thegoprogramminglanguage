package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	for i, arg := range os.Args {
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
	fmt.Println(s)
}
