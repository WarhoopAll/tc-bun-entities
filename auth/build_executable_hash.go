package model

import "github.com/uptrace/bun"

type BuildExecutableHash struct {
	Build int `json:"build,omitempty"`
	Platform string `json:"platform,omitempty"`
	ExecutableHash []byte `json:"executablehash,omitempty"`
}

type BuildExecutableHashSlice []BuildExecutableHash

type DBBuildExecutableHash struct {
	bun.BaseModel `bun:"table:build_executable_hash,alias:build_executable_hash"`
	Build int `bun:"build"`
	Platform string `bun:"platform"`
	ExecutableHash []byte `bun:"executableHash"`
}

type DBBuildExecutableHashSlice []DBBuildExecutableHash

func (entry *BuildExecutableHash) ToDB() *DBBuildExecutableHash {
	if entry == nil {
		return nil
	}
	return &DBBuildExecutableHash{
		Build: entry.Build,
		Platform: entry.Platform,
		ExecutableHash: entry.ExecutableHash,
	}
}

func (entry *DBBuildExecutableHash) ToWeb() *BuildExecutableHash {
	if entry == nil {
		return nil
	}
	return &BuildExecutableHash{
		Build: entry.Build,
		Platform: entry.Platform,
		ExecutableHash: entry.ExecutableHash,
	}
}

func (data BuildExecutableHashSlice) ToDB() DBBuildExecutableHashSlice {
	result := make(DBBuildExecutableHashSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBBuildExecutableHashSlice) ToWeb() BuildExecutableHashSlice {
	result := make(BuildExecutableHashSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
