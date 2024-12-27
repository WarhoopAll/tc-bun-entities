package model

import "github.com/uptrace/bun"

type Updates struct {
	Name string `json:"name,omitempty"`
	Hash string `json:"hash,omitempty"`
	State string `json:"state,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Speed int `json:"speed,omitempty"`
}

type UpdatesSlice []Updates

type DBUpdates struct {
	bun.BaseModel `bun:"table:updates,alias:updates"`
	Name string `bun:"name"`
	Hash string `bun:"hash"`
	State string `bun:"state"`
	Timestamp time.Time `bun:"timestamp"`
	Speed int `bun:"speed"`
}

type DBUpdatesSlice []DBUpdates

func (entry *Updates) ToDB() *DBUpdates {
	if entry == nil {
		return nil
	}
	return &DBUpdates{
		Name: entry.Name,
		Hash: entry.Hash,
		State: entry.State,
		Timestamp: entry.Timestamp,
		Speed: entry.Speed,
	}
}

func (entry *DBUpdates) ToWeb() *Updates {
	if entry == nil {
		return nil
	}
	return &Updates{
		Name: entry.Name,
		Hash: entry.Hash,
		State: entry.State,
		Timestamp: entry.Timestamp,
		Speed: entry.Speed,
	}
}

func (data UpdatesSlice) ToDB() DBUpdatesSlice {
	result := make(DBUpdatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBUpdatesSlice) ToWeb() UpdatesSlice {
	result := make(UpdatesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
