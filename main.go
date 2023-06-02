package main

import (
	"fmt"
	"net/http"
	"os"
)

// Declare the global map variable
var FileHash = make(map[string][]string)

func main() {
	os.Mkdir("./uploads", os.ModePerm)

	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/list", listFilesHandler)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server is running on port 8080.")

}
