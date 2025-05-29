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
		SeatWind: 1,
	}

	doraIndicator, err := model.CreateTile("3p")
	if err != nil {
		log.Fatal(err)
	}

	gameState := model.GameState{
		PrevalentWind:  1,
		Round:          1,
		Honba:          0,
		DoraIndicators: []model.Tile{doraIndicator},
	}
	fmt.Println(gameState)

	// hand, err := model.CreateHand("1m2m3m4p0p6p7s8s9sWdWdEwEw")
	// hand, err := model.CreateHand("1p2p3p5s5s6s7sWdWdWdGdGdGd")
	// hand, err := model.CreateHand("WdWdWd3m4m5m5m6m7m7m8m8m8m")
	// hand, err := model.CreateHand("1p1p1p2p2p3p3p4p6p7p8p7p8p")

	// hand, err := model.NewHandWithDraw("1p1p1p2p2p2p3p3p3p4p4p5p5p", "4p")
	// hand, err := model.NewHandWithDraw("1p2p3p7p8p9p1s1s1s9p9p7m8m", "6m")
	// hand, err := model.NewHandWithDraw("1p1p1p9p9p9p9s9s9sWdWdEwEw", "Ew")
	// hand, err := model.NewHandWithDraw("RdRdRdGdGdGdWdWd7p8p9p1p2p", "3p")
	// hand, err := model.NewHandWithDraw("NwNwNw1p2p3p4p5p6p7p8pGdGd", "9p")
	hand, err := model.NewHandWithDraw("1p1p1p2p2p2p3p3p3p4p4p5p5p", "5p")
	if err != nil {
		log.Fatal(err)
	}

	hand.State = handState
	hand.Sort()
	fmt.Println(hand)

	parsedHands := hand.Parse()
	fmt.Println("Num Parses: ", len(parsedHands))
	for _, ph := range parsedHands {
		fmt.Println()
		fmt.Println(ph)
		fmt.Println(ph.WaitTypes())
		fmt.Println("Yaku: ", model.CalculateYaku(ph, gameState))
	}
}
