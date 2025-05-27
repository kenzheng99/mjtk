package model

import (
	"fmt"
)

type Hand struct {
	Tiles []Tile
	Draw  Tile
	State HandState
}

const HandLength = 13

func CreateHand(handStr string) (Hand, error) {
	if len(handStr) != HandLength*2 {
		return Hand{}, ErrInvalidHand(handStr, nil)
	}

	hand := Hand{}
	hand.Tiles = make([]Tile, 0, HandLength)
	for i := 0; i < len(handStr)-1; i += 2 {
		tileStr := handStr[i : i+2]
		tile, err := CreateTile(tileStr)
		if err != nil {
			return Hand{}, ErrInvalidHand(handStr, err)
		}
		hand.Tiles = append(hand.Tiles, tile)
	}

	return hand, nil
}

func (h Hand) String() string {
	var handStr string
	for _, tile := range h.Tiles {
		handStr += tile.String() + " "
	}
	return fmt.Sprintf("Hand{%s, %s, %s}", handStr[:len(handStr)-1], h.Draw, h.State)
}

func ErrInvalidHand(s string, cause error) error {
	return fmt.Errorf("invalid hand: %q: %w", s, cause)
}
