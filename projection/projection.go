package project

import (
	"math"
)

func ProjectTo2D(m *Map, angleX, angleY, zoom float64) [][]Vertex {
	sinX, cosX := math.Sin(angleX), math.Cos(angleX)
	sinY, cosY := math.Sin(angleY), math.Cos(angleY)

	proj := make([][]Vertex, m.height)
	for i := 0; i < m.height; i++ {
		proj[i] = make([]Vertex, m.width)
		for j := 0; j < m.width; j++ {
			z := m.data[i][j]

			//rotate around x
			rotY := float64(j)*cosY - z*sinX
			rotZ := float64(j)*sinY - z*cosX

			//rotate around y
			rotX := float64(i)*cosY - rotZ*sinY
			rotatedZ := float64(i)*sinY + rotZ*cosY

			//to 2d
			x2d := (rotX - rotY) * zoom
			y2d := (rotX+rotY)*0.5*zoom - rotatedZ*zoom*0.3

			proj[i][j] = Vertex{x: x2d, y: y2d, z: rotatedZ}
		}
	}
	return proj
}
