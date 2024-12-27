package model

import "github.com/uptrace/bun"

type CharacterInstance struct {
	Guid int `json:"guid,omitempty"`
	Instance int `json:"instance,omitempty"`
	Permanent int8 `json:"permanent,omitempty"`
	ExtendState int8 `json:"extendstate,omitempty"`
}

type CharacterInstanceSlice []CharacterInstance

type DBCharacterInstance struct {
	bun.BaseModel `bun:"table:character_instance,alias:character_instance"`
	Guid int `bun:"guid"`
	Instance int `bun:"instance"`
	Permanent int8 `bun:"permanent"`
	ExtendState int8 `bun:"extendState"`
}

type DBCharacterInstanceSlice []DBCharacterInstance

func (entry *CharacterInstance) ToDB() *DBCharacterInstance {
	if entry == nil {
		return nil
	}
	return &DBCharacterInstance{
		Guid: entry.Guid,
		Instance: entry.Instance,
		Permanent: entry.Permanent,
		ExtendState: entry.ExtendState,
	}
}

func (entry *DBCharacterInstance) ToWeb() *CharacterInstance {
	if entry == nil {
		return nil
	}
	return &CharacterInstance{
		Guid: entry.Guid,
		Instance: entry.Instance,
		Permanent: entry.Permanent,
		ExtendState: entry.ExtendState,
	}
}

func (data CharacterInstanceSlice) ToDB() DBCharacterInstanceSlice {
	result := make(DBCharacterInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterInstanceSlice) ToWeb() CharacterInstanceSlice {
	result := make(CharacterInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
