package server

import (
	"fmt"
)

var port string = "9000"

/* @@@TODO (0xbiel): Port function is bad, ugly and confusing, 
                     need a better approach to it.
                     Disabling it for now.
func Port() {
	if(len(os.Args) == 2) {
		fmt.Println("Please, specify a port.")
	} else {
		port := os.Args[2]
		fmt.Printf("Port is %s\n", port)
		Start(port);
	}
}*/

func Start() {
	fmt.Printf("Starting local server on port %s.\n", port)
}
