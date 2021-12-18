package main

import (
	"fmt"

	"github.com/DennisPing/learn-ebiten/controller"
	"github.com/DennisPing/learn-ebiten/model"
	"github.com/DennisPing/learn-ebiten/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	GameName string
}

func NewGame() *Game {
	return &Game{GameName: "Example Game"}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (game *Game) Layout(width, height int) (int, int) {
	return 1280, 962
}

func main() {
	fmt.Println("Hello, from main!")

	// Instantiate a new ModelStruct pointer
	var m = &model.ModelStruct{Feature: "This is a feature"}
	fmt.Printf("Our Model Feature looks like this '%s'\n", m.Feature)
	view.ViewFunction()
	controller.ControllerFunction()

	var g = NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(g.GameName)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}

}
