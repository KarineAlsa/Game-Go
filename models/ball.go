
package models

import (
	"time"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)
const (
    WindowWidth  = 800
    WindowHeight = 400
    PaddleWidth  = 15
    PaddleHeight = 70
    
)
var (
    player1Y = (WindowHeight - PaddleHeight) / 2
    player2Y = (WindowHeight - PaddleHeight) / 2
    BallX    = WindowWidth / 2
    BallY    = (WindowHeight / 2)
    BallDX   = -1
    BallDY   = -1
    Score1   = 0
    Score2   = 0
	Player1Label *widget.Label
    Player2Label *widget.Label
)
type Ball struct {
	
	ballX, ballY, ballSize int
	ballDX, ballDY int
	ball *canvas.Circle
	status bool
}

func NewBall(ballSize int,ballx int, bally int,balldx int, balldy int, img *canvas.Circle) *Ball{
	return &Ball{
		ballSize: ballSize,
		ballX: ballx,
		ballY: bally,
		ballDX: balldx,
		ballDY: balldy,
		ball: img,
	}
}

func (b *Ball) Run(p1 *Player1,p2 *Player2) {
	b.status=true
	for b.status{
		b.ballX += b.ballDX
        b.ballY += b.ballDY

        if b.ballY <= 0 || b.ballY >= (WindowHeight-45)-b.ballSize {
            b.ballDY *= -1
        }

        if b.ballX <= PaddleWidth && b.ballY >= p1.Player1Y && b.ballY <= p1.Player1Y+PaddleHeight {
        	b.ballDX *= -1
        }

        if b.ballX >= WindowWidth-PaddleWidth-b.ballSize-15 && b.ballY >= p2.Player2Y && b.ballY <= p2.Player2Y+PaddleHeight {
            b.ballDX *= -1
        }

        if b.ballX < 0 {
            b.ballX = WindowWidth / 2
            b.ballY = WindowHeight / 2
            b.ballDX = -b.ballDX
            Score2++
			SetPlayer2Score(Score2)
        } else if b.ballX > WindowWidth {
            b.ballX = WindowWidth / 2
            b.ballY = WindowHeight / 2
            b.ballDX = -b.ballDX
            Score1++
			SetPlayer1Score(Score1)
        }
		b.ball.Move(fyne.NewPos(float32(b.ballX), float32(b.ballY)))
		Player1Label.SetText(fmt.Sprintf("Player 1: %d", Score1))
        Player2Label.SetText(fmt.Sprintf("Player 2: %d", Score2))
		
		if Score1 ==3|| Score2 ==3{
			b.status=false
			
		}
		
        time.Sleep(6 * time.Millisecond)
	}	
}
