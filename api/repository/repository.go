package repository

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Queue will store the queue type
type Queue struct {
	Domain   string
	Weigth   int
	Priority int
}

// Repository will define the methods to be implemented
type Repository interface {
	Read() []*Queue
}

func (que *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/repository/domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := ""
	var q *Queue
	var queue []*Queue
	domain := ""
	weight := 0
	priority := 0
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			continue
		} else if line == "alpha" || line == "omega" || line == "beta" {
			domain = line
		} else if strings.Contains(line, "weight") {
			weight, _ = strconv.Atoi(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "priority") {
			priority, _ = strconv.Atoi(strings.Split(line, ":")[1])
			q = &Queue{
				Domain:   domain,
				Weigth:   weight,
				Priority: priority,
			}
			queue = append(queue, q)
		}
	}

	return queue
}
