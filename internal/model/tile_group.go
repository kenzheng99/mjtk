package model

import (
	"slices"
)

type TileGroup struct {
	Tiles  []Tile
	isOpen bool
}

func (tg TileGroup) Sort() {
	slices.SortFunc(tg.Tiles, CmpTile)
}

// -------------------
// boolean checks
// -------------------
func (tg TileGroup) IsSequence() bool {
	if len(tg.Tiles) != 3 {
		return false
	}
	tg.Sort()
	return tg.Tiles[0].Next().Equals(tg.Tiles[1]) && tg.Tiles[1].Next().Equals(tg.Tiles[2])
}

func (tg TileGroup) IsTriplet() bool {
	if len(tg.Tiles) != 3 {
		return false
	}
	tg.Sort()
	return tg.Tiles[0].Equals(tg.Tiles[1]) && tg.Tiles[1].Equals(tg.Tiles[2])
}

func (tg TileGroup) IsPair() bool {
	if len(tg.Tiles) != 2 {
		return false
	}
	return tg.Tiles[0].Equals(tg.Tiles[1])
}

func (tg TileGroup) IsQuad() bool {
	if len(tg.Tiles) != 4 {
		return false
	}
	return tg.Tiles[0].Equals(tg.Tiles[1]) && tg.Tiles[1].Equals(tg.Tiles[2]) && tg.Tiles[2].Equals(tg.Tiles[3])
}

func (tg TileGroup) IsComplete() bool {
	return tg.IsSequence() || tg.IsTriplet() || tg.IsQuad()
}

// ---------------
// adding
// ---------------
func (tg TileGroup) CanAdd(t Tile) bool {
	if tg.Empty() {
		return true
	}
	tg.Sort()
	if tg.Len() == 1 {
		return tg.Tiles[0].Equals(t) || tg.Tiles[0].Next().Equals(t)
	}
	if tg.Len() == 2 {
		canAddTriplet := tg.Tiles[0].Equals(t) && tg.Tiles[1].Equals(t)
		canAddSequence := tg.Tiles[0].Next().Equals(tg.Tiles[1]) && tg.Tiles[1].Next().Equals(t)
		return canAddTriplet || canAddSequence
	}
	return false
}
func (tg TileGroup) CanAddPair(t Tile) bool {
	if tg.Empty() {
		return true
	}
	if tg.Len() == 1 {
		return tg.Tiles[0].Equals(t)
	}
	return false
}

func (tg *TileGroup) Add(t Tile) {
	tg.Tiles = append(tg.Tiles, t)
}

// ---------
// utility
// ---------
func (tg TileGroup) Len() int {
	return len(tg.Tiles)
}

func (tg TileGroup) Empty() bool {
	return len(tg.Tiles) == 0
}

func (tg TileGroup) String() string {
	str := ""
	for _, t := range tg.Tiles {
		str += t.String()
	}
	return str
}
