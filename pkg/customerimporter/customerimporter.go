package customerimporter

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

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (e email) isValid() bool {
	return emailPattern.MatchString(e.String())
}

func (e email) domain() string {
	parts := strings.Split(e.String(), "@")
	if e.isValid() {
		return parts[1]
	}

	return "*invalid"
}

func FromCsv(csvpath string) {
	file, err := os.Open(csvpath)
	if err != nil {
		fmt.Println("Cannot opent the file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	domainCounts := make(map[string]int)

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		domainCounts[email(record[2]).domain()]++
	}

	domainCountList := make([]domain, 0, len(domainCounts))
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
