package mmap

type Map struct {
	Width  int
	Height int
	Data   [][]float64
}

func InitMap(data [][]float64) *Map {
	return &Map{
		Width:  len(data[0]),
		Height: len(data),
		Data:   data,
	}
}
