package model

import "github.com/uptrace/bun"

type CreatureMovementOverride struct {
	SpawnId int `json:"spawnid,omitempty"`
	Ground int8 `json:"ground,omitempty"`
	Swim int8 `json:"swim,omitempty"`
	Flight int8 `json:"flight,omitempty"`
	Rooted int8 `json:"rooted,omitempty"`
	Chase int8 `json:"chase,omitempty"`
	Random int8 `json:"random,omitempty"`
	InteractionPauseTimer int `json:"interactionpausetimer,omitempty"`
}

type CreatureMovementOverrideSlice []CreatureMovementOverride

type DBCreatureMovementOverride struct {
	bun.BaseModel `bun:"table:creature_movement_override,alias:creature_movement_override"`
	SpawnId int `bun:"SpawnId"`
	Ground int8 `bun:"Ground"`
	Swim int8 `bun:"Swim"`
	Flight int8 `bun:"Flight"`
	Rooted int8 `bun:"Rooted"`
	Chase int8 `bun:"Chase"`
	Random int8 `bun:"Random"`
	InteractionPauseTimer int `bun:"InteractionPauseTimer"`
}

type DBCreatureMovementOverrideSlice []DBCreatureMovementOverride

func (entry *CreatureMovementOverride) ToDB() *DBCreatureMovementOverride {
	if entry == nil {
		return nil
	}
	return &DBCreatureMovementOverride{
		SpawnId: entry.SpawnId,
		Ground: entry.Ground,
		Swim: entry.Swim,
		Flight: entry.Flight,
		Rooted: entry.Rooted,
		Chase: entry.Chase,
		Random: entry.Random,
		InteractionPauseTimer: entry.InteractionPauseTimer,
	}
}

func (entry *DBCreatureMovementOverride) ToWeb() *CreatureMovementOverride {
	if entry == nil {
		return nil
	}
	return &CreatureMovementOverride{
		SpawnId: entry.SpawnId,
		Ground: entry.Ground,
		Swim: entry.Swim,
		Flight: entry.Flight,
		Rooted: entry.Rooted,
		Chase: entry.Chase,
		Random: entry.Random,
		InteractionPauseTimer: entry.InteractionPauseTimer,
	}
}

func (data CreatureMovementOverrideSlice) ToDB() DBCreatureMovementOverrideSlice {
	result := make(DBCreatureMovementOverrideSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureMovementOverrideSlice) ToWeb() CreatureMovementOverrideSlice {
	result := make(CreatureMovementOverrideSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
