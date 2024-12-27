package model

import "github.com/uptrace/bun"

type CharacterPet struct {
	Id int `json:"id,omitempty"`
	Entry int `json:"entry,omitempty"`
	Owner int `json:"owner,omitempty"`
	Modelid int `json:"modelid,omitempty"`
	CreatedBySpell int32 `json:"createdbyspell,omitempty"`
	PetType int8 `json:"pettype,omitempty"`
	Level int16 `json:"level,omitempty"`
	Exp int `json:"exp,omitempty"`
	Reactstate int8 `json:"reactstate,omitempty"`
	Name string `json:"name,omitempty"`
	Renamed int8 `json:"renamed,omitempty"`
	Slot int8 `json:"slot,omitempty"`
	Curhealth int `json:"curhealth,omitempty"`
	Curmana int `json:"curmana,omitempty"`
	Curhappiness int `json:"curhappiness,omitempty"`
	Savetime int `json:"savetime,omitempty"`
	Abdata string `json:"abdata,omitempty"`
}

type CharacterPetSlice []CharacterPet

type DBCharacterPet struct {
	bun.BaseModel `bun:"table:character_pet,alias:character_pet"`
	Id int `bun:"id"`
	Entry int `bun:"entry"`
	Owner int `bun:"owner"`
	Modelid int `bun:"modelid"`
	CreatedBySpell int32 `bun:"CreatedBySpell"`
	PetType int8 `bun:"PetType"`
	Level int16 `bun:"level"`
	Exp int `bun:"exp"`
	Reactstate int8 `bun:"Reactstate"`
	Name string `bun:"name"`
	Renamed int8 `bun:"renamed"`
	Slot int8 `bun:"slot"`
	Curhealth int `bun:"curhealth"`
	Curmana int `bun:"curmana"`
	Curhappiness int `bun:"curhappiness"`
	Savetime int `bun:"savetime"`
	Abdata string `bun:"abdata"`
}

type DBCharacterPetSlice []DBCharacterPet

func (entry *CharacterPet) ToDB() *DBCharacterPet {
	if entry == nil {
		return nil
	}
	return &DBCharacterPet{
		Id: entry.Id,
		Entry: entry.Entry,
		Owner: entry.Owner,
		Modelid: entry.Modelid,
		CreatedBySpell: entry.CreatedBySpell,
		PetType: entry.PetType,
		Level: entry.Level,
		Exp: entry.Exp,
		Reactstate: entry.Reactstate,
		Name: entry.Name,
		Renamed: entry.Renamed,
		Slot: entry.Slot,
		Curhealth: entry.Curhealth,
		Curmana: entry.Curmana,
		Curhappiness: entry.Curhappiness,
		Savetime: entry.Savetime,
		Abdata: entry.Abdata,
	}
}

func (entry *DBCharacterPet) ToWeb() *CharacterPet {
	if entry == nil {
		return nil
	}
	return &CharacterPet{
		Id: entry.Id,
		Entry: entry.Entry,
		Owner: entry.Owner,
		Modelid: entry.Modelid,
		CreatedBySpell: entry.CreatedBySpell,
		PetType: entry.PetType,
		Level: entry.Level,
		Exp: entry.Exp,
		Reactstate: entry.Reactstate,
		Name: entry.Name,
		Renamed: entry.Renamed,
		Slot: entry.Slot,
		Curhealth: entry.Curhealth,
		Curmana: entry.Curmana,
		Curhappiness: entry.Curhappiness,
		Savetime: entry.Savetime,
		Abdata: entry.Abdata,
	}
}

func (data CharacterPetSlice) ToDB() DBCharacterPetSlice {
	result := make(DBCharacterPetSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterPetSlice) ToWeb() CharacterPetSlice {
	result := make(CharacterPetSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
