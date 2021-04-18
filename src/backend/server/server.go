package server

import (
	"fmt"
	"os"
)

func Port() {
	if(len(os.Args) == 2) {
		fmt.Println("Please, specify a port.")
	} else {
		port := os.Args[2]
		fmt.Printf("Port is %s\n", port)
}

func Start() {
	fmt.Println("Start!")
}
