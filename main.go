// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("GOPATH"))
}
