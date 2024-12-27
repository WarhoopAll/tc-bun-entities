package model

import "github.com/uptrace/bun"

type QuestGreeting struct {
	ID int `json:"id,omitempty"`
	Type int8 `json:"type,omitempty"`
	GreetEmoteType int16 `json:"greetemotetype,omitempty"`
	GreetEmoteDelay int `json:"greetemotedelay,omitempty"`
	Greeting string `json:"greeting,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestGreetingSlice []QuestGreeting

type DBQuestGreeting struct {
	bun.BaseModel `bun:"table:quest_greeting,alias:quest_greeting"`
	ID int `bun:"ID"`
	Type int8 `bun:"Type"`
	GreetEmoteType int16 `bun:"GreetEmoteType"`
	GreetEmoteDelay int `bun:"GreetEmoteDelay"`
	Greeting string `bun:"Greeting"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestGreetingSlice []DBQuestGreeting

func (entry *QuestGreeting) ToDB() *DBQuestGreeting {
	if entry == nil {
		return nil
	}
	return &DBQuestGreeting{
		ID: entry.ID,
		Type: entry.Type,
		GreetEmoteType: entry.GreetEmoteType,
		GreetEmoteDelay: entry.GreetEmoteDelay,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestGreeting) ToWeb() *QuestGreeting {
	if entry == nil {
		return nil
	}
	return &QuestGreeting{
		ID: entry.ID,
		Type: entry.Type,
		GreetEmoteType: entry.GreetEmoteType,
		GreetEmoteDelay: entry.GreetEmoteDelay,
		Greeting: entry.Greeting,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestGreetingSlice) ToDB() DBQuestGreetingSlice {
	result := make(DBQuestGreetingSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestGreetingSlice) ToWeb() QuestGreetingSlice {
	result := make(QuestGreetingSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
