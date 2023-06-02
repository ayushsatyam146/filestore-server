package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleFileListing(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleFileListing(w http.ResponseWriter, r *http.Request) {

	files, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	response := strings.Join(fileNames, "\n")
	fmt.Fprint(w, response)
}
