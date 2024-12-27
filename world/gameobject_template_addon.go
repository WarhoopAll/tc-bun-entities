package model

import "github.com/uptrace/bun"

type GameobjectTemplateAddon struct {
	Entry int `json:"entry,omitempty"`
	Faction int16 `json:"faction,omitempty"`
	Flags int `json:"flags,omitempty"`
	Mingold int `json:"mingold,omitempty"`
	Maxgold int `json:"maxgold,omitempty"`
	Artkit0 int `json:"artkit0,omitempty"`
	Artkit1 int `json:"artkit1,omitempty"`
	Artkit2 int `json:"artkit2,omitempty"`
	Artkit3 int `json:"artkit3,omitempty"`
}

type GameobjectTemplateAddonSlice []GameobjectTemplateAddon

type DBGameobjectTemplateAddon struct {
	bun.BaseModel `bun:"table:gameobject_template_addon,alias:gameobject_template_addon"`
	Entry int `bun:"entry"`
	Faction int16 `bun:"faction"`
	Flags int `bun:"flags"`
	Mingold int `bun:"mingold"`
	Maxgold int `bun:"maxgold"`
	Artkit0 int `bun:"artkit0"`
	Artkit1 int `bun:"artkit1"`
	Artkit2 int `bun:"artkit2"`
	Artkit3 int `bun:"artkit3"`
}

type DBGameobjectTemplateAddonSlice []DBGameobjectTemplateAddon

func (entry *GameobjectTemplateAddon) ToDB() *DBGameobjectTemplateAddon {
	if entry == nil {
		return nil
	}
	return &DBGameobjectTemplateAddon{
		Entry: entry.Entry,
		Faction: entry.Faction,
		Flags: entry.Flags,
		Mingold: entry.Mingold,
		Maxgold: entry.Maxgold,
		Artkit0: entry.Artkit0,
		Artkit1: entry.Artkit1,
		Artkit2: entry.Artkit2,
		Artkit3: entry.Artkit3,
	}
}

func (entry *DBGameobjectTemplateAddon) ToWeb() *GameobjectTemplateAddon {
	if entry == nil {
		return nil
	}
	return &GameobjectTemplateAddon{
		Entry: entry.Entry,
		Faction: entry.Faction,
		Flags: entry.Flags,
		Mingold: entry.Mingold,
		Maxgold: entry.Maxgold,
		Artkit0: entry.Artkit0,
		Artkit1: entry.Artkit1,
		Artkit2: entry.Artkit2,
		Artkit3: entry.Artkit3,
	}
}

func (data GameobjectTemplateAddonSlice) ToDB() DBGameobjectTemplateAddonSlice {
	result := make(DBGameobjectTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBGameobjectTemplateAddonSlice) ToWeb() GameobjectTemplateAddonSlice {
	result := make(GameobjectTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
