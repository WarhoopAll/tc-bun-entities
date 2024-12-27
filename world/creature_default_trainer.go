package model

import "github.com/uptrace/bun"

type CreatureDefaultTrainer struct {
	CreatureId int `json:"creatureid,omitempty"`
	TrainerId int `json:"trainerid,omitempty"`
}

type CreatureDefaultTrainerSlice []CreatureDefaultTrainer

type DBCreatureDefaultTrainer struct {
	bun.BaseModel `bun:"table:creature_default_trainer,alias:creature_default_trainer"`
	CreatureId int `bun:"CreatureId"`
	TrainerId int `bun:"TrainerId"`
}

type DBCreatureDefaultTrainerSlice []DBCreatureDefaultTrainer

func (entry *CreatureDefaultTrainer) ToDB() *DBCreatureDefaultTrainer {
	if entry == nil {
		return nil
	}
	return &DBCreatureDefaultTrainer{
		CreatureId: entry.CreatureId,
		TrainerId: entry.TrainerId,
	}
}

func (entry *DBCreatureDefaultTrainer) ToWeb() *CreatureDefaultTrainer {
	if entry == nil {
		return nil
	}
	return &CreatureDefaultTrainer{
		CreatureId: entry.CreatureId,
		TrainerId: entry.TrainerId,
	}
}

func (data CreatureDefaultTrainerSlice) ToDB() DBCreatureDefaultTrainerSlice {
	result := make(DBCreatureDefaultTrainerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureDefaultTrainerSlice) ToWeb() CreatureDefaultTrainerSlice {
	result := make(CreatureDefaultTrainerSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
