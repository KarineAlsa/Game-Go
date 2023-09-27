// score.go
package models
var (
    Player1Score int
    Player2Score int
    Player1Updates chan int
    Player2Updates chan int
)
func init() {
    Player1Updates = make(chan int)
    Player2Updates = make(chan int)
}

func SetPlayer1Score(newScore int) {
    
    Player1Score = newScore
    Player1Updates <- newScore
}

func SetPlayer2Score(newScore int) {
 
    Player2Score = newScore

    Player2Updates <- newScore
}

func GetPlayer1Score() int {
   
    return Player1Score
}

func GetPlayer2Score() int {
    
    return Player2Score
}
