package models
import (
	
	"fyne.io/fyne/v2/canvas"
	
)

type Player2 struct {
	Player2X, Player2Y int
	PaddleHeight, PaddleWidth int
	Player *canvas.Image
}

func NewPlayer2(player2y int, player2x int, paddleHeight int, paddleWidth int, img *canvas.Image) *Player2{
	return &Player2{
		Player2X: player2x,
		Player2Y: player2y,
		PaddleHeight: paddleHeight,
		PaddleWidth: paddleWidth,
		Player: img,
	}
}
