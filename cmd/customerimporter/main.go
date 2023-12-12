package main

import (
	"fmt"
	"os"

	"github.com/zys-gd/go_test_teamwork/pkg/customerimporter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/customerimporter/main.go <csv_file>")
		os.Exit(1)
	}

	customerimporter.FromCsv(os.Args[1])
}
