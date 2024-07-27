package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
  "math/rand"
)

type Game struct {
  canvas       [canvasY][canvasX]Point
  birdPosition Point
  hasLost      bool
}

type Point struct {
  X, Y int
  Type string
}

func InitGame() *Game {
  game := new(Game)
  game.hasLost = false
  points := [canvasY][canvasX]Point{}
  for i := range points {
    for j := range points[i] {
      points[i][j] = Point{
        X: i,
        Y: j,
        Type: Canvaspixel,
      }
    }
    game.birdPosition = Point{
    	X:    8,
    	Y:    7,
    	Type: Bird,
    }
  }
  game.canvas = points
  return game
}

func (game *Game) RunGame() {
  go game.inputListen()
  go game.spawnTower()
  // game loop
  for {
    // some game logic
    if game.hasLost {
      fmt.Println("You died")
      os.Exit(0)
    }
    game.birdPosition.Y++
    if game.birdPosition.Y == canvasY {
      game.hasLost = true
    } else if game.canvas[game.birdPosition.Y][game.birdPosition.X].Type == Tower {
      game.hasLost = true
    }
    if game.hasLost {
      fmt.Println("You died")
      os.Exit(0)
    }
    game.tickAndRender()
    time.Sleep(time.Millisecond * 150)
    clearScreen()
  }
}

func (game *Game) spawnTower() {
  for {
    time.Sleep(3 * time.Second)
    towerLengthBottomOffset := rand.Intn(canvasY - 5)
    fmt.Println(towerLengthBottomOffset)

    for i := 0; i < canvasY; i++ {
      game.canvas[i][canvasX - 1] = Point{
      	X:    i,
      	Y:    canvasY,
      	Type: Tower,
      }
    }
    game.canvas[towerLengthBottomOffset][canvasX - 1].Type = Canvaspixel
    game.canvas[towerLengthBottomOffset + 1][canvasX - 1].Type = Canvaspixel
    game.canvas[towerLengthBottomOffset + 2][canvasX - 1].Type = Canvaspixel
  }
}

func (game *Game) tickAndRender() {
  for i := range game.canvas {
    for j := range game.canvas[i] {
      // render
      if i == game.birdPosition.Y && j == game.birdPosition.X {
        fmt.Print(Bird)
      } else {
        fmt.Print(game.canvas[i][j].Type)
      }
      if game.canvas[i][j].Type != Bird && j != 0 {
        game.canvas[i][j - 1] = game.canvas[i][j]
        game.canvas[i][j].Type = Canvaspixel
      }
    }
    fmt.Println()
  }
}

func (game *Game) inputListen() {
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    input := string(b)
    if input == Up {
      if game.birdPosition.Y < 5 {
        game.birdPosition.Y = 0
      } else {
        game.birdPosition.Y -= 5
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
  Tower       = "❒"
  Bird        = "✈"
  Canvaspixel = " "

  // Directions
  Up        = " "

  // Canvas settings
  canvasX   = 50
  canvasY   = 20
)
