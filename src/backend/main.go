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
	fmt.Println("  -p: Set local server port.")
	fmt.Println("  -h: Show help.\n")
}

func main() {
	if(len(os.Args) < 2) {
		help()
	} else {
		switch(os.Args[1]) {
			case "-s":
				server.Start()
				break
			case "-p":
				server.Port()
				break
			case "-h":
				help()
				break
			default:
				help()
				break
		}
	}
}
