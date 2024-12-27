package model

import "github.com/uptrace/bun"

type ArenaTeamMember struct {
	ArenaTeamId int `json:"arenateamid,omitempty"`
	Guid int `json:"guid,omitempty"`
	WeekGames int16 `json:"weekgames,omitempty"`
	WeekWins int16 `json:"weekwins,omitempty"`
	SeasonGames int16 `json:"seasongames,omitempty"`
	SeasonWins int16 `json:"seasonwins,omitempty"`
	PersonalRating int16 `json:"personalrating,omitempty"`
}

type ArenaTeamMemberSlice []ArenaTeamMember

type DBArenaTeamMember struct {
	bun.BaseModel `bun:"table:arena_team_member,alias:arena_team_member"`
	ArenaTeamId int `bun:"arenaTeamId"`
	Guid int `bun:"guid"`
	WeekGames int16 `bun:"weekGames"`
	WeekWins int16 `bun:"weekWins"`
	SeasonGames int16 `bun:"seasonGames"`
	SeasonWins int16 `bun:"seasonWins"`
	PersonalRating int16 `bun:"personalRating"`
}

type DBArenaTeamMemberSlice []DBArenaTeamMember

func (entry *ArenaTeamMember) ToDB() *DBArenaTeamMember {
	if entry == nil {
		return nil
	}
	return &DBArenaTeamMember{
		ArenaTeamId: entry.ArenaTeamId,
		Guid: entry.Guid,
		WeekGames: entry.WeekGames,
		WeekWins: entry.WeekWins,
		SeasonGames: entry.SeasonGames,
		SeasonWins: entry.SeasonWins,
		PersonalRating: entry.PersonalRating,
	}
}

func (entry *DBArenaTeamMember) ToWeb() *ArenaTeamMember {
	if entry == nil {
		return nil
	}
	return &ArenaTeamMember{
		ArenaTeamId: entry.ArenaTeamId,
		Guid: entry.Guid,
		WeekGames: entry.WeekGames,
		WeekWins: entry.WeekWins,
		SeasonGames: entry.SeasonGames,
		SeasonWins: entry.SeasonWins,
		PersonalRating: entry.PersonalRating,
	}
}

func (data ArenaTeamMemberSlice) ToDB() DBArenaTeamMemberSlice {
	result := make(DBArenaTeamMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBArenaTeamMemberSlice) ToWeb() ArenaTeamMemberSlice {
	result := make(ArenaTeamMemberSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
