package model

import "github.com/uptrace/bun"

type Trainer struct {
	Id int `json:"id,omitempty"`
	Type int8 `json:"type,omitempty"`
	Requirement int `json:"requirement,omitempty"`
	Greeting string `json:"greeting,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type TrainerSlice []Trainer

type DBTrainer struct {
	bun.BaseModel `bun:"table:trainer,alias:trainer"`
	Id int `bun:"Id"`
	Type int8 `bun:"Type"`
	Requirement int `bun:"Requirement"`
	Greeting string `bun:"Greeting"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBTrainerSlice []DBTrainer

func (entry *Trainer) ToDB() *DBTrainer {
	if entry == nil {
		return nil
	}
	return &DBTrainer{
		Id: entry.Id,
		Type: entry.Type,
		Requirement: entry.Requirement,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBTrainer) ToWeb() *Trainer {
	if entry == nil {
		return nil
	}
	return &Trainer{
		Id: entry.Id,
		Type: entry.Type,
		Requirement: entry.Requirement,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data TrainerSlice) ToDB() DBTrainerSlice {
	result := make(DBTrainerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBTrainerSlice) ToWeb() TrainerSlice {
	result := make(TrainerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
