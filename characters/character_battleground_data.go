package model

import "github.com/uptrace/bun"

type CharacterBattlegroundData struct {
	Guid int `json:"guid,omitempty"`
	InstanceId int `json:"instanceid,omitempty"`
	Team int16 `json:"team,omitempty"`
	JoinX float64 `json:"joinx,omitempty"`
	JoinY float64 `json:"joiny,omitempty"`
	JoinZ float64 `json:"joinz,omitempty"`
	JoinO float64 `json:"joino,omitempty"`
	JoinMapId int16 `json:"joinmapid,omitempty"`
	TaxiStart int `json:"taxistart,omitempty"`
	TaxiEnd int `json:"taxiend,omitempty"`
	MountSpell int `json:"mountspell,omitempty"`
}

type CharacterBattlegroundDataSlice []CharacterBattlegroundData

type DBCharacterBattlegroundData struct {
	bun.BaseModel `bun:"table:character_battleground_data,alias:character_battleground_data"`
	Guid int `bun:"guid"`
	InstanceId int `bun:"instanceId"`
	Team int16 `bun:"team"`
	JoinX float64 `bun:"joinX"`
	JoinY float64 `bun:"joinY"`
	JoinZ float64 `bun:"joinZ"`
	JoinO float64 `bun:"joinO"`
	JoinMapId int16 `bun:"joinMapId"`
	TaxiStart int `bun:"taxiStart"`
	TaxiEnd int `bun:"taxiEnd"`
	MountSpell int `bun:"mountSpell"`
}

type DBCharacterBattlegroundDataSlice []DBCharacterBattlegroundData

func (entry *CharacterBattlegroundData) ToDB() *DBCharacterBattlegroundData {
	if entry == nil {
		return nil
	}
	return &DBCharacterBattlegroundData{
		Guid: entry.Guid,
		InstanceId: entry.InstanceId,
		Team: entry.Team,
		JoinX: entry.JoinX,
		JoinY: entry.JoinY,
		JoinZ: entry.JoinZ,
		JoinO: entry.JoinO,
		JoinMapId: entry.JoinMapId,
		TaxiStart: entry.TaxiStart,
		TaxiEnd: entry.TaxiEnd,
		MountSpell: entry.MountSpell,
	}
}

func (entry *DBCharacterBattlegroundData) ToWeb() *CharacterBattlegroundData {
	if entry == nil {
		return nil
	}
	return &CharacterBattlegroundData{
		Guid: entry.Guid,
		InstanceId: entry.InstanceId,
		Team: entry.Team,
		JoinX: entry.JoinX,
		JoinY: entry.JoinY,
		JoinZ: entry.JoinZ,
		JoinO: entry.JoinO,
		JoinMapId: entry.JoinMapId,
		TaxiStart: entry.TaxiStart,
		TaxiEnd: entry.TaxiEnd,
		MountSpell: entry.MountSpell,
	}
}

func (data CharacterBattlegroundDataSlice) ToDB() DBCharacterBattlegroundDataSlice {
	result := make(DBCharacterBattlegroundDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterBattlegroundDataSlice) ToWeb() CharacterBattlegroundDataSlice {
	result := make(CharacterBattlegroundDataSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
