package model

import "github.com/uptrace/bun"

type CreatureTemplateAddon struct {
	Entry int `json:"entry,omitempty"`
	PathId int `json:"path_id,omitempty"`
	Mount int `json:"mount,omitempty"`
	MountCreatureID int `json:"mountcreatureid,omitempty"`
	StandState int8 `json:"standstate,omitempty"`
	AnimTier int8 `json:"animtier,omitempty"`
	VisFlags int8 `json:"visflags,omitempty"`
	SheathState int8 `json:"sheathstate,omitempty"`
	PvPFlags int8 `json:"pvpflags,omitempty"`
	Emote int `json:"emote,omitempty"`
	VisibilityDistanceType int8 `json:"visibilitydistancetype,omitempty"`
	Auras string `json:"auras,omitempty"`
}

type CreatureTemplateAddonSlice []CreatureTemplateAddon

type DBCreatureTemplateAddon struct {
	bun.BaseModel `bun:"table:creature_template_addon,alias:creature_template_addon"`
	Entry int `bun:"entry"`
	PathId int `bun:"path_id"`
	Mount int `bun:"mount"`
	MountCreatureID int `bun:"MountCreatureID"`
	StandState int8 `bun:"StandState"`
	AnimTier int8 `bun:"AnimTier"`
	VisFlags int8 `bun:"VisFlags"`
	SheathState int8 `bun:"SheathState"`
	PvPFlags int8 `bun:"PvPFlags"`
	Emote int `bun:"emote"`
	VisibilityDistanceType int8 `bun:"visibilityDistanceType"`
	Auras string `bun:"auras"`
}

type DBCreatureTemplateAddonSlice []DBCreatureTemplateAddon

func (entry *CreatureTemplateAddon) ToDB() *DBCreatureTemplateAddon {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplateAddon{
		Entry: entry.Entry,
		PathId: entry.PathId,
		Mount: entry.Mount,
		MountCreatureID: entry.MountCreatureID,
		StandState: entry.StandState,
		AnimTier: entry.AnimTier,
		VisFlags: entry.VisFlags,
		SheathState: entry.SheathState,
		PvPFlags: entry.PvPFlags,
		Emote: entry.Emote,
		VisibilityDistanceType: entry.VisibilityDistanceType,
		Auras: entry.Auras,
	}
}

func (entry *DBCreatureTemplateAddon) ToWeb() *CreatureTemplateAddon {
	if entry == nil {
		return nil
	}
	return &CreatureTemplateAddon{
		Entry: entry.Entry,
		PathId: entry.PathId,
		Mount: entry.Mount,
		MountCreatureID: entry.MountCreatureID,
		StandState: entry.StandState,
		AnimTier: entry.AnimTier,
		VisFlags: entry.VisFlags,
		SheathState: entry.SheathState,
		PvPFlags: entry.PvPFlags,
		Emote: entry.Emote,
		VisibilityDistanceType: entry.VisibilityDistanceType,
		Auras: entry.Auras,
	}
}

func (data CreatureTemplateAddonSlice) ToDB() DBCreatureTemplateAddonSlice {
	result := make(DBCreatureTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateAddonSlice) ToWeb() CreatureTemplateAddonSlice {
	result := make(CreatureTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
