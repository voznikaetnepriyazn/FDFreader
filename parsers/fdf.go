package fdf

import (
	"os"
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

}

//defer file.Close()
