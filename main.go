package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getServerIp(w http.ResponseWriter, r *http.Request) {

	url := "https://ipinfo.io/json"

	timeoutSecs := 10

	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)
	defer cancel()

	// NOTE provide an http request method
	req, err := http.NewRequestWithContext(context, "GET", url, nil)
	if err != nil {
		http.Error(w, "[ERROR]:  Timed out", http.StatusInternalServerError)
		return
	}

	// NOTE provide connection
	var httpClient = &http.Client{}

	// NOTE create httpClient request
	resp, err := httpClient.Do(req)

	if err != nil {

		http.Error(w, "[ERROR]:  Timed out", http.StatusInternalServerError)
		return
	}

	// NOTE check status code
	if resp.StatusCode != 200 {
		log.Printf("[ERROR]: %d Request to %s", resp.StatusCode, url)

		fmt.Fprintf(w, "[ERROR]: %d Request to %s", resp.StatusCode, url)
		return
	}

	defer resp.Body.Close()

	// NOTE read response
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("%s\n", err)
	}

	fmt.Println(string(body))
	fmt.Fprint(w, string(body))

	log.Printf("data fetched successfully\n")
}

func main() {

	var portNumber string = "8080"

	http.HandleFunc("/", getServerIp)
	log.Printf("listening on port %s", portNumber)
	http.ListenAndServe("localhost"+":"+portNumber, nil)

}
