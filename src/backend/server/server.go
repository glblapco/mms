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

/* @@@(Biel A. P.): Using regexp to match trials 
 * of Arbitrary Command Execution. If string 
 * matches, the program prints a message and 
 * exits, so the command does not get 
 * interpreted as a port argument.
 */

func filterArg(str string) string {
	pattern = regexp.MustCompile(`[^\w@%+=:,./-]`)

	if(pattern.MatchString(str)) {
		fmt.Println("No commands here!")
		os.Exit(1)
	}
	return str
}

func home(rw http.ResponseWriter, req *http.Request) {
	if(req.URL.Path != "/") {
		http.Error(rw, "Page Not Found.", http.StatusNotFound)
		return
	}

	if (req.Method != "GET") {
		http.Error(rw, "Not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("Home!")
}

/* @@@(Biel A. P.): If no port number is 
 * specified after '-s', start local 
 * server on port 8080. If port number 
 * is specified after '-s', start local 
 * server on given port.*/

func Start() {
	if(len(os.Args) == 2) {
		port = port
	} else {
		port = ":" + filterArg(os.Args[2])
	}
	fmt.Printf("Starting server at localhost%s\n", port)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}
