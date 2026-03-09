package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//Task 1: Read and Process CSV Data
//Write a program that reads this file, processes each line, and prints only the names of students whose grade is above 60.

func ReadFile(File string) error {

	inputFile, err := os.Open(File)
	defer inputFile.Close()

	if err != nil {
		return err
	}

	inputReader := bufio.NewReader(inputFile)

	fmt.Println("Students with passing grades:")
	for {
		line, readError := inputReader.ReadString('\n')
		parts := strings.Split(line, ",")
		name := parts[0]
		grade, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		if readError != nil {
			if readError == io.EOF {
				break
			}
		}

		if grade > 60 {
			fmt.Println("-", name)
		}
	}
	return nil
}
