package model

import "github.com/uptrace/bun"

type CalendarInvites struct {
	Id int `json:"id,omitempty"`
	Event int `json:"event,omitempty"`
	Invitee int `json:"invitee,omitempty"`
	Sender int `json:"sender,omitempty"`
	Status int8 `json:"status,omitempty"`
	Statustime int `json:"statustime,omitempty"`
	Rank int8 `json:"rank,omitempty"`
	Text string `json:"text,omitempty"`
}

type CalendarInvitesSlice []CalendarInvites

type DBCalendarInvites struct {
	bun.BaseModel `bun:"table:calendar_invites,alias:calendar_invites"`
	Id int `bun:"id"`
	Event int `bun:"event"`
	Invitee int `bun:"invitee"`
	Sender int `bun:"sender"`
	Status int8 `bun:"status"`
	Statustime int `bun:"statustime"`
	Rank int8 `bun:"rank"`
	Text string `bun:"text"`
}

type DBCalendarInvitesSlice []DBCalendarInvites

func (entry *CalendarInvites) ToDB() *DBCalendarInvites {
	if entry == nil {
		return nil
	}
	return &DBCalendarInvites{
		Id: entry.Id,
		Event: entry.Event,
		Invitee: entry.Invitee,
		Sender: entry.Sender,
		Status: entry.Status,
		Statustime: entry.Statustime,
		Rank: entry.Rank,
		Text: entry.Text,
	}
}

func (entry *DBCalendarInvites) ToWeb() *CalendarInvites {
	if entry == nil {
		return nil
	}
	return &CalendarInvites{
		Id: entry.Id,
		Event: entry.Event,
		Invitee: entry.Invitee,
		Sender: entry.Sender,
		Status: entry.Status,
		Statustime: entry.Statustime,
		Rank: entry.Rank,
		Text: entry.Text,
	}
}

func (data CalendarInvitesSlice) ToDB() DBCalendarInvitesSlice {
	result := make(DBCalendarInvitesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCalendarInvitesSlice) ToWeb() CalendarInvitesSlice {
	result := make(CalendarInvitesSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
