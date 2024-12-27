package model

import "github.com/uptrace/bun"

type CharacterSpellCooldown struct {
	Guid int `json:"guid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	Item int `json:"item,omitempty"`
	Time int `json:"time,omitempty"`
	CategoryId int `json:"categoryid,omitempty"`
	CategoryEnd int `json:"categoryend,omitempty"`
}

type CharacterSpellCooldownSlice []CharacterSpellCooldown

type DBCharacterSpellCooldown struct {
	bun.BaseModel `bun:"table:character_spell_cooldown,alias:character_spell_cooldown"`
	Guid int `bun:"guid"`
	Spell int32 `bun:"spell"`
	Item int `bun:"item"`
	Time int `bun:"time"`
	CategoryId int `bun:"categoryId"`
	CategoryEnd int `bun:"categoryEnd"`
}

type DBCharacterSpellCooldownSlice []DBCharacterSpellCooldown

func (entry *CharacterSpellCooldown) ToDB() *DBCharacterSpellCooldown {
	if entry == nil {
		return nil
	}
	return &DBCharacterSpellCooldown{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Item: entry.Item,
		Time: entry.Time,
		CategoryId: entry.CategoryId,
		CategoryEnd: entry.CategoryEnd,
	}
}

func (entry *DBCharacterSpellCooldown) ToWeb() *CharacterSpellCooldown {
	if entry == nil {
		return nil
	}
	return &CharacterSpellCooldown{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Item: entry.Item,
		Time: entry.Time,
		CategoryId: entry.CategoryId,
		CategoryEnd: entry.CategoryEnd,
	}
}

func (data CharacterSpellCooldownSlice) ToDB() DBCharacterSpellCooldownSlice {
	result := make(DBCharacterSpellCooldownSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterSpellCooldownSlice) ToWeb() CharacterSpellCooldownSlice {
	result := make(CharacterSpellCooldownSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
