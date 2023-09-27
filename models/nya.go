
package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Nya struct {
	
	Nya *canvas.Image
}

func NewNya(img *canvas.Image) *Nya{
	return &Nya{
		
		Nya: img,
	}
}

func (n *Nya)MoveNya() {

	n.Nya.Move(fyne.NewPos(-100, -100))
    for {
        border := rand.Intn(4) // 0: izquierda, 1: arriba, 2: abajo

        
        var startX, startY float32
        switch border {
        case 0: 
            startX = -float32(n.Nya.Size().Width)
            startY = rand.Float32() * float32(355)
        case 1: 
			startX = rand.Float32() * float32(800)
			startY = -float32(n.Nya.Size().Height)

        case 2: 
			startX = rand.Float32() * float32(800)
			startY = float32(355)
        
        }
      
        directionX := rand.Float32() * 2 - 1 
        directionY := rand.Float32() * 2 - 1 

        for startX >= -float32(n.Nya.Size().Width) && startX <= float32(800) &&
            startY >= -float32(n.Nya.Size().Height+100) && startY <= float32(355) {

            n.Nya.Move(fyne.NewPos(startX, startY))
            startX += directionX
            startY += directionY
            time.Sleep(10 * time.Millisecond) 
        }
		
        
        time.Sleep(time.Duration(rand.Intn(5000)+1000) * time.Millisecond)
    }
}

