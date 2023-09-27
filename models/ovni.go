
package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Ovni struct {
	
	Ovni *canvas.Image
}

func NewOvni(img *canvas.Image) *Ovni{
	return &Ovni{
		
		Ovni: img,
	}
}

func (o *Ovni)MoveOvni() {

	o.Ovni.Move(fyne.NewPos(-100, -100))
    for {
        border := rand.Intn(3) // 0: izquierda, 1: arriba, 2: abajo
      
        var startX, startY float32
        switch border {
        case 0: 
            startX = -float32(o.Ovni.Size().Width)
            startY = rand.Float32() * float32(355)
        case 1: 
			startX = rand.Float32() * float32(800)
			startY = -float32(o.Ovni.Size().Height)

        case 2: 
			startX = rand.Float32() * float32(800)
			startY = float32(355)    
        }

        directionX := rand.Float32() * 2 - 1 // Valor entre -1 y 1
        directionY := rand.Float32() * 2 - 1 // Valor entre -1 y 1

        secs:=time.Duration(rand.Float32() * float32(10))
       
        for startX >= -float32(o.Ovni.Size().Width) && startX <= float32(800) &&
            startY >= -float32(o.Ovni.Size().Height+100) && startY <= float32(355) {

            o.Ovni.Move(fyne.NewPos(startX, startY))
            
			startX += directionX
			startY += directionY
			
            time.Sleep(secs * time.Millisecond) 
        }
		
        time.Sleep(time.Duration(rand.Intn(5000)+1000) * time.Millisecond)
    }
}

