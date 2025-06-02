package scorer

import (
	"errors"

	"github.com/kenzheng99/mjtk/internal/model"
)

type DoraCounts struct {
	Dora int
	Ura  int
	Aka  int
}

func (dc DoraCounts) Total() int {
	return dc.Dora + dc.Ura + dc.Aka
}

func ScoreDora(ph model.ParsedHand, gs model.GameState) (DoraCounts, error) {
	if len(gs.DoraIndicators) == 0 {
		return DoraCounts{}, errors.New("no dora indicator found in gamestate")
	}
	var dora, ura []model.Tile
	var doraCounts DoraCounts

	for _, doraIndicator := range gs.DoraIndicators {
		dora = append(dora, doraIndicator.NextWrap())
	}

	for _, uraIndicator := range gs.UraIndicators {
		ura = append(ura, uraIndicator.NextWrap())
	}

	for _, tile := range ph.Tiles {
		if tile.IsAka() {
			doraCounts.Aka++
			tile = tile.RemoveAka()
		}

		if hasTile(dora, tile) {
			doraCounts.Dora++
		}

		if hasTile(ura, tile) {
			doraCounts.Ura++
		}
	}
	return doraCounts, nil
}

func hasTile(tiles []model.Tile, target model.Tile) bool {
	for _, tile := range tiles {
		if target == tile {
			return true
		}
	}
	return false
}
