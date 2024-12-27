package model

import "github.com/uptrace/bun"

type PetSpellCooldown struct {
	Guid int `json:"guid,omitempty"`
	Spell int32 `json:"spell,omitempty"`
	Time int `json:"time,omitempty"`
	CategoryId int `json:"categoryid,omitempty"`
	CategoryEnd int `json:"categoryend,omitempty"`
}

type PetSpellCooldownSlice []PetSpellCooldown

type DBPetSpellCooldown struct {
	bun.BaseModel `bun:"table:pet_spell_cooldown,alias:pet_spell_cooldown"`
	Guid int `bun:"guid"`
	Spell int32 `bun:"spell"`
	Time int `bun:"time"`
	CategoryId int `bun:"categoryId"`
	CategoryEnd int `bun:"categoryEnd"`
}

type DBPetSpellCooldownSlice []DBPetSpellCooldown

func (entry *PetSpellCooldown) ToDB() *DBPetSpellCooldown {
	if entry == nil {
		return nil
	}
	return &DBPetSpellCooldown{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Time: entry.Time,
		CategoryId: entry.CategoryId,
		CategoryEnd: entry.CategoryEnd,
	}
}

func (entry *DBPetSpellCooldown) ToWeb() *PetSpellCooldown {
	if entry == nil {
		return nil
	}
	return &PetSpellCooldown{
		Guid: entry.Guid,
		Spell: entry.Spell,
		Time: entry.Time,
		CategoryId: entry.CategoryId,
		CategoryEnd: entry.CategoryEnd,
	}
}

func (data PetSpellCooldownSlice) ToDB() DBPetSpellCooldownSlice {
	result := make(DBPetSpellCooldownSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetSpellCooldownSlice) ToWeb() PetSpellCooldownSlice {
	result := make(PetSpellCooldownSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
