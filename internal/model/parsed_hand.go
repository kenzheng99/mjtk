package model

import "fmt"

type ParsedHand struct {
	Hand

	Pair   TileGroup
	Groups [4]TileGroup
}

// parse a given hand into different winning interpretations
func (h Hand) Parse() []ParsedHand {
	var parses []ParsedHand
	ph := ParsedHand{
		Hand: h,
	}
	parses = append(parses, parseRecursive(ph, 0, 4)...)
	parses = append(parses, parseRecursive(ph, 0, 0)...)

	parsesSet := make(map[string]struct{})
	var deduplicated []ParsedHand
	for _, parse := range parses {
		if _, ok := parsesSet[parse.String()]; !ok {
			deduplicated = append(deduplicated, parse)
			parsesSet[parse.String()] = struct{}{}
		}
	}
	return deduplicated
}

// group 4 is pair
func parseRecursive(ph ParsedHand, index int, group int) []ParsedHand {
	if group < 0 || group > 4 || index < 0 || index > len(ph.Hand.Tiles) {
		return nil
	}

	// add current tile
	tile := ph.Tiles[index]
	if group == 4 {
		if !ph.Pair.CanAddPair(tile) {
			return nil
		}
		ph.Pair.Add(tile)
	} else {
		if !ph.Groups[group].CanAdd(tile) {
			return nil
		}
		ph.Groups[group].Add(tile)
	}

	// base case
	if ph.IsComplete() {
		return []ParsedHand{ph}
	}

	// parse next steps
	var parses []ParsedHand
	parses = append(parses, parseRecursive(ph, index+1, 4)...)

	for i := 0; i < 4; i++ {
		isEmpty := ph.Groups[i].Empty()
		parses = append(parses, parseRecursive(ph, index+1, i)...)
		if isEmpty {
			break
		}
	}
	return parses
}

func (ph ParsedHand) IsComplete() bool {
	if !ph.Pair.IsPair() {
		return false
	}
	for _, tg := range ph.Groups {
		if !tg.IsComplete() {
			return false
		}
	}
	return true
}

func (ph ParsedHand) WaitTypes() WaitType {
	var waits WaitType
	for _, tile := range ph.Pair.Tiles {
		if ph.LastDraw.Equals(tile) {
			waits |= WaitTanki
		}
	}
	for _, group := range ph.Groups {
		for i, tile := range group.Tiles {
			if ph.LastDraw.Equals(tile) {
				waits |= group.WaitTypes(i)
				break
			}
		}
	}
	return waits
}

func (ph ParsedHand) String() string {
	return fmt.Sprintf("%s %s %s %s %s", ph.Pair, ph.Groups[0], ph.Groups[1], ph.Groups[2], ph.Groups[3])
}
