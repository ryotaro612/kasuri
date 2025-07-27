package main

import (
	"fmt"
	"os"

	"github.com/ryotaro612/ashitaba/internal"
)

func main() {
	_, err := internal.Read(os.Args[1:])
	if err != nil {

	}
	fmt.Println("Hello, World!")
}
