package model

import "github.com/uptrace/bun"

type Realmcharacters struct {
	Realmid int `json:"realmid,omitempty"`
	Acctid int `json:"acctid,omitempty"`
	Numchars int8 `json:"numchars,omitempty"`
}

type RealmcharactersSlice []Realmcharacters

type DBRealmcharacters struct {
	bun.BaseModel `bun:"table:realmcharacters,alias:realmcharacters"`
	Realmid int `bun:"realmid"`
	Acctid int `bun:"acctid"`
	Numchars int8 `bun:"numchars"`
}

type DBRealmcharactersSlice []DBRealmcharacters

func (entry *Realmcharacters) ToDB() *DBRealmcharacters {
	if entry == nil {
		return nil
	}
	return &DBRealmcharacters{
		Realmid: entry.Realmid,
		Acctid: entry.Acctid,
		Numchars: entry.Numchars,
	}
}

func (entry *DBRealmcharacters) ToWeb() *Realmcharacters {
	if entry == nil {
		return nil
	}
	return &Realmcharacters{
		Realmid: entry.Realmid,
		Acctid: entry.Acctid,
		Numchars: entry.Numchars,
	}
}

func (data RealmcharactersSlice) ToDB() DBRealmcharactersSlice {
	result := make(DBRealmcharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRealmcharactersSlice) ToWeb() RealmcharactersSlice {
	result := make(RealmcharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
