package model

import "github.com/uptrace/bun"

type PvpstatsPlayers struct {
	BattlegroundId int `json:"battleground_id,omitempty"`
	CharacterGuid int `json:"character_guid,omitempty"`
	Winner string `json:"winner,omitempty"`
	ScoreKillingBlows int32 `json:"score_killing_blows,omitempty"`
	ScoreDeaths int32 `json:"score_deaths,omitempty"`
	ScoreHonorableKills int32 `json:"score_honorable_kills,omitempty"`
	ScoreBonusHonor int32 `json:"score_bonus_honor,omitempty"`
	ScoreDamageDone int32 `json:"score_damage_done,omitempty"`
	ScoreHealingDone int32 `json:"score_healing_done,omitempty"`
	Attr1 int32 `json:"attr_1,omitempty"`
	Attr2 int32 `json:"attr_2,omitempty"`
	Attr3 int32 `json:"attr_3,omitempty"`
	Attr4 int32 `json:"attr_4,omitempty"`
	Attr5 int32 `json:"attr_5,omitempty"`
}

type PvpstatsPlayersSlice []PvpstatsPlayers

type DBPvpstatsPlayers struct {
	bun.BaseModel `bun:"table:pvpstats_players,alias:pvpstats_players"`
	BattlegroundId int `bun:"battleground_id"`
	CharacterGuid int `bun:"character_guid"`
	Winner string `bun:"winner"`
	ScoreKillingBlows int32 `bun:"score_killing_blows"`
	ScoreDeaths int32 `bun:"score_deaths"`
	ScoreHonorableKills int32 `bun:"score_honorable_kills"`
	ScoreBonusHonor int32 `bun:"score_bonus_honor"`
	ScoreDamageDone int32 `bun:"score_damage_done"`
	ScoreHealingDone int32 `bun:"score_healing_done"`
	Attr1 int32 `bun:"attr_1"`
	Attr2 int32 `bun:"attr_2"`
	Attr3 int32 `bun:"attr_3"`
	Attr4 int32 `bun:"attr_4"`
	Attr5 int32 `bun:"attr_5"`
}

type DBPvpstatsPlayersSlice []DBPvpstatsPlayers

func (entry *PvpstatsPlayers) ToDB() *DBPvpstatsPlayers {
	if entry == nil {
		return nil
	}
	return &DBPvpstatsPlayers{
		BattlegroundId: entry.BattlegroundId,
		CharacterGuid: entry.CharacterGuid,
		Winner: entry.Winner,
		ScoreKillingBlows: entry.ScoreKillingBlows,
		ScoreDeaths: entry.ScoreDeaths,
		ScoreHonorableKills: entry.ScoreHonorableKills,
		ScoreBonusHonor: entry.ScoreBonusHonor,
		ScoreDamageDone: entry.ScoreDamageDone,
		ScoreHealingDone: entry.ScoreHealingDone,
		Attr1: entry.Attr1,
		Attr2: entry.Attr2,
		Attr3: entry.Attr3,
		Attr4: entry.Attr4,
		Attr5: entry.Attr5,
	}
}

func (entry *DBPvpstatsPlayers) ToWeb() *PvpstatsPlayers {
	if entry == nil {
		return nil
	}
	return &PvpstatsPlayers{
		BattlegroundId: entry.BattlegroundId,
		CharacterGuid: entry.CharacterGuid,
		Winner: entry.Winner,
		ScoreKillingBlows: entry.ScoreKillingBlows,
		ScoreDeaths: entry.ScoreDeaths,
		ScoreHonorableKills: entry.ScoreHonorableKills,
		ScoreBonusHonor: entry.ScoreBonusHonor,
		ScoreDamageDone: entry.ScoreDamageDone,
		ScoreHealingDone: entry.ScoreHealingDone,
		Attr1: entry.Attr1,
		Attr2: entry.Attr2,
		Attr3: entry.Attr3,
		Attr4: entry.Attr4,
		Attr5: entry.Attr5,
	}
}

func (data PvpstatsPlayersSlice) ToDB() DBPvpstatsPlayersSlice {
	result := make(DBPvpstatsPlayersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPvpstatsPlayersSlice) ToWeb() PvpstatsPlayersSlice {
	result := make(PvpstatsPlayersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
