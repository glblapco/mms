package server

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"regexp"
)

var port string = ":8080"
var pattern *regexp.Regexp

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Home!")
}

/* Using regexp to match trials of Arbitrary
 * Command Execution. If string matches,
 * the program prints a message and exits,
 * so the command does not get interpreted
 * as a port argument.
 */

func filterArg(str string) string {
	pattern = regexp.MustCompile(`[^\w@%+=:,./-]`)

	if(pattern.MatchString(str)) {
		fmt.Println("Easy bruh! No commands here!")
		os.Exit(1)
	}
	return str
}

/* If no port number is specified after '-s',
 * start local server on port 8080.
 * If port number is specified after '-s',
 * start local server on given port.*/

func Start() {
	if(len(os.Args) == 2) {
		port = port
	} else {
		filterArg(os.Args[2])
		port = ":" + os.Args[2]
	}
	fmt.Printf("Starting server at localhost%s\n", port)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}
