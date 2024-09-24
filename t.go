package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {

		oldstate, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error while turning changing terminal mode: ", err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldstate)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		term.Restore(int(os.Stdin.Fd()), oldstate)
		os.Exit(0)
	}()

		fmt.Print("press :q to quit.")

	// buff := make([]byte,1) if use this then no need to define c
	var c byte
	buff := []byte{c}
	var inputBuffer []byte

	for  {
		_, err := os.Stdin.Read(buff)
		if err!= nil {
			fmt.Println("Error in displaying input: ", err)
			return
		}
		fmt.Printf("%c", buff[0])

		// Append the charector into buffer
		inputBuffer =append(inputBuffer, buff[0])
		// to check if the charectors entered are :q
		if len(inputBuffer)>= 2 && string(inputBuffer[len(inputBuffer)-2:]) == ":q" {
		fmt.Println("You pressed :q : ")	
		break
		}
	}
	//Restore terminal and exit
	term.Restore(int(os.Stdin.Fd()), oldstate)
}