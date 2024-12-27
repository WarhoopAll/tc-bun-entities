package model

import "github.com/uptrace/bun"

type SpellGroupStackRules struct {
	GroupId int `json:"group_id,omitempty"`
	StackRule int8 `json:"stack_rule,omitempty"`
}

type SpellGroupStackRulesSlice []SpellGroupStackRules

type DBSpellGroupStackRules struct {
	bun.BaseModel `bun:"table:spell_group_stack_rules,alias:spell_group_stack_rules"`
	GroupId int `bun:"group_id"`
	StackRule int8 `bun:"stack_rule"`
}

type DBSpellGroupStackRulesSlice []DBSpellGroupStackRules

func (entry *SpellGroupStackRules) ToDB() *DBSpellGroupStackRules {
	if entry == nil {
		return nil
	}
	return &DBSpellGroupStackRules{
		GroupId: entry.GroupId,
		StackRule: entry.StackRule,
	}
}

func (entry *DBSpellGroupStackRules) ToWeb() *SpellGroupStackRules {
	if entry == nil {
		return nil
	}
	return &SpellGroupStackRules{
		GroupId: entry.GroupId,
		StackRule: entry.StackRule,
	}
}

func (data SpellGroupStackRulesSlice) ToDB() DBSpellGroupStackRulesSlice {
	result := make(DBSpellGroupStackRulesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellGroupStackRulesSlice) ToWeb() SpellGroupStackRulesSlice {
	result := make(SpellGroupStackRulesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
