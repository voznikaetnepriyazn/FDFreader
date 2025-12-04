package input

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	angleX float64
	angleY float64
	zoom   float64
}

func InitInput(anglex, angleY, zoom float64) *Input {
	return &Input{
		angleX: 0.7,
		angleY: 0.7,
		zoom:   20.0,
	}
}

func RotateX(current, delta float64) float64 {
	return current + delta
}

func RotateY(current, delta float64) float64 {
	return current + delta
}

func ScaleZoom(current, factor float64) float64 {
	return current * factor
}

func (i *Input) Update() {
	const delta = 0.05
	const zoomFactor = 1.05

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		i.angleY = RotateY(i.angleY, -delta)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		i.angleY = RotateY(i.angleY, delta)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		i.angleX = RotateX(i.angleX, delta)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		i.angleX = RotateX(i.angleX, -delta)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEqual) || ebiten.IsKeyPressed(ebiten.KeyNumpadAdd) {
		i.zoom = ScaleZoom(i.zoom, zoomFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyMinus) || ebiten.IsKeyPressed(ebiten.KeyNumpadSubtract) {
		i.zoom = ScaleZoom(i.zoom, 1/zoomFactor)
	}
}
