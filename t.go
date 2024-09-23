package main

import (
	"fmt"
	"os"
)

func main() {
	// buff := make([]byte,5)
	var c byte
	buff := []byte{c}

	for {
		_, err := os.Stdin.Read(buff)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Printf("%c", buff[0])

		}
}