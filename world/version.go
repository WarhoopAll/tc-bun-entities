package model

import "github.com/uptrace/bun"

type Version struct {
	CoreVersion string `json:"core_version,omitempty"`
	CoreRevision string `json:"core_revision,omitempty"`
	DbVersion string `json:"db_version,omitempty"`
	CacheId int `json:"cache_id,omitempty"`
}

type VersionSlice []Version

type DBVersion struct {
	bun.BaseModel `bun:"table:version,alias:version"`
	CoreVersion string `bun:"core_version"`
	CoreRevision string `bun:"core_revision"`
	DbVersion string `bun:"db_version"`
	CacheId int `bun:"cache_id"`
}

type DBVersionSlice []DBVersion

func (entry *Version) ToDB() *DBVersion {
	if entry == nil {
		return nil
	}
	return &DBVersion{
		CoreVersion: entry.CoreVersion,
		CoreRevision: entry.CoreRevision,
		DbVersion: entry.DbVersion,
		CacheId: entry.CacheId,
	}
}

func (entry *DBVersion) ToWeb() *Version {
	if entry == nil {
		return nil
	}
	return &Version{
		CoreVersion: entry.CoreVersion,
		CoreRevision: entry.CoreRevision,
		DbVersion: entry.DbVersion,
		CacheId: entry.CacheId,
	}
}

func (data VersionSlice) ToDB() DBVersionSlice {
	result := make(DBVersionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBVersionSlice) ToWeb() VersionSlice {
	result := make(VersionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
