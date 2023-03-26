package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"net/http" 
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

    if _, err := os.Stat(storagePath); os.IsNotExist(err) {
        fmt.Println("Directory", storagePath, "does not exist.");
		initStorage(storagePath)
    }
	
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        fmt.Println("Error retrieving file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer file.Close()

    filePath := filepath.Join(storagePath, handler.Filename)
    outFile, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, file)
    if err != nil {
        fmt.Println("Error copying file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}