package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type email string

type domain struct {
	domain string
	count  int
}

func (e email) String() string {
	return string(e)
}

func (e email) validate() bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, e.String())
	if err != nil {
		return false
	}

	return match
}

func (e email) domain() string {
	parts := strings.Split(e.String(), "@")
	if e.validate() {
		return parts[1]
	}

	return "*invalid"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <csv_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	domainCounts := make(map[string]int)
	rowNum := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if rowNum == 0 {
			rowNum++
			continue
		}

		email := email(record[2])
		d := email.domain()

		domainCounts[d]++
	}

	var domainCountList []domain
	for d, count := range domainCounts {
		domainCountList = append(domainCountList, domain{d, count})
	}

	sort.Slice(domainCountList, func(i, j int) bool {
		return domainCountList[i].domain < domainCountList[j].domain
	})

	for _, dc := range domainCountList {
		fmt.Printf("%s: %d\n", dc.domain, dc.count)
	}
}
