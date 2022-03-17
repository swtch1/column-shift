package main

import (
	"fmt"
)

func info(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func fatal(err error, msg string, args ...interface{}) {
	m := fmt.Sprintf(msg, args...)
	fmt.Printf("ERROR: %s: %s", m, err)
	enterToContinue()
}
