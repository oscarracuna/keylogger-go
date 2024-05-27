package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "io"
)

func main() {
  go keylogger()
  waitForExitSignal()
}
func keylogger() {
	//Gotta write the keylogger here but this will do while I learn
	inputFile := "home/chopin/Desktop"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file: ", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 24)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error reading input file: ", err)
			return
		}

		fmt.Printf("Key event: %v\n", buffer)
	}
}

func waitForExitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Exiting...")
	os.Exit(0)
}
