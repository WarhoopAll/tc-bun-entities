package model

import "github.com/uptrace/bun"

type ReputationRewardRate struct {
	Faction int `json:"faction,omitempty"`
	QuestRate float64 `json:"quest_rate,omitempty"`
	QuestDailyRate float64 `json:"quest_daily_rate,omitempty"`
	QuestWeeklyRate float64 `json:"quest_weekly_rate,omitempty"`
	QuestMonthlyRate float64 `json:"quest_monthly_rate,omitempty"`
	QuestRepeatableRate float64 `json:"quest_repeatable_rate,omitempty"`
	CreatureRate float64 `json:"creature_rate,omitempty"`
	SpellRate float64 `json:"spell_rate,omitempty"`
}

type ReputationRewardRateSlice []ReputationRewardRate

type DBReputationRewardRate struct {
	bun.BaseModel `bun:"table:reputation_reward_rate,alias:reputation_reward_rate"`
	Faction int `bun:"faction"`
	QuestRate float64 `bun:"quest_rate"`
	QuestDailyRate float64 `bun:"quest_daily_rate"`
	QuestWeeklyRate float64 `bun:"quest_weekly_rate"`
	QuestMonthlyRate float64 `bun:"quest_monthly_rate"`
	QuestRepeatableRate float64 `bun:"quest_repeatable_rate"`
	CreatureRate float64 `bun:"creature_rate"`
	SpellRate float64 `bun:"spell_rate"`
}

type DBReputationRewardRateSlice []DBReputationRewardRate

func (entry *ReputationRewardRate) ToDB() *DBReputationRewardRate {
	if entry == nil {
		return nil
	}
	return &DBReputationRewardRate{
		Faction: entry.Faction,
		QuestRate: entry.QuestRate,
		QuestDailyRate: entry.QuestDailyRate,
		QuestWeeklyRate: entry.QuestWeeklyRate,
		QuestMonthlyRate: entry.QuestMonthlyRate,
		QuestRepeatableRate: entry.QuestRepeatableRate,
		CreatureRate: entry.CreatureRate,
		SpellRate: entry.SpellRate,
	}
}

func (entry *DBReputationRewardRate) ToWeb() *ReputationRewardRate {
	if entry == nil {
		return nil
	}
	return &ReputationRewardRate{
		Faction: entry.Faction,
		QuestRate: entry.QuestRate,
		QuestDailyRate: entry.QuestDailyRate,
		QuestWeeklyRate: entry.QuestWeeklyRate,
		QuestMonthlyRate: entry.QuestMonthlyRate,
		QuestRepeatableRate: entry.QuestRepeatableRate,
		CreatureRate: entry.CreatureRate,
		SpellRate: entry.SpellRate,
	}
}

func (data ReputationRewardRateSlice) ToDB() DBReputationRewardRateSlice {
	result := make(DBReputationRewardRateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBReputationRewardRateSlice) ToWeb() ReputationRewardRateSlice {
	result := make(ReputationRewardRateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
