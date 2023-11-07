package main

import (
	"fmt"
	"strings"
)

func pingHopCli(){
	for {
		fmt.Print("REQUEST  > ")
		// read input until newline
		var input string
		fmt.Scanln(&input)
		// check exit condition
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		} else if strings.Trim(input, " ") == "" {
			fmt.Println("Empty request, please try again.")
			continue
		} else {
			// split input into array
			inputArray := strings.SplitN(input, ">", 2)
			ip := inputArray[0]
			content := ""
			if len(inputArray) > 1 {
				content = inputArray[1]
			}
			// send request
			newContent := sendRequest(ip, content)
			// print response
			fmt.Println("RESPONSE > " + newContent)
		}
	}
}