package model

import "github.com/uptrace/bun"

type GroupInstance struct {
	Guid int `json:"guid,omitempty"`
	Instance int `json:"instance,omitempty"`
	Permanent int8 `json:"permanent,omitempty"`
}

type GroupInstanceSlice []GroupInstance

type DBGroupInstance struct {
	bun.BaseModel `bun:"table:group_instance,alias:group_instance"`
	Guid int `bun:"guid"`
	Instance int `bun:"instance"`
	Permanent int8 `bun:"permanent"`
}

type DBGroupInstanceSlice []DBGroupInstance

func (entry *GroupInstance) ToDB() *DBGroupInstance {
	if entry == nil {
		return nil
	}
	return &DBGroupInstance{
		Guid: entry.Guid,
		Instance: entry.Instance,
		Permanent: entry.Permanent,
	}
}

func (entry *DBGroupInstance) ToWeb() *GroupInstance {
	if entry == nil {
		return nil
	}
	return &GroupInstance{
		Guid: entry.Guid,
		Instance: entry.Instance,
		Permanent: entry.Permanent,
	}
}

func (data GroupInstanceSlice) ToDB() DBGroupInstanceSlice {
	result := make(DBGroupInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGroupInstanceSlice) ToWeb() GroupInstanceSlice {
	result := make(GroupInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
