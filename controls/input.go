package input

import (
	mmap "FDFreader/model"
	project "FDFreader/projection"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Input struct {
	Mmap   *mmap.Map
	angleX float64
	angleY float64
	zoom   float64
}

func InitInput(Mmap *mmap.Map, anglex, angleY, zoom float64) *Input {
	return &Input{
		Mmap:   Mmap,
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

func (i *Input) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	proj := project.ProjectTo2D(i.Mmap, i.angleX, i.angleY, i.zoom)

	//height
	for y := 0; y < len(proj); y++ {
		for x := 0; x < len(proj[y])-1; x++ {
			v1 := proj[y][x]
			v2 := proj[y][x+1]
			ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, color.White)
		}
	}

	//width
	for x := 0; x < len(proj[0]); x++ {
		for y := 0; y < len(proj)-1; y++ {
			v1 := proj[y][x]
			v2 := proj[y+1][x]
			ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, color.White)
		}
	}
}

func (i *Input) Layout(width, height int) (int, int) {
	return 1024, 768
}
