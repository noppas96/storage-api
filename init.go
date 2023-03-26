package main

import (
	"os"
	"fmt"
)

func initStorage(storagePath string) {
	err := os.Mkdir(storagePath, 0775)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created successfully!", storagePath)
}