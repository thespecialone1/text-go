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

	// buff := make([]byte,5)
	var c byte
	buff := []byte{c}

	for  {
		_, err := os.Stdin.Read(buff)
		if err!= nil {
			fmt.Println("Error in displaying input: ", err)
			return
		}
		fmt.Printf("%c", buff[0])
	}

}