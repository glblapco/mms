package server

import (
	"fmt"
	"os"
)

var port string = "8080"

/* If no port number is specified after '-s',
 * start local server on port 8080.
 * If port number is specified after '-s',
 * start local server on given port.*/

func Start() {
	if(len(os.Args) == 2) {
		port = port
	} else {
		port = os.Args[2]
	}
		fmt.Printf("Starting local server on port %s.\n", port)
}
