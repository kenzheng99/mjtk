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

	Tanyao    YakuType = "tanyao"
	SeatWind  YakuType = "seat wind"
	RoundWind YakuType = "round wind"
	Haku      YakuType = "white dragon"
	Hatsu     YakuType = "green dragon"
	Chun      YakuType = "red dragon"

	// 2 han, 1 open
	Chanta   YakuType = "chanta"
	Sanshoku YakuType = "sanshoku"
	Iitsu    YakuType = "iitsu"

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
	Kokushi     YakuType = "kokushi"  // double if 13-sided wait
	Suuankou    YakuType = "suuankou" // double if tanki wait
	Daisangen   YakuType = "daisangen"
	Shousuushii YakuType = "shousuushii"
	Daisuushii  YakuType = "daisuushii" // double
	Tsuuiisou   YakuType = "tsuuiisou"
	Chinroutou  YakuType = "chinroutou"
	Ryuuiisou   YakuType = "ryuuisou"
	Chuuren     YakuType = "chuuren"
	Suukantsu   YakuType = "suukantsu"
	Tenhou      YakuType = "tenhou"
	Chiihou     YakuType = "chiihou"

	// special
	Nagashi YakuType = "nagashi" // mangan at draw
)

type Yaku struct {
	name      YakuType
	han       int
	isYakuman bool
}

func (ph ParsedHand) HandYakuType() []Yaku {
	var yaku []Yaku
	yaku = CheckRiichiIppatsuTsumo(yaku, ph)
	yaku = CheckPinfu(yaku, ph)
	yaku = CheckTanyao(yaku, ph)
	yaku = CheckChinitsuHonitsu(yaku, ph)
	return yaku
}

func CheckRiichiIppatsuTsumo(yaku []Yaku, ph ParsedHand) []Yaku {
	if ph.State.IsRiichi {
		yaku = append(yaku, Yaku{
			name: Riichi,
			han:  1,
		})
	}
	if ph.State.IsIppatsu {
		yaku = append(yaku, Yaku{
			name: Ippatsu,
			han:  1,
		})
	}
	if ph.State.IsTsumo {
		yaku = append(yaku, Yaku{
			name: Tsumo,
			han:  1,
		})
	}
	return yaku
}

func CheckPinfu(yaku []Yaku, ph ParsedHand) []Yaku {
	for _, group := range ph.Groups {
		if !group.IsSequence() {
			return yaku
		}
	}
	// TODO check non-yakuhai pair
	if ph.WaitTypes().Has(WaitRyanmen) {
		yaku = append(yaku, Yaku{
			name: Pinfu,
			han:  1,
		})
	}
	return yaku
}

func CheckTanyao(yaku []Yaku, ph ParsedHand) []Yaku {
	for _, tile := range ph.Tiles {
		if tile.IsHonor() || tile.IsTerminal() {
			return yaku
		}
	}
	yaku = append(yaku, Yaku{
		name: Tanyao,
		han:  1,
	})
	return yaku
}

func CheckChinitsuHonitsu(yaku []Yaku, ph ParsedHand) []Yaku {
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
		yaku = append(yaku, Yaku{
			name: Chinitsu,
			han:  6,
		})
	} else if sameSuit {
		yaku = append(yaku, Yaku{
			name: Honitsu,
			han:  3,
		})
	}
	return yaku
}

func CheckYakuhai(yaku []Yaku, ph ParsedHand) []Yaku {
	return yaku
}

func (y Yaku) String() string {
	return string(y.name)
}
