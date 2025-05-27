package model

type HandState struct {
	IsDealer  bool
	IsRiichi  bool
	IsIppatsu bool
	IsTsumo   bool
	IsRon     bool
	IsClosed  bool
	IsTenpai  bool

	IsDoubleRiichi bool
	IsHaitei       bool
	IsHoutei       bool
	IsRinshan      bool
	IsChankan      bool
}

func (hs HandState) String() string {
	var str string
	if hs.IsDealer {
		str += "dealer, "
	}
	if hs.IsRiichi {
		str += "riichi, "
	}
	if hs.IsIppatsu {
		str += "ippatsu,"
	}
	if hs.IsTsumo {
		str += "tsumo, "
	}
	if len(str) >= 0 {
		str = str[:len(str)-2]
	}

	return str
}
