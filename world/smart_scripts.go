package model

import "github.com/uptrace/bun"

type SmartScripts struct {
	Entryorguid int `json:"entryorguid,omitempty"`
	SourceType int8 `json:"source_type,omitempty"`
	Id int16 `json:"id,omitempty"`
	Link int16 `json:"link,omitempty"`
	EventType int8 `json:"event_type,omitempty"`
	EventPhaseMask int16 `json:"event_phase_mask,omitempty"`
	EventChance int8 `json:"event_chance,omitempty"`
	EventFlags int16 `json:"event_flags,omitempty"`
	EventParam1 int `json:"event_param1,omitempty"`
	EventParam2 int `json:"event_param2,omitempty"`
	EventParam3 int `json:"event_param3,omitempty"`
	EventParam4 int `json:"event_param4,omitempty"`
	EventParam5 int `json:"event_param5,omitempty"`
	ActionType int8 `json:"action_type,omitempty"`
	ActionParam1 int `json:"action_param1,omitempty"`
	ActionParam2 int `json:"action_param2,omitempty"`
	ActionParam3 int `json:"action_param3,omitempty"`
	ActionParam4 int `json:"action_param4,omitempty"`
	ActionParam5 int `json:"action_param5,omitempty"`
	ActionParam6 int `json:"action_param6,omitempty"`
	TargetType int8 `json:"target_type,omitempty"`
	TargetParam1 int `json:"target_param1,omitempty"`
	TargetParam2 int `json:"target_param2,omitempty"`
	TargetParam3 int `json:"target_param3,omitempty"`
	TargetParam4 int `json:"target_param4,omitempty"`
	TargetX float64 `json:"target_x,omitempty"`
	TargetY float64 `json:"target_y,omitempty"`
	TargetZ float64 `json:"target_z,omitempty"`
	TargetO float64 `json:"target_o,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type SmartScriptsSlice []SmartScripts

type DBSmartScripts struct {
	bun.BaseModel `bun:"table:smart_scripts,alias:smart_scripts"`
	Entryorguid int `bun:"entryorguid"`
	SourceType int8 `bun:"source_type"`
	Id int16 `bun:"id"`
	Link int16 `bun:"link"`
	EventType int8 `bun:"event_type"`
	EventPhaseMask int16 `bun:"event_phase_mask"`
	EventChance int8 `bun:"event_chance"`
	EventFlags int16 `bun:"event_flags"`
	EventParam1 int `bun:"event_param1"`
	EventParam2 int `bun:"event_param2"`
	EventParam3 int `bun:"event_param3"`
	EventParam4 int `bun:"event_param4"`
	EventParam5 int `bun:"event_param5"`
	ActionType int8 `bun:"action_type"`
	ActionParam1 int `bun:"action_param1"`
	ActionParam2 int `bun:"action_param2"`
	ActionParam3 int `bun:"action_param3"`
	ActionParam4 int `bun:"action_param4"`
	ActionParam5 int `bun:"action_param5"`
	ActionParam6 int `bun:"action_param6"`
	TargetType int8 `bun:"target_type"`
	TargetParam1 int `bun:"target_param1"`
	TargetParam2 int `bun:"target_param2"`
	TargetParam3 int `bun:"target_param3"`
	TargetParam4 int `bun:"target_param4"`
	TargetX float64 `bun:"target_x"`
	TargetY float64 `bun:"target_y"`
	TargetZ float64 `bun:"target_z"`
	TargetO float64 `bun:"target_o"`
	Comment string `bun:"comment"`
}

type DBSmartScriptsSlice []DBSmartScripts

func (entry *SmartScripts) ToDB() *DBSmartScripts {
	if entry == nil {
		return nil
	}
	return &DBSmartScripts{
		Entryorguid: entry.Entryorguid,
		SourceType: entry.SourceType,
		Id: entry.Id,
		Link: entry.Link,
		EventType: entry.EventType,
		EventPhaseMask: entry.EventPhaseMask,
		EventChance: entry.EventChance,
		EventFlags: entry.EventFlags,
		EventParam1: entry.EventParam1,
		EventParam2: entry.EventParam2,
		EventParam3: entry.EventParam3,
		EventParam4: entry.EventParam4,
		EventParam5: entry.EventParam5,
		ActionType: entry.ActionType,
		ActionParam1: entry.ActionParam1,
		ActionParam2: entry.ActionParam2,
		ActionParam3: entry.ActionParam3,
		ActionParam4: entry.ActionParam4,
		ActionParam5: entry.ActionParam5,
		ActionParam6: entry.ActionParam6,
		TargetType: entry.TargetType,
		TargetParam1: entry.TargetParam1,
		TargetParam2: entry.TargetParam2,
		TargetParam3: entry.TargetParam3,
		TargetParam4: entry.TargetParam4,
		TargetX: entry.TargetX,
		TargetY: entry.TargetY,
		TargetZ: entry.TargetZ,
		TargetO: entry.TargetO,
		Comment: entry.Comment,
	}
}

func (entry *DBSmartScripts) ToWeb() *SmartScripts {
	if entry == nil {
		return nil
	}
	return &SmartScripts{
		Entryorguid: entry.Entryorguid,
		SourceType: entry.SourceType,
		Id: entry.Id,
		Link: entry.Link,
		EventType: entry.EventType,
		EventPhaseMask: entry.EventPhaseMask,
		EventChance: entry.EventChance,
		EventFlags: entry.EventFlags,
		EventParam1: entry.EventParam1,
		EventParam2: entry.EventParam2,
		EventParam3: entry.EventParam3,
		EventParam4: entry.EventParam4,
		EventParam5: entry.EventParam5,
		ActionType: entry.ActionType,
		ActionParam1: entry.ActionParam1,
		ActionParam2: entry.ActionParam2,
		ActionParam3: entry.ActionParam3,
		ActionParam4: entry.ActionParam4,
		ActionParam5: entry.ActionParam5,
		ActionParam6: entry.ActionParam6,
		TargetType: entry.TargetType,
		TargetParam1: entry.TargetParam1,
		TargetParam2: entry.TargetParam2,
		TargetParam3: entry.TargetParam3,
		TargetParam4: entry.TargetParam4,
		TargetX: entry.TargetX,
		TargetY: entry.TargetY,
		TargetZ: entry.TargetZ,
		TargetO: entry.TargetO,
		Comment: entry.Comment,
	}
}

func (data SmartScriptsSlice) ToDB() DBSmartScriptsSlice {
	result := make(DBSmartScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSmartScriptsSlice) ToWeb() SmartScriptsSlice {
	result := make(SmartScriptsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
