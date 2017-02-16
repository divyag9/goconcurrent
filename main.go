package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sync"

	concurrent "github.com/divyag9/goconcurrent/packages/concurrent"
)

//Result represents the result of running the command
type Result struct {
	res string
	err error
}

func main() {
	runtime.GOMAXPROCS(2)
	// Parsing the command line arguments
	numRoutines, execCommand := getCommandLineArguments()
	var wg sync.WaitGroup
	wg.Add(*numRoutines)
	// Calling concurrently
	for i := 0; i < *numRoutines; i++ {
		result := make(chan Result)
		go callExecuteCommand(result, &wg, *execCommand)
		output := <-result
		fmt.Printf("Result:%s, Error: %s", output.res, output.err)
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

func callExecuteCommand(result chan Result, wg *sync.WaitGroup, command string) {
	defer wg.Done()
	res, err := concurrent.ExecuteCommand(command)
	result <- Result{res: res, err: err}
}
