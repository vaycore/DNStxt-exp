// Package utils Created by vaycore on 2021-12-31.
package utils

import (
	"bufio"
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func FileReadingLines(filename string) []string {
	var result []string
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return result
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); IsNotEmpty(line) {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return result
	}
	return result
}

func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func IsEmpty(value string) bool {
	return len(value) == 0 || len(strings.TrimSpace(value)) == 0
}
