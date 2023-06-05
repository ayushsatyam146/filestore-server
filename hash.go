package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func calculateFileHash(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(fileData)
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash, nil
}

func getHashHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		getIsFileHashed(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getIsFileHashed(w http.ResponseWriter, r *http.Request) {

	var input map[string]string
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := make(map[string]bool)
	for key, value := range input {
		_, ok := FileHash[value]
		output[key] = ok
	}
	response, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
