package model

import "github.com/uptrace/bun"

type CalendarEvents struct {
	Id int `json:"id,omitempty"`
	Creator int `json:"creator,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Type int8 `json:"type,omitempty"`
	Dungeon int `json:"dungeon,omitempty"`
	Eventtime int `json:"eventtime,omitempty"`
	Flags int `json:"flags,omitempty"`
	Time2 int `json:"time2,omitempty"`
}

type CalendarEventsSlice []CalendarEvents

type DBCalendarEvents struct {
	bun.BaseModel `bun:"table:calendar_events,alias:calendar_events"`
	Id int `bun:"id"`
	Creator int `bun:"creator"`
	Title string `bun:"title"`
	Description string `bun:"description"`
	Type int8 `bun:"type"`
	Dungeon int `bun:"dungeon"`
	Eventtime int `bun:"eventtime"`
	Flags int `bun:"flags"`
	Time2 int `bun:"time2"`
}

type DBCalendarEventsSlice []DBCalendarEvents

func (entry *CalendarEvents) ToDB() *DBCalendarEvents {
	if entry == nil {
		return nil
	}
	return &DBCalendarEvents{
		Id: entry.Id,
		Creator: entry.Creator,
		Title: entry.Title,
		Description: entry.Description,
		Type: entry.Type,
		Dungeon: entry.Dungeon,
		Eventtime: entry.Eventtime,
		Flags: entry.Flags,
		Time2: entry.Time2,
	}
}

func (entry *DBCalendarEvents) ToWeb() *CalendarEvents {
	if entry == nil {
		return nil
	}
	return &CalendarEvents{
		Id: entry.Id,
		Creator: entry.Creator,
		Title: entry.Title,
		Description: entry.Description,
		Type: entry.Type,
		Dungeon: entry.Dungeon,
		Eventtime: entry.Eventtime,
		Flags: entry.Flags,
		Time2: entry.Time2,
	}
}

func (data CalendarEventsSlice) ToDB() DBCalendarEventsSlice {
	result := make(DBCalendarEventsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCalendarEventsSlice) ToWeb() CalendarEventsSlice {
	result := make(CalendarEventsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
