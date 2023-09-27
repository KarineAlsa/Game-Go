package models
import (
	"fyne.io/fyne/v2/canvas"
)
type Player1 struct {
	Player1X, Player1Y int
	PaddleHeight, PaddleWidth int
	Player *canvas.Image
}

func NewPlayer1(player1y int, player1x int, paddleHeight int, paddleWidth int, img *canvas.Image) *Player1{
	return &Player1{
		Player1X: player1x,
		Player1Y: player1y,
		PaddleHeight: paddleHeight,
		PaddleWidth: paddleWidth,
		Player: img,
	}
}
