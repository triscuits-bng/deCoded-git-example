// this is an app that displays output 
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() { 
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter you age: ")
	age, err := buf.ReadBytes('\n')

	if err != nil {
		fmt.Println(age)
	} else (
		fmt.Println(err)
	)
	fmt.Println(age)
	fmt.Println(err)
}

