package model

import "github.com/uptrace/bun"

type Petition struct {
	Ownerguid int `json:"ownerguid,omitempty"`
	Petitionguid int `json:"petitionguid,omitempty"`
	Name string `json:"name,omitempty"`
	Type int8 `json:"type,omitempty"`
}

type PetitionSlice []Petition

type DBPetition struct {
	bun.BaseModel `bun:"table:petition,alias:petition"`
	Ownerguid int `bun:"ownerguid"`
	Petitionguid int `bun:"petitionguid"`
	Name string `bun:"name"`
	Type int8 `bun:"type"`
}

type DBPetitionSlice []DBPetition

func (entry *Petition) ToDB() *DBPetition {
	if entry == nil {
		return nil
	}
	return &DBPetition{
		Ownerguid: entry.Ownerguid,
		Petitionguid: entry.Petitionguid,
		Name: entry.Name,
		Type: entry.Type,
	}
}

func (entry *DBPetition) ToWeb() *Petition {
	if entry == nil {
		return nil
	}
	return &Petition{
		Ownerguid: entry.Ownerguid,
		Petitionguid: entry.Petitionguid,
		Name: entry.Name,
		Type: entry.Type,
	}
}

func (data PetitionSlice) ToDB() DBPetitionSlice {
	result := make(DBPetitionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetitionSlice) ToWeb() PetitionSlice {
	result := make(PetitionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
