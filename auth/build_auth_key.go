package model

import "github.com/uptrace/bun"

type BuildAuthKey struct {
	Build int `json:"build,omitempty"`
	Platform string `json:"platform,omitempty"`
	Arch string `json:"arch,omitempty"`
	Type string `json:"type,omitempty"`
	Key []byte `json:"key,omitempty"`
}

type BuildAuthKeySlice []BuildAuthKey

type DBBuildAuthKey struct {
	bun.BaseModel `bun:"table:build_auth_key,alias:build_auth_key"`
	Build int `bun:"build"`
	Platform string `bun:"platform"`
	Arch string `bun:"arch"`
	Type string `bun:"type"`
	Key []byte `bun:"key"`
}

type DBBuildAuthKeySlice []DBBuildAuthKey

func (entry *BuildAuthKey) ToDB() *DBBuildAuthKey {
	if entry == nil {
		return nil
	}
	return &DBBuildAuthKey{
		Build: entry.Build,
		Platform: entry.Platform,
		Arch: entry.Arch,
		Type: entry.Type,
		Key: entry.Key,
	}
}

func (entry *DBBuildAuthKey) ToWeb() *BuildAuthKey {
	if entry == nil {
		return nil
	}
	return &BuildAuthKey{
		Build: entry.Build,
		Platform: entry.Platform,
		Arch: entry.Arch,
		Type: entry.Type,
		Key: entry.Key,
	}
}

func (data BuildAuthKeySlice) ToDB() DBBuildAuthKeySlice {
	result := make(DBBuildAuthKeySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBuildAuthKeySlice) ToWeb() BuildAuthKeySlice {
	result := make(BuildAuthKeySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
