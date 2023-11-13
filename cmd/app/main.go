package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/damirsharip/myoffice-test/internal/service"
)

func main() {
	filePath := flag.String("file", "file.txt", "path to file")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	s := service.NewService()
	s.UrlHandler(file)

	return
}
