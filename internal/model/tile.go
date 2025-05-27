package model

import (
	"fmt"
	"strconv"
)

// Tile represents a single mahjong tile.
// The value is 0-9, and the suit m, p, s for regular tiles
// 0 represents Red Fives (akadora)
// For Honor Tiles:
// - Ew (East Wind) is {1, w}
// - Sw (South Wind) is {2, w}
// - Ww (West Wind) is {3, w}
// - Nw (North Wind) is {4, w}
// - Wd (White Dragon) is {5, d}
// - Gd (Green Dragon) is {6, d}
// - Rd (Red Dragon) is {7, d}
type Tile struct {
	Value int
	Suit  byte
}

func CreateTile(tileStr string) (Tile, error) {
	if len(tileStr) != 2 {
		return Tile{}, ErrInvalidTile(tileStr)
	}

	suit := tileStr[1]
	var value int
	if suit == 'm' || suit == 'p' || suit == 's' {
		var err error
		value, err = strconv.Atoi(string(tileStr[0]))
		if err != nil || value < 0 || value > 9 {
			return Tile{}, ErrInvalidTile(tileStr)
		}
	} else if suit == 'w' || suit == 'd' {
		switch tileStr {
		case "Ew":
			value = 1
		case "Sw":
			value = 2
		case "Ww":
			value = 3
		case "Nw":
			value = 4
		case "Wd":
			value = 5
		case "Gd":
			value = 6
		case "Rd":
			value = 7
		default:
			return Tile{}, ErrInvalidTile(tileStr)
		}
	} else {
		return Tile{}, ErrInvalidTile(tileStr)
	}

	return Tile{
		Value: value,
		Suit:  suit,
	}, nil
}

func (t Tile) IsHonor() bool {
	return t.Suit == 'w' || t.Suit == 'd'
}

func (t Tile) String() string {
	var value string
	if t.IsHonor() {
		switch t.Value {
		case 1:
			value = "E"
		case 2:
			value = "S"
		case 3:
			value = "W"
		case 4:
			value = "N"
		case 5:
			value = "W"
		case 6:
			value = "G"
		case 7:
			value = "R"
		default:
			value = "?"
		}
	} else {
		value = strconv.Itoa(t.Value)
	}
	return fmt.Sprintf("%s%c", value, t.Suit)
}

func ErrInvalidTile(s string) error {
	return fmt.Errorf("invalid tile: %q", s)
}
