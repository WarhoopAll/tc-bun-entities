package model

import "github.com/uptrace/bun"

type CreatureTemplateMovement struct {
	CreatureId int `json:"creatureid,omitempty"`
	Ground int8 `json:"ground,omitempty"`
	Swim int8 `json:"swim,omitempty"`
	Flight int8 `json:"flight,omitempty"`
	Rooted int8 `json:"rooted,omitempty"`
	Chase int8 `json:"chase,omitempty"`
	Random int8 `json:"random,omitempty"`
	InteractionPauseTimer int `json:"interactionpausetimer,omitempty"`
}

type CreatureTemplateMovementSlice []CreatureTemplateMovement

type DBCreatureTemplateMovement struct {
	bun.BaseModel `bun:"table:creature_template_movement,alias:creature_template_movement"`
	CreatureId int `bun:"CreatureId"`
	Ground int8 `bun:"Ground"`
	Swim int8 `bun:"Swim"`
	Flight int8 `bun:"Flight"`
	Rooted int8 `bun:"Rooted"`
	Chase int8 `bun:"Chase"`
	Random int8 `bun:"Random"`
	InteractionPauseTimer int `bun:"InteractionPauseTimer"`
}

type DBCreatureTemplateMovementSlice []DBCreatureTemplateMovement

func (entry *CreatureTemplateMovement) ToDB() *DBCreatureTemplateMovement {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplateMovement{
		CreatureId: entry.CreatureId,
		Ground: entry.Ground,
		Swim: entry.Swim,
		Flight: entry.Flight,
		Rooted: entry.Rooted,
		Chase: entry.Chase,
		Random: entry.Random,
		InteractionPauseTimer: entry.InteractionPauseTimer,
	}
}

func (entry *DBCreatureTemplateMovement) ToWeb() *CreatureTemplateMovement {
	if entry == nil {
		return nil
	}
	return &CreatureTemplateMovement{
		CreatureId: entry.CreatureId,
		Ground: entry.Ground,
		Swim: entry.Swim,
		Flight: entry.Flight,
		Rooted: entry.Rooted,
		Chase: entry.Chase,
		Random: entry.Random,
		InteractionPauseTimer: entry.InteractionPauseTimer,
	}
}

func (data CreatureTemplateMovementSlice) ToDB() DBCreatureTemplateMovementSlice {
	result := make(DBCreatureTemplateMovementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateMovementSlice) ToWeb() CreatureTemplateMovementSlice {
	result := make(CreatureTemplateMovementSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
