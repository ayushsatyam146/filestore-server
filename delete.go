package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func deleteSpecifiedFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleSpecifiedFileDeletion(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSpecifiedFileDeletion(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	if filename == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("./uploads", filename)
	hash, err := calculateFileHash(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lerr := os.Remove(filePath)
	if lerr != nil {
		http.Error(w, lerr.Error(), http.StatusInternalServerError)
		return
	}

	hashSlice := FileHash[hash]
	for i, v := range hashSlice {
		if v == filename {
			hashSlice = append(hashSlice[:i], hashSlice[i+1:]...)
			break
		}
	}
	FileHash[hash] = hashSlice
	if len(FileHash[hash]) == 0 {
		delete(FileHash, hash)
	}

	fmt.Fprintf(w, "File %s removed successfully!", filename)
}
