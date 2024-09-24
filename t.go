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

		signal.Ignore(syscall.SIGINT, syscall.SIGTERM)	// Not Working
	sigChan := make(chan os.Signal, 1)		//chan o.Signal will listen for signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) //Registers the channel to receive notifications
	go func() {
		<-sigChan
		term.Restore(int(os.Stdin.Fd()), oldstate)
		os.Exit(0)
	}()
		fmt.Print("press :q to quit. \n")

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
		break
		}
	}
	
	//Restore terminal and exit
	term.Restore(int(os.Stdin.Fd()), oldstate)
}