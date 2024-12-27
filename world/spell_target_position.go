package model

import "github.com/uptrace/bun"

type SpellTargetPosition struct {
	ID int `json:"id,omitempty"`
	EffectIndex int8 `json:"effectindex,omitempty"`
	MapID int16 `json:"mapid,omitempty"`
	PositionX float64 `json:"positionx,omitempty"`
	PositionY float64 `json:"positiony,omitempty"`
	PositionZ float64 `json:"positionz,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type SpellTargetPositionSlice []SpellTargetPosition

type DBSpellTargetPosition struct {
	bun.BaseModel `bun:"table:spell_target_position,alias:spell_target_position"`
	ID int `bun:"ID"`
	EffectIndex int8 `bun:"EffectIndex"`
	MapID int16 `bun:"MapID"`
	PositionX float64 `bun:"PositionX"`
	PositionY float64 `bun:"PositionY"`
	PositionZ float64 `bun:"PositionZ"`
	Orientation float64 `bun:"Orientation"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBSpellTargetPositionSlice []DBSpellTargetPosition

func (entry *SpellTargetPosition) ToDB() *DBSpellTargetPosition {
	if entry == nil {
		return nil
	}
	return &DBSpellTargetPosition{
		ID: entry.ID,
		EffectIndex: entry.EffectIndex,
		MapID: entry.MapID,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBSpellTargetPosition) ToWeb() *SpellTargetPosition {
	if entry == nil {
		return nil
	}
	return &SpellTargetPosition{
		ID: entry.ID,
		EffectIndex: entry.EffectIndex,
		MapID: entry.MapID,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Orientation: entry.Orientation,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data SpellTargetPositionSlice) ToDB() DBSpellTargetPositionSlice {
	result := make(DBSpellTargetPositionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellTargetPositionSlice) ToWeb() SpellTargetPositionSlice {
	result := make(SpellTargetPositionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
