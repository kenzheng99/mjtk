package model

import "strings"

type WaitType uint8

const (
	WaitRyanmen WaitType = 1 << iota
	WaitKanchan
	WaitPenchan
	WaitShanpon
	WaitTanki
)

func (w WaitType) Has(t WaitType) bool {
	return w&t != 0
}

func (w WaitType) String() string {
	var waits []string
	if w.Has(WaitRyanmen) {
		waits = append(waits, "ryanmen")
	}
	if w.Has(WaitKanchan) {
		waits = append(waits, "kanchan")
	}
	if w.Has(WaitPenchan) {
		waits = append(waits, "penchan")
	}
	if w.Has(WaitShanpon) {
		waits = append(waits, "shanpon")
	}
	if w.Has(WaitTanki) {
		waits = append(waits, "tanki")
	}
	return "WaitType: " + strings.Join(waits, ", ")
}

// only considers triples, doesn't include tanki
func (tg TileGroup) WaitTypes(waitIndex int) WaitType {
	if tg.IsTriplet() {
		return WaitShanpon
	}

	if !tg.IsSequence() {
		return 0
	}

	if waitIndex == 1 {
		return WaitKanchan
	}

	if waitIndex == 0 && tg.Tiles[2].IsTerminal() ||
		waitIndex == 2 && tg.Tiles[0].IsTerminal() {
		return WaitPenchan
	}

	return WaitRyanmen
}
