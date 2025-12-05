package project

import (
	mmap "FDFreader/model"
	"math"
)

func ProjectTo2D(m *mmap.Map, angleX, angleY, zoom float64) [][]mmap.Vertex {
	sinX, cosX := math.Sin(angleX), math.Cos(angleX)
	sinY, cosY := math.Sin(angleY), math.Cos(angleY)

	proj := make([][]mmap.Vertex, m.Height)
	for i := 0; i < m.Height; i++ {
		proj[i] = make([]mmap.Vertex, m.Width)
		for j := 0; j < m.Width; j++ {
			z := m.Data[i][j]

			//rotate around x
			rotY := float64(j)*cosY - z*sinX
			rotZ := float64(j)*sinY - z*cosX

			//rotate around y
			rotX := float64(i)*cosY - rotZ*sinY
			rotatedZ := float64(i)*sinY + rotZ*cosY

			//to 2d
			x2d := (rotX - rotY) * zoom
			y2d := (rotX+rotY)*0.5*zoom - rotatedZ*zoom*0.3

			proj[i][j] = mmap.Vertex{X: x2d, Y: y2d, Z: rotatedZ}
		}
	}
	return proj
}
