package model

import "github.com/uptrace/bun"

type PetLevelstats struct {
	CreatureEntry int `json:"creature_entry,omitempty"`
	Level int8 `json:"level,omitempty"`
	Hp int16 `json:"hp,omitempty"`
	Mana int16 `json:"mana,omitempty"`
	Armor int `json:"armor,omitempty"`
	Str int16 `json:"str,omitempty"`
	Agi int16 `json:"agi,omitempty"`
	Sta int16 `json:"sta,omitempty"`
	Inte int16 `json:"inte,omitempty"`
	Spi int16 `json:"spi,omitempty"`
	MinDmg int16 `json:"min_dmg,omitempty"`
	MaxDmg int16 `json:"max_dmg,omitempty"`
}

type PetLevelstatsSlice []PetLevelstats

type DBPetLevelstats struct {
	bun.BaseModel `bun:"table:pet_levelstats,alias:pet_levelstats"`
	CreatureEntry int `bun:"creature_entry"`
	Level int8 `bun:"level"`
	Hp int16 `bun:"hp"`
	Mana int16 `bun:"mana"`
	Armor int `bun:"armor"`
	Str int16 `bun:"str"`
	Agi int16 `bun:"agi"`
	Sta int16 `bun:"sta"`
	Inte int16 `bun:"inte"`
	Spi int16 `bun:"spi"`
	MinDmg int16 `bun:"min_dmg"`
	MaxDmg int16 `bun:"max_dmg"`
}

type DBPetLevelstatsSlice []DBPetLevelstats

func (entry *PetLevelstats) ToDB() *DBPetLevelstats {
	if entry == nil {
		return nil
	}
	return &DBPetLevelstats{
		CreatureEntry: entry.CreatureEntry,
		Level: entry.Level,
		Hp: entry.Hp,
		Mana: entry.Mana,
		Armor: entry.Armor,
		Str: entry.Str,
		Agi: entry.Agi,
		Sta: entry.Sta,
		Inte: entry.Inte,
		Spi: entry.Spi,
		MinDmg: entry.MinDmg,
		MaxDmg: entry.MaxDmg,
	}
}

func (entry *DBPetLevelstats) ToWeb() *PetLevelstats {
	if entry == nil {
		return nil
	}
	return &PetLevelstats{
		CreatureEntry: entry.CreatureEntry,
		Level: entry.Level,
		Hp: entry.Hp,
		Mana: entry.Mana,
		Armor: entry.Armor,
		Str: entry.Str,
		Agi: entry.Agi,
		Sta: entry.Sta,
		Inte: entry.Inte,
		Spi: entry.Spi,
		MinDmg: entry.MinDmg,
		MaxDmg: entry.MaxDmg,
	}
}

func (data PetLevelstatsSlice) ToDB() DBPetLevelstatsSlice {
	result := make(DBPetLevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetLevelstatsSlice) ToWeb() PetLevelstatsSlice {
	result := make(PetLevelstatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
