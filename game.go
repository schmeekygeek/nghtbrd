package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Game struct {
  Canvas [canvasY][canvasX]Point
  BirdPosition Point
}

type Point struct {
  X, Y int
  Type string
}

func InitGame() *Game {
  game := new(Game)
  points := [canvasY][canvasX]Point{}
  for i := range points {
    for j := range points[i] {
      points[i][j] = Point{
        X: i,
        Y: j,
        Type: Canvaspixel,
      }
    }
    game.BirdPosition = Point{
    	X:    8,
    	Y:    7,
    	Type: Bird,
    }
  }
  game.Canvas = points
  return game
}

func (game *Game) RunGame() {
  go game.keyboardListen()
  // game loop
  for {
    game.BirdPosition.Y++
    if game.BirdPosition.Y == canvasY {
      fmt.Println("You died")
      os.Exit(0)
      break
    }
    game.render()
    time.Sleep(time.Millisecond * 150)
    clearScreen()
  }
}

func (game *Game) render() {
  game.Canvas[game.BirdPosition.Y][game.BirdPosition.X].Type = Bird
  for i := range game.Canvas {
    for j := range game.Canvas[i] {
      fmt.Print(game.Canvas[i][j].Type)
    }
    fmt.Println()
  }
  game.Canvas[game.BirdPosition.Y][game.BirdPosition.X].Type = Canvaspixel
}

func (game *Game) keyboardListen() {
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    input := string(b)
    if input == Up {
      if game.BirdPosition.Y < 5 {
        game.BirdPosition.Y = 0
      } else {
        game.BirdPosition.Y -= 5
      }
    }
  }
}

func clearScreen() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

const (
  Tower       = "#"
  Bird        = "@"
  Canvaspixel = " "

  // Directions
  Up        = " "

  // Canvas settings
  canvasX   = 100
  canvasY   = 30
)
