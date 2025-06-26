package scorer

import "testing"

type scorePointsTest struct {
	name string
	// inputs
	han      int
	fu       int
	isTsumo  bool
	isDealer bool
	honba    int
	// expected
	expectedHandScore HandScore
}

func TestScorePointsTable(t *testing.T) {
	var tests = []scorePointsTest{
		// nondealer ron
		{"1han_30fu_ron", 1, 30, false, false, 0, HandScore{1000, 0, ScoreRegular}},
		{"1han_40fu_ron", 1, 40, false, false, 0, HandScore{1300, 0, ScoreRegular}},
		{"1han_50fu_ron", 1, 50, false, false, 0, HandScore{1600, 0, ScoreRegular}},
		{"1han_60fu_ron", 1, 60, false, false, 0, HandScore{2000, 0, ScoreRegular}},
		{"1han_70fu_ron", 1, 70, false, false, 0, HandScore{2300, 0, ScoreRegular}},
		{"2han_25fu_ron", 2, 25, false, false, 0, HandScore{1600, 0, ScoreRegular}},
		{"2han_30fu_ron", 2, 30, false, false, 0, HandScore{2000, 0, ScoreRegular}},
		{"2han_40fu_ron", 2, 40, false, false, 0, HandScore{2600, 0, ScoreRegular}},
		{"2han_50fu_ron", 2, 50, false, false, 0, HandScore{3200, 0, ScoreRegular}},
		{"2han_60fu_ron", 2, 60, false, false, 0, HandScore{3900, 0, ScoreRegular}},
		{"2han_70fu_ron", 2, 70, false, false, 0, HandScore{4500, 0, ScoreRegular}},
		{"3han_25fu_ron", 3, 25, false, false, 0, HandScore{3200, 0, ScoreRegular}},
		{"3han_30fu_ron", 3, 30, false, false, 0, HandScore{3900, 0, ScoreRegular}},
		{"3han_40fu_ron", 3, 40, false, false, 0, HandScore{5200, 0, ScoreRegular}},
		{"3han_50fu_ron", 3, 50, false, false, 0, HandScore{6400, 0, ScoreRegular}},
		{"3han_60fu_ron", 3, 60, false, false, 0, HandScore{7700, 0, ScoreRegular}},
		{"3han_70fu_ron", 3, 70, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"4han_25fu_ron", 4, 25, false, false, 0, HandScore{6400, 0, ScoreRegular}},
		{"4han_30fu_ron", 4, 30, false, false, 0, HandScore{7700, 0, ScoreRegular}},
		{"4han_40fu_ron", 4, 40, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"4han_50fu_ron", 4, 50, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"4han_60fu_ron", 4, 60, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"4han_70fu_ron", 4, 70, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"5han_30fu_ron", 5, 30, false, false, 0, HandScore{8000, 0, ScoreMangan}},
		{"6han_30fu_ron", 6, 30, false, false, 0, HandScore{12000, 0, ScoreHaneman}},
		{"7han_30fu_ron", 7, 30, false, false, 0, HandScore{12000, 0, ScoreHaneman}},
		{"8han_30fu_ron", 8, 30, false, false, 0, HandScore{16000, 0, ScoreBaiman}},
		{"9han_30fu_ron", 9, 30, false, false, 0, HandScore{16000, 0, ScoreBaiman}},
		{"10han_30fu_ron", 10, 30, false, false, 0, HandScore{16000, 0, ScoreBaiman}},
		{"11han_30fu_ron", 11, 30, false, false, 0, HandScore{24000, 0, ScoreSanbaiman}},
		{"12han_30fu_ron", 12, 30, false, false, 0, HandScore{24000, 0, ScoreSanbaiman}},
		{"13han_30fu_ron", 13, 30, false, false, 0, HandScore{32000, 0, ScoreYakuman}},
		{"26han_30fu_ron", 26, 30, false, false, 0, HandScore{64000, 0, ScoreDoubleYakuman}},

		// nondealer tsumo
		{"1han_30fu_tsumo", 1, 30, true, false, 0, HandScore{300, 500, ScoreRegular}},
		{"1han_40fu_tsumo", 1, 40, true, false, 0, HandScore{400, 700, ScoreRegular}},
		{"1han_50fu_tsumo", 1, 50, true, false, 0, HandScore{400, 800, ScoreRegular}},
		{"1han_60fu_tsumo", 1, 60, true, false, 0, HandScore{500, 1000, ScoreRegular}},
		{"1han_70fu_tsumo", 1, 70, true, false, 0, HandScore{600, 1200, ScoreRegular}},
		{"2han_20fu_tsumo", 2, 20, true, false, 0, HandScore{400, 700, ScoreRegular}},
		{"2han_30fu_tsumo", 2, 30, true, false, 0, HandScore{500, 1000, ScoreRegular}},
		{"2han_40fu_tsumo", 2, 40, true, false, 0, HandScore{700, 1300, ScoreRegular}},
		{"2han_50fu_tsumo", 2, 50, true, false, 0, HandScore{800, 1600, ScoreRegular}},
		{"2han_60fu_tsumo", 2, 60, true, false, 0, HandScore{1000, 2000, ScoreRegular}},
		{"2han_70fu_tsumo", 2, 70, true, false, 0, HandScore{1200, 2300, ScoreRegular}},
		{"3han_20fu_tsumo", 3, 20, true, false, 0, HandScore{700, 1300, ScoreRegular}},
		{"3han_25fu_tsumo", 3, 25, true, false, 0, HandScore{800, 1600, ScoreRegular}},
		{"3han_30fu_tsumo", 3, 30, true, false, 0, HandScore{1000, 2000, ScoreRegular}},
		{"3han_40fu_tsumo", 3, 40, true, false, 0, HandScore{1300, 2600, ScoreRegular}},
		{"3han_50fu_tsumo", 3, 50, true, false, 0, HandScore{1600, 3200, ScoreRegular}},
		{"3han_60fu_tsumo", 3, 60, true, false, 0, HandScore{2000, 3900, ScoreRegular}},
		{"3han_70fu_tsumo", 3, 70, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"4han_20fu_tsumo", 4, 20, true, false, 0, HandScore{1300, 2600, ScoreRegular}},
		{"4han_25fu_tsumo", 4, 25, true, false, 0, HandScore{1600, 3200, ScoreRegular}},
		{"4han_30fu_tsumo", 4, 30, true, false, 0, HandScore{2000, 3900, ScoreRegular}},
		{"4han_40fu_tsumo", 4, 40, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"4han_50fu_tsumo", 4, 50, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"4han_60fu_tsumo", 4, 60, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"4han_70fu_tsumo", 4, 70, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"5han_30fu_tsumo", 5, 30, true, false, 0, HandScore{2000, 4000, ScoreMangan}},
		{"6han_30fu_tsumo", 6, 30, true, false, 0, HandScore{3000, 6000, ScoreHaneman}},
		{"7han_30fu_tsumo", 7, 30, true, false, 0, HandScore{3000, 6000, ScoreHaneman}},
		{"8han_30fu_tsumo", 8, 30, true, false, 0, HandScore{4000, 8000, ScoreBaiman}},
		{"9han_30fu_tsumo", 9, 30, true, false, 0, HandScore{4000, 8000, ScoreBaiman}},
		{"10han_30fu_tsumo", 10, 30, true, false, 0, HandScore{4000, 8000, ScoreBaiman}},
		{"11han_30fu_tsumo", 11, 30, true, false, 0, HandScore{6000, 12000, ScoreSanbaiman}},
		{"12han_30fu_tsumo", 12, 30, true, false, 0, HandScore{6000, 12000, ScoreSanbaiman}},
		{"13han_30fu_tsumo", 13, 30, true, false, 0, HandScore{8000, 16000, ScoreYakuman}},
		{"26han_30fu_tsumo", 26, 30, true, false, 0, HandScore{16000, 32000, ScoreDoubleYakuman}},

		// dealer ron
		{"1han_30fu_dealer_ron", 1, 30, false, true, 0, HandScore{1500, 0, ScoreRegular}},
		{"1han_40fu_dealer_ron", 1, 40, false, true, 0, HandScore{2000, 0, ScoreRegular}},
		{"1han_50fu_dealer_ron", 1, 50, false, true, 0, HandScore{2400, 0, ScoreRegular}},
		{"1han_60fu_dealer_ron", 1, 60, false, true, 0, HandScore{2900, 0, ScoreRegular}},
		{"1han_70fu_dealer_ron", 1, 70, false, true, 0, HandScore{3400, 0, ScoreRegular}},
		{"2han_25fu_dealer_ron", 2, 25, false, true, 0, HandScore{2400, 0, ScoreRegular}},
		{"2han_30fu_dealer_ron", 2, 30, false, true, 0, HandScore{2900, 0, ScoreRegular}},
		{"2han_40fu_dealer_ron", 2, 40, false, true, 0, HandScore{3900, 0, ScoreRegular}},
		{"2han_50fu_dealer_ron", 2, 50, false, true, 0, HandScore{4800, 0, ScoreRegular}},
		{"2han_60fu_dealer_ron", 2, 60, false, true, 0, HandScore{5800, 0, ScoreRegular}},
		{"2han_70fu_dealer_ron", 2, 70, false, true, 0, HandScore{6800, 0, ScoreRegular}},
		{"3han_25fu_dealer_ron", 3, 25, false, true, 0, HandScore{4800, 0, ScoreRegular}},
		{"3han_30fu_dealer_ron", 3, 30, false, true, 0, HandScore{5800, 0, ScoreRegular}},
		{"3han_40fu_dealer_ron", 3, 40, false, true, 0, HandScore{7700, 0, ScoreRegular}},
		{"3han_50fu_dealer_ron", 3, 50, false, true, 0, HandScore{9600, 0, ScoreRegular}},
		{"3han_60fu_dealer_ron", 3, 60, false, true, 0, HandScore{11600, 0, ScoreRegular}},
		{"3han_70fu_dealer_ron", 3, 70, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"4han_25fu_dealer_ron", 4, 25, false, true, 0, HandScore{9600, 0, ScoreRegular}},
		{"4han_30fu_dealer_ron", 4, 30, false, true, 0, HandScore{11600, 0, ScoreRegular}},
		{"4han_40fu_dealer_ron", 4, 40, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"4han_50fu_dealer_ron", 4, 50, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"4han_60fu_dealer_ron", 4, 60, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"4han_70fu_dealer_ron", 4, 70, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"5han_30fu_dealer_ron", 5, 30, false, true, 0, HandScore{12000, 0, ScoreMangan}},
		{"6han_30fu_dealer_ron", 6, 30, false, true, 0, HandScore{18000, 0, ScoreHaneman}},
		{"7han_30fu_dealer_ron", 7, 30, false, true, 0, HandScore{18000, 0, ScoreHaneman}},
		{"8han_30fu_dealer_ron", 8, 30, false, true, 0, HandScore{24000, 0, ScoreBaiman}},
		{"9han_30fu_dealer_ron", 9, 30, false, true, 0, HandScore{24000, 0, ScoreBaiman}},
		{"10han_30fu_dealer_ron", 10, 30, false, true, 0, HandScore{24000, 0, ScoreBaiman}},
		{"11han_30fu_dealer_ron", 11, 30, false, true, 0, HandScore{36000, 0, ScoreSanbaiman}},
		{"12han_30fu_dealer_ron", 12, 30, false, true, 0, HandScore{36000, 0, ScoreSanbaiman}},
		{"13han_30fu_dealer_ron", 13, 30, false, true, 0, HandScore{48000, 0, ScoreYakuman}},
		{"26han_30fu_dealer_ron", 26, 30, false, true, 0, HandScore{96000, 0, ScoreDoubleYakuman}},

		// dealer tsumo
		{"1han_30fu_dealer_tsumo", 1, 30, true, true, 0, HandScore{500, 500, ScoreRegular}},
		{"1han_40fu_dealer_tsumo", 1, 40, true, true, 0, HandScore{700, 700, ScoreRegular}},
		{"1han_50fu_dealer_tsumo", 1, 50, true, true, 0, HandScore{800, 800, ScoreRegular}},
		{"1han_60fu_dealer_tsumo", 1, 60, true, true, 0, HandScore{1000, 1000, ScoreRegular}},
		{"1han_70fu_dealer_tsumo", 1, 70, true, true, 0, HandScore{1200, 1200, ScoreRegular}},
		{"2han_20fu_dealer_tsumo", 2, 20, true, true, 0, HandScore{700, 700, ScoreRegular}},
		{"2han_30fu_dealer_tsumo", 2, 30, true, true, 0, HandScore{1000, 1000, ScoreRegular}},
		{"2han_40fu_dealer_tsumo", 2, 40, true, true, 0, HandScore{1300, 1300, ScoreRegular}},
		{"2han_50fu_dealer_tsumo", 2, 50, true, true, 0, HandScore{1600, 1600, ScoreRegular}},
		{"2han_60fu_dealer_tsumo", 2, 60, true, true, 0, HandScore{2000, 2000, ScoreRegular}},
		{"2han_70fu_dealer_tsumo", 2, 70, true, true, 0, HandScore{2300, 2300, ScoreRegular}},
		{"3han_20fu_dealer_tsumo", 3, 20, true, true, 0, HandScore{1300, 1300, ScoreRegular}},
		{"3han_25fu_dealer_tsumo", 3, 25, true, true, 0, HandScore{1600, 1600, ScoreRegular}},
		{"3han_30fu_dealer_tsumo", 3, 30, true, true, 0, HandScore{2000, 2000, ScoreRegular}},
		{"3han_40fu_dealer_tsumo", 3, 40, true, true, 0, HandScore{2600, 2600, ScoreRegular}},
		{"3han_50fu_dealer_tsumo", 3, 50, true, true, 0, HandScore{3200, 3200, ScoreRegular}},
		{"3han_60fu_dealer_tsumo", 3, 60, true, true, 0, HandScore{3900, 3900, ScoreRegular}},
		{"3han_70fu_dealer_tsumo", 3, 70, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"4han_20fu_dealer_tsumo", 4, 20, true, true, 0, HandScore{2600, 2600, ScoreRegular}},
		{"4han_25fu_dealer_tsumo", 4, 25, true, true, 0, HandScore{3200, 3200, ScoreRegular}},
		{"4han_30fu_dealer_tsumo", 4, 30, true, true, 0, HandScore{3900, 3900, ScoreRegular}},
		{"4han_40fu_dealer_tsumo", 4, 40, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"4han_50fu_dealer_tsumo", 4, 50, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"4han_60fu_dealer_tsumo", 4, 60, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"4han_70fu_dealer_tsumo", 4, 70, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"5han_30fu_dealer_tsumo", 5, 30, true, true, 0, HandScore{4000, 4000, ScoreMangan}},
		{"6han_30fu_dealer_tsumo", 6, 30, true, true, 0, HandScore{6000, 6000, ScoreHaneman}},
		{"7han_30fu_dealer_tsumo", 7, 30, true, true, 0, HandScore{6000, 6000, ScoreHaneman}},
		{"8han_30fu_dealer_tsumo", 8, 30, true, true, 0, HandScore{8000, 8000, ScoreBaiman}},
		{"9han_30fu_dealer_tsumo", 9, 30, true, true, 0, HandScore{8000, 8000, ScoreBaiman}},
		{"10han_30fu_dealer_tsumo", 10, 30, true, true, 0, HandScore{8000, 8000, ScoreBaiman}},
		{"11han_30fu_dealer_tsumo", 11, 30, true, true, 0, HandScore{12000, 12000, ScoreSanbaiman}},
		{"12han_30fu_dealer_tsumo", 12, 30, true, true, 0, HandScore{12000, 12000, ScoreSanbaiman}},
		{"13han_30fu_dealer_tsumo", 13, 30, true, true, 0, HandScore{16000, 16000, ScoreYakuman}},
		{"26han_30fu_dealer_tsumo", 26, 30, true, true, 0, HandScore{32000, 32000, ScoreDoubleYakuman}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handScore, err := ScorePoints(tt.han, tt.fu, tt.isTsumo, tt.isDealer, tt.honba)
			if err != nil {
				t.Error(err)
			}
			if handScore != tt.expectedHandScore {
				t.Errorf("wrong hand score, actual: %s, expected: %s",
					handScore,
					tt.expectedHandScore)
			}
			if handScore.Type != tt.expectedHandScore.Type {
				t.Errorf("wrong hand type, actual: %s, expected: %s",
					handScore.Type,
					tt.expectedHandScore.Type)
			}
		})
	}

}
