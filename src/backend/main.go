package main

import (
	"fmt"
	server "proj/mms/src/backend/server"
	"os"
)

func help() {
	fmt.Println(`
 _____ _____ ___ 
|     |     |_ -|
|_|_|_|_|_|_|___|
	`)
	fmt.Println("Usage: mms [Arguments]")
	fmt.Println("Arguments:")
	fmt.Println("  -s: Start local server.")
	fmt.Println("  -h: Show help.\n")
	fmt.Println("Examples:")
	fmt.Println("  $ mms -s: Starts local server on port 8080.")
	fmt.Println("  $ mms -s 9000: Starts local server on port 9000.\n")
}

func main() {
	if(len(os.Args) < 2) {
		help()
	} else {
		switch(os.Args[1]) {
			case "-s":
				server.Start()
				break
			case "-h":
				help()
				break
			default:
				fmt.Println("Invalid argument.")
				break
		}
	}
}
