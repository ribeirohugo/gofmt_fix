package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileName       = "log.txt"
	outputFileName = "output.txt"

	command            = "gofmt -s -w"
	separateLineString = ":1: File is not `gofmt`-ed with `-s` (gofmt)"
)

func main() {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)

	var files []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, separateLineString) {
			fields := strings.Split(line, separateLineString)
			files = append(files, fields[0])
		}
	}

	createFile(files)
}

func createFile(lines []string) {
	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := range lines {
		line := fmt.Sprintf("%s %s\n\r", command, lines[i])

		_, err = file.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}
