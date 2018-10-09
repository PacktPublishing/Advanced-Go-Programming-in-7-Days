package main

import (
	"errors"
	"io"
	"log"
	"reflect"
	"fmt"
)

var SomeError = errors.New("error: description")
var EOFError = errors.New("EOF")

// Implement error interface
type CommandError struct {
	err string
}

func (e CommandError) Error() string {
	return e.err
}

func Avoid() error {
	return &CommandError{"Avoid Command"}
}

func main() {
	// No 2 errors are the same
	if EOFError == io.EOF {
		log.Fatal("Should not happen")
	}

	// prints *errors.errorString so its not ideal for switch statements
	fmt.Println(reflect.TypeOf(SomeError))

	switch SomeError.(type) {
	case error:
		fmt.Println("Its an error")
	}

	switch SomeError {
	case SomeError:
		fmt.Println("Its the same error")
	}

	// Type hint is required for switch
	var invalidCommand error = CommandError{"Invalid Command"}

	switch invalidCommand.(type) {
	case CommandError:
		fmt.Println(invalidCommand)
	}

	err := Avoid()
	// runtime time check
	if err, ok := err.(*CommandError); ok {
		fmt.Println(err)
	}
}
