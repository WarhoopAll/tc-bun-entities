package model

import "github.com/uptrace/bun"

type BuildInfo struct {
	Build int `json:"build,omitempty"`
	MajorVersion int `json:"majorversion,omitempty"`
	MinorVersion int `json:"minorversion,omitempty"`
	BugfixVersion int `json:"bugfixversion,omitempty"`
	HotfixVersion string `json:"hotfixversion,omitempty"`
}

type BuildInfoSlice []BuildInfo

type DBBuildInfo struct {
	bun.BaseModel `bun:"table:build_info,alias:build_info"`
	Build int `bun:"build"`
	MajorVersion int `bun:"majorVersion"`
	MinorVersion int `bun:"minorVersion"`
	BugfixVersion int `bun:"bugfixVersion"`
	HotfixVersion string `bun:"hotfixVersion"`
}

type DBBuildInfoSlice []DBBuildInfo

func (entry *BuildInfo) ToDB() *DBBuildInfo {
	if entry == nil {
		return nil
	}
	return &DBBuildInfo{
		Build: entry.Build,
		MajorVersion: entry.MajorVersion,
		MinorVersion: entry.MinorVersion,
		BugfixVersion: entry.BugfixVersion,
		HotfixVersion: entry.HotfixVersion,
	}
}

func (entry *DBBuildInfo) ToWeb() *BuildInfo {
	if entry == nil {
		return nil
	}
	return &BuildInfo{
		Build: entry.Build,
		MajorVersion: entry.MajorVersion,
		MinorVersion: entry.MinorVersion,
		BugfixVersion: entry.BugfixVersion,
		HotfixVersion: entry.HotfixVersion,
	}
}

func (data BuildInfoSlice) ToDB() DBBuildInfoSlice {
	result := make(DBBuildInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBuildInfoSlice) ToWeb() BuildInfoSlice {
	result := make(BuildInfoSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
