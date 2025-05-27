package main

import (
	"fmt"
	"log"

	"github.com/kenzheng99/mjtk/internal/model"
)

func main() {
	fmt.Println("start mjtk!")
	tile, err := model.CreateTile("Wd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tile)

	handState := model.HandState{
		IsDealer: true,
		IsRiichi: true,
		IsTsumo:  true,
	}
	fmt.Println(handState)

	hand, err := model.CreateHand("1p2p3p4s5s6s7m8m9mWdWdEwEw")
	if err != nil {
		log.Fatal(err)
	}
	hand.Draw = tile
	hand.State = handState
	fmt.Println(hand)

}
