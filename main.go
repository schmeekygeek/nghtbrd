package main

func main() {
  game := InitGame()
  forever := make(chan(bool))
  game.RunGame()
  <-forever
}
