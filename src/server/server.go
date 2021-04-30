package server

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"regexp"
//	"io/ioutil"
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

func serve(rw http.ResponseWriter, req *http.Request) {
	if (req.Method != "GET") {
		http.Error(rw, "Not supported.", http.StatusNotFound)
		return
	}
	http.ServeFile(rw, req, req.URL.Path[1:])
}

func newTask(rw http.ResponseWriter, req *http.Request) {
	//fmt.Println("Form.")
	company := req.FormValue("Company")
	task := req.FormValue("Task")
	deadline := req.FormValue("Deadline")

	fmt.Fprintf(rw, "%s\n", company)
	fmt.Fprintf(rw, "%s\n", task)
	fmt.Fprintf(rw, "%s\n", deadline)

	data := "<p><b>" + company + "</b></p>" + "<p>" + task + "</p>" + "<p>" + deadline + "</p>"

	f, err := os.OpenFile("tasks.html", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if(err != nil) {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

//@@@TODO (Biel A. P.): After storing the data, redirect to Home Page again.

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
	http.HandleFunc("/", serve)
	http.HandleFunc("/new", newTask)
	log.Fatal(http.ListenAndServe(port, nil))
}
