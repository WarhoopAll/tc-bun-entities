package model

import "github.com/uptrace/bun"

type CreatureText struct {
	CreatureID int `json:"creatureid,omitempty"`
	GroupID int8 `json:"groupid,omitempty"`
	ID int8 `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	Type int8 `json:"type,omitempty"`
	Language int8 `json:"language,omitempty"`
	Probability float64 `json:"probability,omitempty"`
	Emote int `json:"emote,omitempty"`
	Duration int `json:"duration,omitempty"`
	Sound int `json:"sound,omitempty"`
	BroadcastTextId int `json:"broadcasttextid,omitempty"`
	TextRange int8 `json:"textrange,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type CreatureTextSlice []CreatureText

type DBCreatureText struct {
	bun.BaseModel `bun:"table:creature_text,alias:creature_text"`
	CreatureID int `bun:"CreatureID"`
	GroupID int8 `bun:"GroupID"`
	ID int8 `bun:"ID"`
	Text string `bun:"Text"`
	Type int8 `bun:"Type"`
	Language int8 `bun:"Language"`
	Probability float64 `bun:"Probability"`
	Emote int `bun:"Emote"`
	Duration int `bun:"Duration"`
	Sound int `bun:"Sound"`
	BroadcastTextId int `bun:"BroadcastTextId"`
	TextRange int8 `bun:"TextRange"`
	Comment string `bun:"comment"`
}

type DBCreatureTextSlice []DBCreatureText

func (entry *CreatureText) ToDB() *DBCreatureText {
	if entry == nil {
		return nil
	}
	return &DBCreatureText{
		CreatureID: entry.CreatureID,
		GroupID: entry.GroupID,
		ID: entry.ID,
		Text: entry.Text,
		Type: entry.Type,
		Language: entry.Language,
		Probability: entry.Probability,
		Emote: entry.Emote,
		Duration: entry.Duration,
		Sound: entry.Sound,
		BroadcastTextId: entry.BroadcastTextId,
		TextRange: entry.TextRange,
		Comment: entry.Comment,
	}
}

func (entry *DBCreatureText) ToWeb() *CreatureText {
	if entry == nil {
		return nil
	}
	return &CreatureText{
		CreatureID: entry.CreatureID,
		GroupID: entry.GroupID,
		ID: entry.ID,
		Text: entry.Text,
		Type: entry.Type,
		Language: entry.Language,
		Probability: entry.Probability,
		Emote: entry.Emote,
		Duration: entry.Duration,
		Sound: entry.Sound,
		BroadcastTextId: entry.BroadcastTextId,
		TextRange: entry.TextRange,
		Comment: entry.Comment,
	}
}

func (data CreatureTextSlice) ToDB() DBCreatureTextSlice {
	result := make(DBCreatureTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTextSlice) ToWeb() CreatureTextSlice {
	result := make(CreatureTextSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
