package model

import "github.com/uptrace/bun"

type PetNameGeneration struct {
	Id int `json:"id,omitempty"`
	Word string `json:"word,omitempty"`
	Entry int `json:"entry,omitempty"`
	Half int8 `json:"half,omitempty"`
}

type PetNameGenerationSlice []PetNameGeneration

type DBPetNameGeneration struct {
	bun.BaseModel `bun:"table:pet_name_generation,alias:pet_name_generation"`
	Id int `bun:"id"`
	Word string `bun:"word"`
	Entry int `bun:"entry"`
	Half int8 `bun:"half"`
}

type DBPetNameGenerationSlice []DBPetNameGeneration

func (entry *PetNameGeneration) ToDB() *DBPetNameGeneration {
	if entry == nil {
		return nil
	}
	return &DBPetNameGeneration{
		Id: entry.Id,
		Word: entry.Word,
		Entry: entry.Entry,
		Half: entry.Half,
	}
}

func (entry *DBPetNameGeneration) ToWeb() *PetNameGeneration {
	if entry == nil {
		return nil
	}
	return &PetNameGeneration{
		Id: entry.Id,
		Word: entry.Word,
		Entry: entry.Entry,
		Half: entry.Half,
	}
}

func (data PetNameGenerationSlice) ToDB() DBPetNameGenerationSlice {
	result := make(DBPetNameGenerationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetNameGenerationSlice) ToWeb() PetNameGenerationSlice {
	result := make(PetNameGenerationSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
