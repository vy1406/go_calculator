package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReactLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file")
		return nil, errors.New("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Error reading file")
	}

	file.Close()
	return lines, nil
}
