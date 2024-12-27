package model

import "github.com/uptrace/bun"

type PetitionSign struct {
	Ownerguid int `json:"ownerguid,omitempty"`
	Petitionguid int `json:"petitionguid,omitempty"`
	Playerguid int `json:"playerguid,omitempty"`
	PlayerAccount int `json:"player_account,omitempty"`
	Type int8 `json:"type,omitempty"`
}

type PetitionSignSlice []PetitionSign

type DBPetitionSign struct {
	bun.BaseModel `bun:"table:petition_sign,alias:petition_sign"`
	Ownerguid int `bun:"ownerguid"`
	Petitionguid int `bun:"petitionguid"`
	Playerguid int `bun:"playerguid"`
	PlayerAccount int `bun:"player_account"`
	Type int8 `bun:"type"`
}

type DBPetitionSignSlice []DBPetitionSign

func (entry *PetitionSign) ToDB() *DBPetitionSign {
	if entry == nil {
		return nil
	}
	return &DBPetitionSign{
		Ownerguid: entry.Ownerguid,
		Petitionguid: entry.Petitionguid,
		Playerguid: entry.Playerguid,
		PlayerAccount: entry.PlayerAccount,
		Type: entry.Type,
	}
}

func (entry *DBPetitionSign) ToWeb() *PetitionSign {
	if entry == nil {
		return nil
	}
	return &PetitionSign{
		Ownerguid: entry.Ownerguid,
		Petitionguid: entry.Petitionguid,
		Playerguid: entry.Playerguid,
		PlayerAccount: entry.PlayerAccount,
		Type: entry.Type,
	}
}

func (data PetitionSignSlice) ToDB() DBPetitionSignSlice {
	result := make(DBPetitionSignSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBPetitionSignSlice) ToWeb() PetitionSignSlice {
	result := make(PetitionSignSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
