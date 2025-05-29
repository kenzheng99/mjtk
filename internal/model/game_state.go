package model

import "fmt"

type GameState struct {
	PrevalentWind int // 1 for East, 2 for South
	Round         int // 1 - 4, e.g. East 4
	SeatWind      int // 1 = East, 2 = South, 3 = West, 4 = North
	Honba         int // number of honba (repeats)
	RiichiSticks  int // number of riichi sticks on table

	Wall           []Tile
	DoraIndicators []Tile
	UraIndicators  []Tile
	NextDrawIndex  int
}

var WindMap = map[int]string{
	1: "East",
	2: "South",
	3: "West",
	4: "North",
}

func (gs GameState) String() string {
	return fmt.Sprintf("GameState: %s %d repeat %d, %d sticks, dora: %s, ura: %s",
		WindMap[gs.PrevalentWind],
		gs.Round,
		gs.Honba,
		gs.RiichiSticks,
		gs.DoraIndicators,
		gs.UraIndicators,
	)
}
