package scenes

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "image/color"
    "juego2/models"
    
    
)


var (
    image1  *canvas.Image
    image2  *canvas.Image
    ballCircle   *canvas.Circle
    
    backgroundImage *canvas.Image
)
var balli *models.Ball
var Player1 *models.Player1
var Player2 *models.Player2
var nya *models.Nya
var ovni *models.Ovni
type UI struct {
    Content fyne.CanvasObject
}

func NewInit() *UI {

    ballCircle = canvas.NewCircle(color.White)
    models.Player1Label = widget.NewLabel("Player 1: 0")
    models.Player2Label = widget.NewLabel("Player 2: 0")
    
    image1 = canvas.NewImageFromFile("./assets/player1.png")
    image1.Resize(fyne.NewSize(models.PaddleWidth,models.PaddleHeight))

    image2 = canvas.NewImageFromFile("./assets/player2.png")
    image2.Resize(fyne.NewSize(models.PaddleWidth,models.PaddleHeight))
    
    decorativeNya := canvas.NewImageFromFile("./assets/nyan.png")
    decorativeNya.Resize(fyne.NewSize(100,28))
    
    
    nya = models.NewNya(decorativeNya)

    decorativeOvni := canvas.NewImageFromFile("./assets/ovni.png")
    decorativeOvni.Resize(fyne.NewSize(70,70))
    
    
    ovni = models.NewOvni(decorativeOvni)

    ballCircle.Resize(fyne.NewSize(10, 10))
    backgroundImage = canvas.NewImageFromFile("./assets/background.jpeg")
    backgroundImage.Resize(fyne.NewSize(models.WindowWidth,models.WindowHeight))

    balli = models.NewBall(10,models.WindowWidth/2,models.WindowHeight/2,models.BallDX,models.BallDY,ballCircle)

    Player1 = models.NewPlayer1((
        models.WindowHeight - models.PaddleHeight) / 2,
        0,
        models.PaddleHeight,
        models.PaddleWidth,
        image1)

    Player2 = models.NewPlayer2((
        models.WindowHeight - models.PaddleHeight) / 2,
        models.WindowWidth-models.PaddleWidth-15,
        models.PaddleHeight,
        models.PaddleWidth,
        image2)

    gameObjects := container.NewWithoutLayout(Player1.Player, Player2.Player, ballCircle, nya.Nya, ovni.Ovni)
    scoreObjects := container.NewGridWithColumns(2, models.Player1Label, models.Player2Label)
    backgroundWithGame := container.NewStack(backgroundImage,gameObjects)
    window:=container.NewBorder(scoreObjects, nil, nil, nil, backgroundWithGame)
    
    window.Resize(fyne.NewSize(models.WindowWidth, models.WindowHeight))
    

    go balli.Run(Player1,Player2)
    go nya.MoveNya()
    go ovni.MoveOvni()

    Player2.Player.Move(fyne.NewPos(models.WindowWidth-models.PaddleWidth-15, float32(Player2.Player2Y)))
    Player1.Player.Move(fyne.NewPos(float32(Player1.Player1X), float32(Player1.Player1Y)))
    
    return &UI{
        Content: window,   
        
    }
    
    

}

