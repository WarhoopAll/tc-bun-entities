package model

import "github.com/uptrace/bun"

type CreatureAddon struct {
	Guid int `json:"guid,omitempty"`
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

type CreatureAddonSlice []CreatureAddon

type DBCreatureAddon struct {
	bun.BaseModel `bun:"table:creature_addon,alias:creature_addon"`
	Guid int `bun:"guid"`
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

type DBCreatureAddonSlice []DBCreatureAddon

func (entry *CreatureAddon) ToDB() *DBCreatureAddon {
	if entry == nil {
		return nil
	}
	return &DBCreatureAddon{
		Guid: entry.Guid,
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

func (entry *DBCreatureAddon) ToWeb() *CreatureAddon {
	if entry == nil {
		return nil
	}
	return &CreatureAddon{
		Guid: entry.Guid,
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

func (data CreatureAddonSlice) ToDB() DBCreatureAddonSlice {
	result := make(DBCreatureAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureAddonSlice) ToWeb() CreatureAddonSlice {
	result := make(CreatureAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
