package model

import "github.com/uptrace/bun"

type CharacterAccountData struct {
	Guid int `json:"guid,omitempty"`
	Type int8 `json:"type,omitempty"`
	Time int `json:"time,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type CharacterAccountDataSlice []CharacterAccountData

type DBCharacterAccountData struct {
	bun.BaseModel `bun:"table:character_account_data,alias:character_account_data"`
	Guid int `bun:"guid"`
	Type int8 `bun:"type"`
	Time int `bun:"time"`
	Data []byte `bun:"data"`
}

type DBCharacterAccountDataSlice []DBCharacterAccountData

func (entry *CharacterAccountData) ToDB() *DBCharacterAccountData {
	if entry == nil {
		return nil
	}
	return &DBCharacterAccountData{
		Guid: entry.Guid,
		Type: entry.Type,
		Time: entry.Time,
		Data: entry.Data,
	}
}

func (entry *DBCharacterAccountData) ToWeb() *CharacterAccountData {
	if entry == nil {
		return nil
	}
	return &CharacterAccountData{
		Guid: entry.Guid,
		Type: entry.Type,
		Time: entry.Time,
		Data: entry.Data,
	}
}

func (data CharacterAccountDataSlice) ToDB() DBCharacterAccountDataSlice {
	result := make(DBCharacterAccountDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterAccountDataSlice) ToWeb() CharacterAccountDataSlice {
	result := make(CharacterAccountDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
