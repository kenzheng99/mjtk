package main

import (
	"fmt"
	"log"

	"github.com/kenzheng99/mjtk/internal/model"
)

func main() {
	fmt.Println("start mjtk!")

	handState := model.HandState{
		IsDealer: true,
		IsRiichi: true,
		IsTsumo:  true,
	}

	// hand, err := model.CreateHand("1m2m3m4p0p6p7s8s9sWdWdEwEw")
	// hand, err := model.CreateHand("1p1p1p2p2p2p3p3p3p4p4p5p5p")
	// hand, err := model.CreateHand("1p2p3p5s5s6s7sWdWdWdGdGdGd")
	// hand, err := model.CreateHand("WdWdWd3m4m5m5m6m7m7m8m8m8m")
	hand, err := model.CreateHand("1p1p1p2p2p3p3p4p6p7p8p7p8p")
	if err != nil {
		log.Fatal(err)
	}
	err = hand.Draw("9p")
	if err != nil {
		log.Fatal(err)
	}
	hand.State = handState
	hand.Sort()
	fmt.Println(hand)

	parsedHands := hand.Parse()
	fmt.Println("Parsed count: ", len(parsedHands))
	for _, ph := range parsedHands {
		fmt.Println(ph)
		fmt.Println(ph.WaitTypes())
		fmt.Println(ph.HandYakuType())
	}
}
