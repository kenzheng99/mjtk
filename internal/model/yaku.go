package model

type YakuType string

const (
	// 1 han closed only
	Riichi   YakuType = "riichi"
	Tsumo    YakuType = "tsumo"
	Ippatsu  YakuType = "ippatsu"
	Pinfu    YakuType = "pinfu"
	Iipeikou YakuType = "iipeikou"

	// 1 han
	Haitei  YakuType = "haitei"
	Houtei  YakuType = "houtei"
	Rinshan YakuType = "rinshan"
	Chankan YakuType = "chankan"

	PrevalentWind YakuType = "prevalent wind"
	SeatWind      YakuType = "seat wind"
	WhiteDragon   YakuType = "white dragon"
	GreenDragon   YakuType = "green dragon"
	RedDragon     YakuType = "red dragon"
	Tanyao        YakuType = "tanyao"

	// 2 han, 1 open
	Chanta   YakuType = "chanta"
	Sanshoku YakuType = "sanshoku"
	Ittsu    YakuType = "ittsu"

	// 2 han
	DoubleRiichi   YakuType = "double riichi" // masks riichi
	Toitoi         YakuType = "toitoi"
	Sanankou       YakuType = "sanankou"
	SanshokuDoukou YakuType = "sanshoku doukou"
	Sankantsu      YakuType = "sankantsu"
	Chiitoi        YakuType = "chiitoi"   // masks everything else
	Honroutou      YakuType = "honroutou" // masks chanta
	Shousangen     YakuType = "shousangen"

	// 3 han, 2 open
	Honitsu YakuType = "honitsu"
	Junchan YakuType = "junchan" // masks chanta

	// 3 han
	Ryanpeikou YakuType = "ryanpeikou" // masks iipeikou

	// 6 han
	Chinitsu YakuType = "chinitsu" // masks honitsu

	// yakuman
	Kokushi       YakuType = "kokushi" // double if 13-sided wait
	Kokushi13     YakuType = "kokushi 13-sided"
	Suuankou      YakuType = "suuankou" // double if tanki wait
	SuuankouTanki YakuType = "suuankou tanki"
	Daisangen     YakuType = "daisangen"
	Shousuushii   YakuType = "shousuushii"
	Daisuushii    YakuType = "daisuushii" // double
	Tsuuiisou     YakuType = "tsuuiisou"
	Chinroutou    YakuType = "chinroutou"
	Ryuuiisou     YakuType = "ryuuisou"
	Chuuren       YakuType = "chuuren"
	Suukantsu     YakuType = "suukantsu"
	Tenhou        YakuType = "tenhou"
	Chiihou       YakuType = "chiihou"

	// special
	Nagashi YakuType = "nagashi" // mangan at draw
)

func CalculateYaku(ph ParsedHand, gs GameState) []YakuType {
	var yaku []YakuType
	yaku = CheckRiichiIppatsuTsumo(yaku, ph)
	yaku = CheckPinfu(yaku, ph, gs)
	yaku = CheckYakuhaiShousangen(yaku, ph, gs)
	yaku = CheckTanyao(yaku, ph)
	yaku = CheckToitoiSanankou(yaku, ph)
	yaku = CheckIipeikouRyanpeikou(yaku, ph)
	yaku = CheckIttsu(yaku, ph)
	yaku = CheckSanshoku(yaku, ph)
	yaku = CheckChantaJunchanHonroutou(yaku, ph)
	yaku = CheckHonitsuChinitsu(yaku, ph)
	return yaku
}

func CheckRiichiIppatsuTsumo(yaku []YakuType, ph ParsedHand) []YakuType {
	if ph.State.IsRiichi {
		yaku = append(yaku, Riichi)
	}
	if ph.State.IsIppatsu {
		yaku = append(yaku, Ippatsu)
	}
	if ph.State.IsTsumo {
		yaku = append(yaku, Tsumo)
	}
	return yaku
}

func CheckPinfu(yaku []YakuType, ph ParsedHand, gs GameState) []YakuType {
	for _, group := range ph.Groups {
		if !group.IsSequence() {
			return yaku
		}
	}
	pairTile := ph.Pair.Tiles[0]
	if pairTile.IsWind(gs.PrevalentWind) ||
		pairTile.IsWind(ph.State.SeatWind) ||
		pairTile.IsDragon() {
		return yaku
	}
	if ph.WaitTypes().Has(WaitRyanmen) {
		yaku = append(yaku, Pinfu)
	}
	return yaku
}

func CheckYakuhaiShousangen(yaku []YakuType, ph ParsedHand, gs GameState) []YakuType {
	dragonTriplets := 0
	for _, group := range ph.Groups {
		if !group.IsTriplet() {
			continue
		}
		tile := group.Tiles[0]
		if tile.IsWind(gs.PrevalentWind) {
			yaku = append(yaku, PrevalentWind)
		}
		if tile.IsWind(ph.State.SeatWind) {
			yaku = append(yaku, SeatWind)
		}
		if tile.IsWhiteDragon() {
			yaku = append(yaku, WhiteDragon)
			dragonTriplets++
		}
		if tile.IsGreenDragon() {
			yaku = append(yaku, GreenDragon)
			dragonTriplets++
		}
		if tile.IsRedDragon() {
			yaku = append(yaku, RedDragon)
			dragonTriplets++
		}
	}
	if dragonTriplets == 2 && ph.Pair.Tiles[0].IsDragon() {
		yaku = append(yaku, Shousangen)
	}
	return yaku
}

func CheckTanyao(yaku []YakuType, ph ParsedHand) []YakuType {
	for _, tile := range ph.Tiles {
		if tile.IsHonor() || tile.IsTerminal() {
			return yaku
		}
	}
	yaku = append(yaku, Tanyao)
	return yaku
}

func CheckIipeikouRyanpeikou(yaku []YakuType, ph ParsedHand) []YakuType {
	firstIipeikou := ph.Groups[0].Equals(ph.Groups[1])
	midIipeikou := ph.Groups[1].Equals(ph.Groups[2])
	lastIipeikou := ph.Groups[2].Equals(ph.Groups[3])

	if firstIipeikou && lastIipeikou {
		yaku = append(yaku, Ryanpeikou)
	} else if firstIipeikou || midIipeikou || lastIipeikou {
		yaku = append(yaku, Iipeikou)
	}
	return yaku
}

func CheckIttsu(yaku []YakuType, ph ParsedHand) []YakuType {
	ittsu012 := ph.Groups[0].IsSequence() &&
		ph.Groups[0].Tiles[0].FaceValue() == 1 &&
		ph.Groups[1].IsSequence() &&
		ph.Groups[1].Tiles[0].FaceValue() == 4 &&
		ph.Groups[2].IsSequence() &&
		ph.Groups[2].Tiles[0].FaceValue() == 7 &&
		ph.Groups[0].SuitEquals(ph.Groups[1]) &&
		ph.Groups[1].SuitEquals(ph.Groups[2])

	ittsu123 := ph.Groups[1].IsSequence() &&
		ph.Groups[1].Tiles[0].FaceValue() == 1 &&
		ph.Groups[2].IsSequence() &&
		ph.Groups[2].Tiles[0].FaceValue() == 4 &&
		ph.Groups[3].IsSequence() &&
		ph.Groups[3].Tiles[0].FaceValue() == 7 &&
		ph.Groups[1].SuitEquals(ph.Groups[2]) &&
		ph.Groups[2].SuitEquals(ph.Groups[3])

	if ittsu012 || ittsu123 {
		yaku = append(yaku, Ittsu)
	}
	return yaku
}

func CheckSanshoku(yaku []YakuType, ph ParsedHand) []YakuType {
	equal01 := ph.Groups[0].ValueEquals(ph.Groups[1])
	equal02 := ph.Groups[0].ValueEquals(ph.Groups[2])
	equal03 := ph.Groups[0].ValueEquals(ph.Groups[3])
	equal12 := ph.Groups[1].ValueEquals(ph.Groups[2])
	equal13 := ph.Groups[1].ValueEquals(ph.Groups[3])

	suitEqual01 := ph.Groups[0].SuitEquals(ph.Groups[1])
	suitEqual02 := ph.Groups[0].SuitEquals(ph.Groups[2])
	suitEqual03 := ph.Groups[0].SuitEquals(ph.Groups[3])
	suitEqual12 := ph.Groups[1].SuitEquals(ph.Groups[2])
	suitEqual13 := ph.Groups[1].SuitEquals(ph.Groups[3])

	isSanshoku := false
	isDoukou := false
	if (equal01 && equal02 && !suitEqual01 && !suitEqual02) ||
		(equal01 && equal03 && !suitEqual01 && !suitEqual03) ||
		(equal02 && equal03 && !suitEqual02 && !suitEqual03) {
		isSanshoku = true
		isDoukou = ph.Groups[0].IsTriplet()
	} else if equal12 && equal13 && !suitEqual12 && !suitEqual13 {
		isSanshoku = true
		isDoukou = ph.Groups[1].IsTriplet()
	}

	if isDoukou {
		yaku = append(yaku, SanshokuDoukou)
	} else if isSanshoku {
		yaku = append(yaku, Sanshoku)
	}
	return yaku
}

func CheckToitoiSanankou(yaku []YakuType, ph ParsedHand) []YakuType {
	for _, group := range ph.Groups {
		if group.IsSequence() {
			return yaku
		}
	}
	yaku = append(yaku, Toitoi)
	yaku = append(yaku, Sanankou)
	return yaku
}

func CheckChantaJunchanHonroutou(yaku []YakuType, ph ParsedHand) []YakuType {
	hasHonors := false
	hasSequence := false
	for _, group := range ph.Groups {
		if group.Tiles[0].IsHonor() {
			hasHonors = true
		} else if !group.HasTerminal() {
			return yaku
		} else if group.IsSequence() {
			hasSequence = true
		}
	}

	if !hasSequence {
		yaku = append(yaku, Honroutou)
	} else if !hasHonors {
		yaku = append(yaku, Junchan)
	} else {
		yaku = append(yaku, Chanta)
	}
	return yaku
}

func CheckHonitsuChinitsu(yaku []YakuType, ph ParsedHand) []YakuType {
	hasHonors := false
	sameSuit := true
	var suit byte

	for _, tile := range ph.Tiles {
		if tile.IsHonor() {
			hasHonors = true
		} else if suit == 0 {
			suit = tile.Suit
		} else {
			sameSuit = tile.Suit == suit
		}
	}

	if !hasHonors && sameSuit {
		yaku = append(yaku, Chinitsu)
	} else if sameSuit {
		yaku = append(yaku, Honitsu)
	}
	return yaku
}
