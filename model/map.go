package mmap

type Map struct {
	width  int
	height int
	data   [][]float64
}

func InitMap(width, height int, data [][]float64) *Map {
	return &Map{
		width:  len(data[0]),
		height: len(data),
		data:   data,
	}
}
