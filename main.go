package main

import (  
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2"
    "juego2/scenes"
    "juego2/models"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Pong")
    myWindow.Resize(fyne.NewSize(800, 400))
    myWindow.CenterOnScreen()
    init:=scenes.NewInit()
    myWindow.SetContent(init.Content)
    myWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
        
        switch event.Name {
        case fyne.KeyUp:
            if scenes.Player2.Player2Y > 0 {
                scenes.Player2.Player2Y -= 15
            }
        case fyne.KeyDown:
            if scenes.Player2.Player2Y < models.WindowHeight-models.PaddleHeight-45 {
                scenes.Player2.Player2Y += 15
            }
        case fyne.KeyW:
            if scenes.Player1.Player1Y > 0 {
                scenes.Player1.Player1Y -= 15
            }
        case fyne.KeyS:
            if scenes.Player1.Player1Y < models.WindowHeight-models.PaddleHeight-45 {
                scenes.Player1.Player1Y += 15
            }
        }
        scenes.Player2.Player.Move(fyne.NewPos(models.WindowWidth-models.PaddleWidth-15, float32(scenes.Player2.Player2Y)))
        scenes.Player1.Player.Move(fyne.NewPos(float32(scenes.Player1.Player1X), float32(scenes.Player1.Player1Y)))
    })
    
    go func() {
        for newScore := range models.Player2Updates {
            
            if newScore == 3  {
                gameOverScene := scenes.NewGameOverScene(2)
                myWindow.SetContent(gameOverScene.Content)

            }
            
            
        }
    }()

    go func() {
        for newScore := range models.Player1Updates {
            if newScore == 3  {
                
                gameOverScene := scenes.NewGameOverScene(1)
                myWindow.SetContent(gameOverScene.Content)
            }
        }
    }()

    myWindow.ShowAndRun()
}

