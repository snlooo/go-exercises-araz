package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Write a program that reads this file and counts the total number of lines.
//
//Print the line count.

func ReadStory(path string) error {

	inputFile, err := os.Open(path)
	defer inputFile.Close()
	if err != nil {
		return err
	}

	newReader := bufio.NewReader(inputFile)

	count := 0
	for {
		_, readError := newReader.ReadString('\n')
		count++
		if readError != nil {
			if readError == io.EOF {
				break
			}
		}
	}

	fmt.Println("#================================#")
	fmt.Println("Total lines in file: ", count)
	return nil
}
