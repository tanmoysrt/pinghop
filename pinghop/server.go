package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"
)

func pingHopServer() {
	// create a http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		bodyString := strings.Trim(string(body), " ")
		if bodyString == "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("SUCCESS"))
		} else {
			// split the body string by > on first occurence
			splittedBodies := strings.SplitN(bodyString, ">", 2)
			systemIp := splittedBodies[0]
			content := ""
			if len(splittedBodies) > 1 {
				content = splittedBodies[1]
			}
			// send request
			newContent := sendRequest(systemIp, content)
			// write response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(newContent))
		}
	})
	http.ListenAndServe(":3333", nil)
}

// private functioon
func sendRequest(ip string, content string) string {
	// Create an HTTP client with a 20-second timeout
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	// create request
	reader := bytes.NewReader([]byte(content))
	request, err := http.NewRequest("POST", "http://"+ip+":3333", reader)
	if err != nil {
		panic("error crafting request")
	}
	// send request
	response, err := client.Do(request)
	/// FAILED from next hop
	if err != nil {
		newContent := ip + " ( FAILED )"
		return newContent
	} else {
		/// SUCCESS from next hop
		defer response.Body.Close()
		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic("error reading response body")
		}
		newContent := ip + " <-- " + string(body)
		return newContent
	}
}
