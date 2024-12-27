package model

import "github.com/uptrace/bun"

type CharacterSocial struct {
	Guid int `json:"guid,omitempty"`
	Friend int `json:"friend,omitempty"`
	Flags int8 `json:"flags,omitempty"`
	Note string `json:"note,omitempty"`
}

type CharacterSocialSlice []CharacterSocial

type DBCharacterSocial struct {
	bun.BaseModel `bun:"table:character_social,alias:character_social"`
	Guid int `bun:"guid"`
	Friend int `bun:"friend"`
	Flags int8 `bun:"flags"`
	Note string `bun:"note"`
}

type DBCharacterSocialSlice []DBCharacterSocial

func (entry *CharacterSocial) ToDB() *DBCharacterSocial {
	if entry == nil {
		return nil
	}
	return &DBCharacterSocial{
		Guid: entry.Guid,
		Friend: entry.Friend,
		Flags: entry.Flags,
		Note: entry.Note,
	}
}

func (entry *DBCharacterSocial) ToWeb() *CharacterSocial {
	if entry == nil {
		return nil
	}
	return &CharacterSocial{
		Guid: entry.Guid,
		Friend: entry.Friend,
		Flags: entry.Flags,
		Note: entry.Note,
	}
}

func (data CharacterSocialSlice) ToDB() DBCharacterSocialSlice {
	result := make(DBCharacterSocialSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterSocialSlice) ToWeb() CharacterSocialSlice {
	result := make(CharacterSocialSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
