package model

import "github.com/uptrace/bun"

type UpdatesInclude struct {
	Path string `json:"path,omitempty"`
	State string `json:"state,omitempty"`
}

type UpdatesIncludeSlice []UpdatesInclude

type DBUpdatesInclude struct {
	bun.BaseModel `bun:"table:updates_include,alias:updates_include"`
	Path string `bun:"path"`
	State string `bun:"state"`
}

type DBUpdatesIncludeSlice []DBUpdatesInclude

func (entry *UpdatesInclude) ToDB() *DBUpdatesInclude {
	if entry == nil {
		return nil
	}
	return &DBUpdatesInclude{
		Path: entry.Path,
		State: entry.State,
	}
}

func (entry *DBUpdatesInclude) ToWeb() *UpdatesInclude {
	if entry == nil {
		return nil
	}
	return &UpdatesInclude{
		Path: entry.Path,
		State: entry.State,
	}
}

func (data UpdatesIncludeSlice) ToDB() DBUpdatesIncludeSlice {
	result := make(DBUpdatesIncludeSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBUpdatesIncludeSlice) ToWeb() UpdatesIncludeSlice {
	result := make(UpdatesIncludeSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
