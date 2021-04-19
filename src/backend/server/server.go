package server

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

var port string = ":8080"

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Home!")
}

/* If no port number is specified after '-s',
 * start local server on port 8080.
 * If port number is specified after '-s',
 * start local server on given port.*/

func Start() {
	if(len(os.Args) == 2) {
		port = port
	} else {
		port = ":" + os.Args[2]
	}

	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))

}
