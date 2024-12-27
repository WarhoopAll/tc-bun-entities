package model

import "github.com/uptrace/bun"

type CharacterAction struct {
	Guid int `json:"guid,omitempty"`
	Spec int8 `json:"spec,omitempty"`
	Button int8 `json:"button,omitempty"`
	Action int `json:"action,omitempty"`
	Type int8 `json:"type,omitempty"`
}

type CharacterActionSlice []CharacterAction

type DBCharacterAction struct {
	bun.BaseModel `bun:"table:character_action,alias:character_action"`
	Guid int `bun:"guid"`
	Spec int8 `bun:"spec"`
	Button int8 `bun:"button"`
	Action int `bun:"action"`
	Type int8 `bun:"type"`
}

type DBCharacterActionSlice []DBCharacterAction

func (entry *CharacterAction) ToDB() *DBCharacterAction {
	if entry == nil {
		return nil
	}
	return &DBCharacterAction{
		Guid: entry.Guid,
		Spec: entry.Spec,
		Button: entry.Button,
		Action: entry.Action,
		Type: entry.Type,
	}
}

func (entry *DBCharacterAction) ToWeb() *CharacterAction {
	if entry == nil {
		return nil
	}
	return &CharacterAction{
		Guid: entry.Guid,
		Spec: entry.Spec,
		Button: entry.Button,
		Action: entry.Action,
		Type: entry.Type,
	}
}

func (data CharacterActionSlice) ToDB() DBCharacterActionSlice {
	result := make(DBCharacterActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterActionSlice) ToWeb() CharacterActionSlice {
	result := make(CharacterActionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
