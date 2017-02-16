package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sync"

	concurrent "github.com/divyag9/goconcurrent/packages/concurrent"
)

func main() {
	runtime.GOMAXPROCS(2)
	// Parsing the command line arguments
	numRoutines, execCommand := getCommandLineArguments()
	var wg sync.WaitGroup
	wg.Add(*numRoutines)
	// Calling concurrently
	for i := 0; i < *numRoutines; i++ {
		go callExecuteCommand(&wg, *execCommand)
	}
	wg.Wait()
}

func getCommandLineArguments() (*int, *string) {
	numRoutines := flag.Int("numroutines", 1, "Number of routines to run concurrently")
	execCommand := flag.String("execcommand", "", "Command to be executed")
	flag.Parse()
	if *execCommand == "" {
		log.Fatal("please pass the required flags: -execcommand(Command to be executed)")
	}

	return numRoutines, execCommand
}

func callExecuteCommand(wg *sync.WaitGroup, command string) {
	defer wg.Done()
	res, err := concurrent.ExecuteCommand(command)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}
