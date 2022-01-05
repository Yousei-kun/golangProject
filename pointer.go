package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(gameName string) {
	gamer.Games = append(gamer.Games, gameName)
}

func main() {
	gamer := Gamer{Name: "Ivan"}

	gamer.AddGame("Mobile Legend")
	gamer.AddGame("PUBG Mobile")
	gamer.AddGame("Cookie Run Kingdom")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
