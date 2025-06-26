package scorer

import (
	"errors"
	"fmt"

	"github.com/kenzheng99/mjtk/internal/model"
)

var yakuHanCommon = map[model.YakuType]int{
	model.Haitei:  1,
	model.Houtei:  1,
	model.Rinshan: 1,
	model.Chankan: 1,

	model.PrevalentWind: 1,
	model.SeatWind:      1,
	model.WhiteDragon:   1,
	model.GreenDragon:   1,
	model.RedDragon:     1,
	model.Tanyao:        1,

	model.Toitoi:     2,
	model.Sanankou:   2,
	model.Honroutou:  2,
	model.Shousangen: 2,
}

var yakuHanClosed = map[model.YakuType]int{
	model.Riichi:   1,
	model.Tsumo:    1,
	model.Ippatsu:  1,
	model.Pinfu:    1,
	model.Iipeikou: 1,

	model.Chanta:       2,
	model.Sanshoku:     2,
	model.Ittsu:        2,
	model.DoubleRiichi: 2,
	model.Chiitoi:      2,

	model.Honitsu:    3,
	model.Junchan:    3,
	model.Ryanpeikou: 3,

	model.Chinitsu: 6,
}

var yakuHanOpen = map[model.YakuType]int{
	model.Chanta:   1,
	model.Sanshoku: 1,
	model.Ittsu:    1,

	model.Honitsu: 2,
	model.Junchan: 2,

	model.Chinitsu: 5,
}

// 1 for regular yakuman, 2 for double yakuman
var yakumanHan = map[model.YakuType]int{
	model.Kokushi:     13,
	model.Suuankou:    13,
	model.Daisangen:   13,
	model.Shousuushii: 13,
	model.Tsuuiisou:   13,
	model.Chinroutou:  13,
	model.Ryuuiisou:   13,
	model.Chuuren:     13,
	model.Suukantsu:   13,
	model.Tenhou:      13,
	model.Chiihou:     13,

	model.SuuankouTanki: 26,
	model.Kokushi13:     26,
	model.Daisuushii:    26,
}

func ScoreYaku(ph model.ParsedHand, yakus []model.YakuType) (int, error) {
	if len(yakus) == 0 {
		return 0, errors.New("yaku list is empty")
	}

	isYakuman := false
	if _, ok := yakumanHan[yakus[0]]; ok {
		isYakuman = true
	}

	if isYakuman {
		return scoreYakuman(yakus)
	}

	if ph.State.IsOpen {
		return scoreOpen(yakus)
	}

	return scoreClosed(yakus)
}

func scoreYakuman(yakus []model.YakuType) (int, error) {
	score := 0
	for _, yaku := range yakus {
		yakuScore, ok := yakumanHan[yaku]
		if !ok {
			return 0, fmt.Errorf("got invalid yakuman: %q", yaku)
		}
		score += yakuScore
	}
	return score, nil
}

func scoreOpen(yakus []model.YakuType) (int, error) {
	score := 0
	for _, yaku := range yakus {
		yakuScore, ok := yakuHanCommon[yaku]
		if !ok {
			yakuScore, ok = yakuHanOpen[yaku]
			if !ok {
				return 0, fmt.Errorf("got invalid yaku for open hand: %q", yaku)
			}
		}
		score += yakuScore
	}
	return score, nil
}

func scoreClosed(yakus []model.YakuType) (int, error) {
	score := 0
	for _, yaku := range yakus {
		yakuScore, ok := yakuHanCommon[yaku]
		if !ok {
			yakuScore, ok = yakuHanClosed[yaku]
			if !ok {
				return 0, fmt.Errorf("got invalid yaku for closed hand: %q", yaku)
			}
		}
		score += yakuScore
	}
	if score >= 13 {
		score = 13 // kazoe yakuman
	}
	return score, nil
}
