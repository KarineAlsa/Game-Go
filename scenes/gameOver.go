package scenes

import (
    "fyne.io/fyne/v2"
    
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
)

type GameOverScene struct {
	Content fyne.CanvasObject
	
}

func NewGameOverScene(play int) *GameOverScene {
	backgroundOver := canvas.NewImageFromFile("./assets/over1.png")
	if play == 1 {
		backgroundOver = canvas.NewImageFromFile("./assets/over1.png")
    	backgroundOver.Resize(fyne.NewSize(800,400))
	}
	if play == 2 {
		backgroundOver = canvas.NewImageFromFile("./assets/over2.png")
    	backgroundOver.Resize(fyne.NewSize(800,400))
	}
    

	
	
    return &GameOverScene{
        Content: container.NewBorder(
			 nil, nil, nil,nil,backgroundOver,
		),
    }
}
