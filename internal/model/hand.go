package model

import (
	"errors"
	"fmt"
	"slices"
)

type Hand struct {
	Tiles    []Tile
	LastDraw Tile
	State    HandState
}

const MaxHandLength = 14

func CreateHand(handStr string) (Hand, error) {
	if len(handStr) > MaxHandLength*2 {
		return Hand{}, ErrInvalidHand(handStr, nil)
	}

	hand := Hand{}
	hand.Tiles = make([]Tile, 0, MaxHandLength)
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
	return fmt.Sprintf("Hand{%s, %s, %s}", handStr[:len(handStr)-1], h.LastDraw, h.State)
}

func (h *Hand) Draw(tileStr string) error {
	if h.Len() >= MaxHandLength {
		return errors.New("hand size exceeded")
	}

	tile, err := CreateTile(tileStr)
	if err != nil {
		return fmt.Errorf("error drawing tile: %w", err)
	}

	h.Tiles = append(h.Tiles, tile)
	h.LastDraw = tile
	return nil
}

func (h Hand) Len() int {
	return len(h.Tiles)
}

func (h *Hand) Sort() {
	slices.SortFunc(h.Tiles, CmpTile)
}

func ErrInvalidHand(s string, cause error) error {
	return fmt.Errorf("invalid hand: %q: %w", s, cause)
}
