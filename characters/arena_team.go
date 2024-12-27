package model

import "github.com/uptrace/bun"

type ArenaTeam struct {
	ArenaTeamId int `json:"arenateamid,omitempty"`
	Name string `json:"name,omitempty"`
	CaptainGuid int `json:"captainguid,omitempty"`
	Type int8 `json:"type,omitempty"`
	Rating int16 `json:"rating,omitempty"`
	SeasonGames int16 `json:"seasongames,omitempty"`
	SeasonWins int16 `json:"seasonwins,omitempty"`
	WeekGames int16 `json:"weekgames,omitempty"`
	WeekWins int16 `json:"weekwins,omitempty"`
	Rank int `json:"rank,omitempty"`
	BackgroundColor int `json:"backgroundcolor,omitempty"`
	EmblemStyle int8 `json:"emblemstyle,omitempty"`
	EmblemColor int `json:"emblemcolor,omitempty"`
	BorderStyle int8 `json:"borderstyle,omitempty"`
	BorderColor int `json:"bordercolor,omitempty"`
}

type ArenaTeamSlice []ArenaTeam

type DBArenaTeam struct {
	bun.BaseModel `bun:"table:arena_team,alias:arena_team"`
	ArenaTeamId int `bun:"arenaTeamId"`
	Name string `bun:"name"`
	CaptainGuid int `bun:"captainGuid"`
	Type int8 `bun:"type"`
	Rating int16 `bun:"rating"`
	SeasonGames int16 `bun:"seasonGames"`
	SeasonWins int16 `bun:"seasonWins"`
	WeekGames int16 `bun:"weekGames"`
	WeekWins int16 `bun:"weekWins"`
	Rank int `bun:"rank"`
	BackgroundColor int `bun:"backgroundColor"`
	EmblemStyle int8 `bun:"emblemStyle"`
	EmblemColor int `bun:"emblemColor"`
	BorderStyle int8 `bun:"borderStyle"`
	BorderColor int `bun:"borderColor"`
}

type DBArenaTeamSlice []DBArenaTeam

func (entry *ArenaTeam) ToDB() *DBArenaTeam {
	if entry == nil {
		return nil
	}
	return &DBArenaTeam{
		ArenaTeamId: entry.ArenaTeamId,
		Name: entry.Name,
		CaptainGuid: entry.CaptainGuid,
		Type: entry.Type,
		Rating: entry.Rating,
		SeasonGames: entry.SeasonGames,
		SeasonWins: entry.SeasonWins,
		WeekGames: entry.WeekGames,
		WeekWins: entry.WeekWins,
		Rank: entry.Rank,
		BackgroundColor: entry.BackgroundColor,
		EmblemStyle: entry.EmblemStyle,
		EmblemColor: entry.EmblemColor,
		BorderStyle: entry.BorderStyle,
		BorderColor: entry.BorderColor,
	}
}

func (entry *DBArenaTeam) ToWeb() *ArenaTeam {
	if entry == nil {
		return nil
	}
	return &ArenaTeam{
		ArenaTeamId: entry.ArenaTeamId,
		Name: entry.Name,
		CaptainGuid: entry.CaptainGuid,
		Type: entry.Type,
		Rating: entry.Rating,
		SeasonGames: entry.SeasonGames,
		SeasonWins: entry.SeasonWins,
		WeekGames: entry.WeekGames,
		WeekWins: entry.WeekWins,
		Rank: entry.Rank,
		BackgroundColor: entry.BackgroundColor,
		EmblemStyle: entry.EmblemStyle,
		EmblemColor: entry.EmblemColor,
		BorderStyle: entry.BorderStyle,
		BorderColor: entry.BorderColor,
	}
}

func (data ArenaTeamSlice) ToDB() DBArenaTeamSlice {
	result := make(DBArenaTeamSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBArenaTeamSlice) ToWeb() ArenaTeamSlice {
	result := make(ArenaTeamSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
