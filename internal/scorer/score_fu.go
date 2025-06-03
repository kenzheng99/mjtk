package scorer

import "github.com/kenzheng99/mjtk/internal/model"

func ScoreFu(ph model.ParsedHand, gs model.GameState) int {
	fu := 20
	fu += tripletFu(ph)
	fu += pairFu(ph, gs)
	fu += waitFu(ph)

	if !ph.State.IsOpen && !ph.State.IsTsumo { // closed ron
		fu += 10
	} else if !ph.State.IsOpen && fu == 20 { // closed tsumo pinfu
		return fu
	} else if ph.State.IsTsumo { // closed or open tsumo
		fu += 2
	} else if !ph.State.IsTsumo && fu == 20 { // open "pinfu"
		fu += 2
	}
	return roundUpToNearest10(fu)
}

func tripletFu(ph model.ParsedHand) int {
	fu := 0
	for _, group := range ph.Groups {
		if !group.IsTriplet() {
			continue
		}

		tripletFu := 2
		if !group.IsOpen {
			tripletFu *= 2
		}
		if group.HasTerminal() || group.HasHonor() {
			tripletFu *= 2
		}
		if len(group.Tiles) == 4 {
			tripletFu *= 4
		}
		fu += tripletFu
	}
	return fu
}

func pairFu(ph model.ParsedHand, gs model.GameState) int {
	pairTile := ph.Pair.Tiles[0]
	fu := 0
	if pairTile.IsWind(gs.PrevalentWind) {
		fu += 2
	}
	if pairTile.IsWind(ph.State.SeatWind) {
		fu += 2
	}
	if pairTile.IsWhiteDragon() ||
		pairTile.IsGreenDragon() ||
		pairTile.IsRedDragon() {
		fu += 2
	}
	return fu
}

func waitFu(ph model.ParsedHand) int {
	waits := ph.WaitTypes()
	if waits.Has(model.WaitKanchan) ||
		waits.Has(model.WaitPenchan) ||
		waits.Has(model.WaitTanki) {
		return 2
	}
	return 0
}

func roundUpToNearest10(n int) int {
	return ((n + 9) / 10) * 10
}
