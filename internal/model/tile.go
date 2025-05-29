package model

import (
	"cmp"
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

func (t Tile) Next() Tile {
	if t.IsHonor() { // honor
		return Tile{}
	}

	if t.Value >= 9 { // terminal
		return Tile{}
	}

	if t.Value == 0 { // aka
		return Tile{
			Suit:  t.Suit,
			Value: 6,
		}
	}

	// regular
	return Tile{
		Suit:  t.Suit,
		Value: t.Value + 1,
	}
}

func (t Tile) NextWrap() Tile {
	if t.IsHonor() {
		if t.Suit == 'd' { // dragons
			return Tile{
				Suit:  t.Suit,
				Value: ((t.Value-5)+1)%3 + 5,
			}
		}
		// winds
		return Tile{
			Suit:  t.Suit,
			Value: (t.Value + 1) % 4,
		}
	}

	if t.Value == 0 { // aka
		return Tile{
			Suit:  t.Suit,
			Value: 6,
		}
	}

	// regular
	return Tile{
		Suit:  t.Suit,
		Value: (t.Value % 9) + 1,
	}
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

// gets tile numeric value. aka returns 5
func (t Tile) FaceValue() int {
	if t.Value == 0 {
		return 5
	}
	return t.Value
}

// equals ignores aka
func (t Tile) Equals(o Tile) bool {
	return t.Suit == o.Suit && t.FaceValue() == o.FaceValue()
}

func (t Tile) IsTerminal() bool {
	return (t.Suit == 'm' || t.Suit == 'p' || t.Suit == 's') && (t.Value == 1 || t.Value == 9)
}

func (t Tile) IsDragon() bool {
	return t.Suit == 'd'

}

func (t Tile) IsWhiteDragon() bool {
	return t.Suit == 'd' && t.Value == 5
}

func (t Tile) IsGreenDragon() bool {
	return t.Suit == 'd' && t.Value == 6
}

func (t Tile) IsRedDragon() bool {
	return t.Suit == 'd' && t.Value == 7
}

func (t Tile) IsWind(windIndex int) bool {
	return t.Suit == 'w' && t.Value == windIndex
}

func ErrInvalidTile(s string) error {
	return fmt.Errorf("invalid tile: %q", s)
}

func CmpTile(a, b Tile) int {
	suitSortOrder := map[byte]int{
		'm': 1,
		'p': 2,
		's': 3,
		'w': 4,
		'd': 5,
	}
	if a.Suit == b.Suit {
		return cmp.Compare(a.FaceValue(), b.FaceValue())
	}
	return cmp.Compare(suitSortOrder[a.Suit], suitSortOrder[b.Suit])
}
