package scorer

import "fmt"

type ScoreType int

const (
	ScoreRegular ScoreType = iota
	ScoreMangan
	ScoreHaneman
	ScoreBaiman
	ScoreSanbaiman
	ScoreYakuman
	ScoreDoubleYakuman
)

// if ron, Payment1 = score, Payment2 = 0
// if non-dealer tsumo, Payment1 = non-dealer payment, Payment2 = dealer payment
// if dealer tsumo, Payment1 = Payment2 = non-dealer payment
type HandScore struct {
	Payment1 int
	Payment2 int
	Type     ScoreType
}

func (hs HandScore) String() string {
	if hs.Payment2 == 0 { // ron
		return fmt.Sprintf("%d", hs.Payment1)
	} else if hs.Payment1 == hs.Payment2 {
		return fmt.Sprintf("%d all", hs.Payment2)
	} else {
		return fmt.Sprintf("%d/%d", hs.Payment1, hs.Payment2)
	}
}

func (t ScoreType) String() string {
	switch t {
	case ScoreMangan:
		return "Mangan"
	case ScoreHaneman:
		return "Haneman"
	case ScoreBaiman:
		return "Baiman"
	case ScoreSanbaiman:
		return "Sanbaiman"
	case ScoreYakuman:
		return "Yakuman"
	case ScoreDoubleYakuman:
		return "Double Yakuman"
	default:
		return ""
	}
}

func ScorePoints(han int, fu int, isTsumo bool, isDealer bool, honba int) (HandScore, error) {
	err := validateHanFu(han, fu, isTsumo)
	if err != nil {
		return HandScore{}, err
	}

	handScore := HandScore{}
	basicPoints := 0
	switch {
	case han >= 26:
		handScore.Type = ScoreDoubleYakuman
		basicPoints = 16000
	case han >= 13:
		handScore.Type = ScoreYakuman
		basicPoints = 8000
	case han >= 11:
		handScore.Type = ScoreSanbaiman
		basicPoints = 6000
	case han >= 8:
		handScore.Type = ScoreBaiman
		basicPoints = 4000
	case han >= 6:
		handScore.Type = ScoreHaneman
		basicPoints = 3000
	case han >= 5:
		handScore.Type = ScoreMangan
		basicPoints = 2000
	default:
		basicPoints = fu * (1 << (2 + han)) // fu * 2^(2+han)
		if basicPoints >= 2000 {
			handScore.Type = ScoreMangan
			basicPoints = 2000
		} else {
			handScore.Type = ScoreRegular
		}
	}

	if isDealer {
		if isTsumo {
			handScore.Payment1 = 2 * basicPoints
			handScore.Payment2 = 2 * basicPoints
		} else {
			handScore.Payment1 = 6 * basicPoints
		}
	} else {
		if isTsumo {
			handScore.Payment1 = basicPoints
			handScore.Payment2 = 2 * basicPoints
		} else {
			handScore.Payment1 = 4 * basicPoints
		}
	}

	handScore.Payment1 = roundUpToNearest100(handScore.Payment1)
	handScore.Payment2 = roundUpToNearest100(handScore.Payment2)

	// honba
	if isTsumo {
		handScore.Payment1 += 100 * honba
		handScore.Payment2 += 100 * honba
	} else {
		handScore.Payment1 += 300 * honba
	}

	return handScore, nil
}

func roundUpToNearest100(n int) int {
	return ((n + 99) / 100) * 100
}

func validateHanFu(han, fu int, isTsumo bool) error {
	if han < 1 || (14 <= han && han <= 25) || han > 26 {
		return fmt.Errorf("han must be between 1-13 or 26, got %d", han)
	}
	if fu < 20 || fu > 110 || (fu%10 != 0 && fu != 25) {
		return fmt.Errorf("fu must be between 20-110 and divisible by 10 (except for 25), got %d", fu)
	}
	if fu == 20 && (!isTsumo || han == 1) {
		return fmt.Errorf("invalid han/fu combination: %d han, %d fu", han, fu)
	}
	if fu == 25 && (han == 1 || (han == 2 && isTsumo)) {
		return fmt.Errorf("invalid han/fu combination: %d han, %d fu, tsumo=%t", han, fu, isTsumo)
	}
	if fu == 110 && han == 1 {
		return fmt.Errorf("invalid han/fu combination: %d han, %d fu", han, fu)
	}
	return nil
}
