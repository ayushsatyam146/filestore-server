package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		isFilePresent := r.FormValue("isFilePresent") == "true"
		fileName := r.FormValue("fileName")
		fileHash := r.FormValue("fileHash")
		handleFileUpload(w, r, isFilePresent, fileName, fileHash)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleFileUpload(w http.ResponseWriter, r *http.Request, isFilePresent bool, fileName string, fileHash string) {
	if isFilePresent {
		if _, err := os.Stat(filepath.Join("./uploads", fileName)); err == nil {
			fmt.Fprintf(w, "File already present on the server!")
			return
		} else {
			// file does not exist
			fmt.Fprintf(w, "File content already present on the server!, creating new file")
			srcFile := FileHash[fileHash][0]
			source, err := os.Open(filepath.Join("./uploads", srcFile))
			if err != nil {
				panic(err)
			}
			defer source.Close()
			dst, err := os.Create(filepath.Join("./uploads", fileName))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, source)
			if err != nil {
				panic(err)
			}
			FileHash[fileHash] = append(FileHash[fileHash], fileName)
			return
		}
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join("./uploads", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hash, err := calculateFileHash(filepath.Join("./uploads", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		FileHash[hash] = append(FileHash[hash], fileHeader.Filename)
	}
	message := fmt.Sprintf("%s uploaded successfully!", fileName)
	fmt.Fprintf(w, message)
}
