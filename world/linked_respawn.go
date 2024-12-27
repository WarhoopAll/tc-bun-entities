package model

import "github.com/uptrace/bun"

type LinkedRespawn struct {
	Guid int `json:"guid,omitempty"`
	LinkedGuid int `json:"linkedguid,omitempty"`
	LinkType int8 `json:"linktype,omitempty"`
}

type LinkedRespawnSlice []LinkedRespawn

type DBLinkedRespawn struct {
	bun.BaseModel `bun:"table:linked_respawn,alias:linked_respawn"`
	Guid int `bun:"guid"`
	LinkedGuid int `bun:"linkedGuid"`
	LinkType int8 `bun:"linkType"`
}

type DBLinkedRespawnSlice []DBLinkedRespawn

func (entry *LinkedRespawn) ToDB() *DBLinkedRespawn {
	if entry == nil {
		return nil
	}
	return &DBLinkedRespawn{
		Guid: entry.Guid,
		LinkedGuid: entry.LinkedGuid,
		LinkType: entry.LinkType,
	}
}

func (entry *DBLinkedRespawn) ToWeb() *LinkedRespawn {
	if entry == nil {
		return nil
	}
	return &LinkedRespawn{
		Guid: entry.Guid,
		LinkedGuid: entry.LinkedGuid,
		LinkType: entry.LinkType,
	}
}

func (data LinkedRespawnSlice) ToDB() DBLinkedRespawnSlice {
	result := make(DBLinkedRespawnSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBLinkedRespawnSlice) ToWeb() LinkedRespawnSlice {
	result := make(LinkedRespawnSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
