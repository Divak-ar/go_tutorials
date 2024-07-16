package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"

	"golang.org/x/image/font/basicfont"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// go mod tidy to install the game engine - ebiten and other dependencies

const (
	screenWidth = 640
	screenHeight = 480
	ballSpeed = 3
	paddleSpeed = 6
)

// These are pixels value 

type Object struct{
	X, Y, W, H int
}

type Ball struct{
	Object
	dxdt int //x velocity 
	dydt int //y velocity
}

type Paddle struct{
	Object
}

type Game struct{
	paddle Paddle
	ball Ball
	score int
	highScore int
}

func main(){
	ebiten.SetWindowTitle("Pong game using Ebit-engine")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	paddle := Paddle{
		Object: Object{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
			},
	}
	ball := Ball{
		Object: Object{
			X: 0,
			Y: 0,
			W: 15,
			H: 15,
			},
		dxdt: ballSpeed,
		dydt: ballSpeed,
	}
	g := &Game{
		paddle: paddle,
		ball: ball,
	}

	err := ebiten.RunGame(g)
	if err != nil{
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int){
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	// we are drawing our paddle and ball object on the screen
	vector.DrawFilledRect(screen, 
		float32(g.paddle.X), float32(g.paddle.Y), 
		float32(g.paddle.W), float32(g.paddle.H), 
		color.White, false,
	)
	vector.DrawFilledRect(screen, 
		float32(g.ball.X), float32(g.ball.Y), 
		float32(g.ball.W), float32(g.ball.H), 
		color.White, false,
	)

	
	scoreStr := "Score: " + fmt.Sprint(g.score)
	// displaying/drawinf score on the screen 
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 10, color.White)

	highScoreStr := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 30, color.White)

}

func (g *Game) Update() error {	
	g.paddle.MoveOnKeyPress()
	g.ball.Move()
	g.CollideWithWall()
	g.CollideWithPaddle()
	return nil
}
func (p *Paddle) MoveOnKeyPress() {
	// if down key is pressed we move our paddle downward by increasing y (0->top, 200->bottom)
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= paddleSpeed
	}
}

func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
}

func (g *Game) Reset() {
	g.ball.X = 0
	g.ball.Y = 0

	g.score = 0
}

func (g *Game) CollideWithWall() {

	// Right wall causes a game over
	if g.ball.X >= screenWidth {
		g.Reset()
	}else if g.ball.X <= 0{
		// ball goes in positive x direction ->
		g.ball.dxdt = ballSpeed
	}else if g.ball.Y <= 0 {
		// ball goes downward 0->200px 
		g.ball.dydt = ballSpeed
	}else if g.ball.Y >= screenHeight {
		// ball goes in positive y direction | (upwards)
		g.ball.dydt = -ballSpeed
	}
}

func (g *Game) CollideWithPaddle() {
	// if ball is in touch with paddle
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y + g.paddle.H {
		// sending the ball simulatig a hit with paddle
		g.ball.dxdt = -g.ball.dxdt
		g.score++
		//updating highscore
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}

