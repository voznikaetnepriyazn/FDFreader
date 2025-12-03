package fdf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MustLoad(nameOfFile string) ([][]float64, error) {
	if _, err := os.Stat(nameOfFile); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(nameOfFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		slice := make([]float64, len(fields))

		for i, j := range fields {
			val, err := strconv.ParseFloat(j, 64)
			if err != nil {
				fmt.Printf("error of parse from string into float")
			}
			slice[i] = val
		}
		matrix = append(matrix, slice)
	}

	return matrix, nil
}
