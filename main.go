package main

import (
	"fmt"
	"net/http"
	"os"
)

// Declare the global map variable
var FileHash = make(map[string][]string)

func main() {
	fmt.Println("Server is running on port 8080.")
	os.Mkdir("./uploads", os.ModePerm)

	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/getHash", getHashHandler)
	http.HandleFunc("/deletefile", deleteSpecifiedFileHandler)
	http.HandleFunc("/list", listFilesHandler)
	http.HandleFunc("/wordcount", wordCountHandler)
	http.HandleFunc("/frequentWords", freqWordsHandler)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
