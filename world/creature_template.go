package model

import "github.com/uptrace/bun"

type CreatureTemplate struct {
	Entry int `json:"entry,omitempty"`
	DifficultyEntry1 int `json:"difficulty_entry_1,omitempty"`
	DifficultyEntry2 int `json:"difficulty_entry_2,omitempty"`
	DifficultyEntry3 int `json:"difficulty_entry_3,omitempty"`
	KillCredit1 int `json:"killcredit1,omitempty"`
	KillCredit2 int `json:"killcredit2,omitempty"`
	Modelid1 int `json:"modelid1,omitempty"`
	Modelid2 int `json:"modelid2,omitempty"`
	Modelid3 int `json:"modelid3,omitempty"`
	Modelid4 int `json:"modelid4,omitempty"`
	Name string `json:"name,omitempty"`
	Subname string `json:"subname,omitempty"`
	IconName string `json:"iconname,omitempty"`
	GossipMenuId int `json:"gossip_menu_id,omitempty"`
	Minlevel int8 `json:"minlevel,omitempty"`
	Maxlevel int8 `json:"maxlevel,omitempty"`
	Exp int16 `json:"exp,omitempty"`
	Faction int16 `json:"faction,omitempty"`
	Npcflag int `json:"npcflag,omitempty"`
	SpeedWalk float64 `json:"speed_walk,omitempty"`
	SpeedRun float64 `json:"speed_run,omitempty"`
	Scale float64 `json:"scale,omitempty"`
	Rank int8 `json:"rank,omitempty"`
	Dmgschool int8 `json:"dmgschool,omitempty"`
	BaseAttackTime int `json:"baseattacktime,omitempty"`
	RangeAttackTime int `json:"rangeattacktime,omitempty"`
	BaseVariance float64 `json:"basevariance,omitempty"`
	RangeVariance float64 `json:"rangevariance,omitempty"`
	UnitClass int8 `json:"unit_class,omitempty"`
	UnitFlags int `json:"unit_flags,omitempty"`
	UnitFlags2 int `json:"unit_flags2,omitempty"`
	Dynamicflags int `json:"dynamicflags,omitempty"`
	Family int8 `json:"family,omitempty"`
	Type int8 `json:"type,omitempty"`
	TypeFlags int `json:"type_flags,omitempty"`
	Lootid int `json:"lootid,omitempty"`
	Pickpocketloot int `json:"pickpocketloot,omitempty"`
	Skinloot int `json:"skinloot,omitempty"`
	PetSpellDataId int `json:"petspelldataid,omitempty"`
	VehicleId int `json:"vehicleid,omitempty"`
	Mingold int `json:"mingold,omitempty"`
	Maxgold int `json:"maxgold,omitempty"`
	AIName string `json:"ainame,omitempty"`
	MovementType int8 `json:"movementtype,omitempty"`
	HoverHeight float64 `json:"hoverheight,omitempty"`
	HealthModifier float64 `json:"healthmodifier,omitempty"`
	ManaModifier float64 `json:"manamodifier,omitempty"`
	ArmorModifier float64 `json:"armormodifier,omitempty"`
	DamageModifier float64 `json:"damagemodifier,omitempty"`
	ExperienceModifier float64 `json:"experiencemodifier,omitempty"`
	RacialLeader int8 `json:"racialleader,omitempty"`
	MovementId int `json:"movementid,omitempty"`
	RegenHealth int8 `json:"regenhealth,omitempty"`
	MechanicImmuneMask int `json:"mechanic_immune_mask,omitempty"`
	SpellSchoolImmuneMask int `json:"spell_school_immune_mask,omitempty"`
	FlagsExtra int `json:"flags_extra,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	StringId string `json:"stringid,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type CreatureTemplateSlice []CreatureTemplate

type DBCreatureTemplate struct {
	bun.BaseModel `bun:"table:creature_template,alias:creature_template"`
	Entry int `bun:"entry"`
	DifficultyEntry1 int `bun:"difficulty_entry_1"`
	DifficultyEntry2 int `bun:"difficulty_entry_2"`
	DifficultyEntry3 int `bun:"difficulty_entry_3"`
	KillCredit1 int `bun:"KillCredit1"`
	KillCredit2 int `bun:"KillCredit2"`
	Modelid1 int `bun:"modelid1"`
	Modelid2 int `bun:"modelid2"`
	Modelid3 int `bun:"modelid3"`
	Modelid4 int `bun:"modelid4"`
	Name string `bun:"name"`
	Subname string `bun:"subname"`
	IconName string `bun:"IconName"`
	GossipMenuId int `bun:"gossip_menu_id"`
	Minlevel int8 `bun:"minlevel"`
	Maxlevel int8 `bun:"maxlevel"`
	Exp int16 `bun:"exp"`
	Faction int16 `bun:"faction"`
	Npcflag int `bun:"npcflag"`
	SpeedWalk float64 `bun:"speed_walk"`
	SpeedRun float64 `bun:"speed_run"`
	Scale float64 `bun:"scale"`
	Rank int8 `bun:"rank"`
	Dmgschool int8 `bun:"dmgschool"`
	BaseAttackTime int `bun:"BaseAttackTime"`
	RangeAttackTime int `bun:"RangeAttackTime"`
	BaseVariance float64 `bun:"BaseVariance"`
	RangeVariance float64 `bun:"RangeVariance"`
	UnitClass int8 `bun:"unit_class"`
	UnitFlags int `bun:"unit_flags"`
	UnitFlags2 int `bun:"unit_flags2"`
	Dynamicflags int `bun:"dynamicflags"`
	Family int8 `bun:"family"`
	Type int8 `bun:"type"`
	TypeFlags int `bun:"type_flags"`
	Lootid int `bun:"lootid"`
	Pickpocketloot int `bun:"pickpocketloot"`
	Skinloot int `bun:"skinloot"`
	PetSpellDataId int `bun:"PetSpellDataId"`
	VehicleId int `bun:"VehicleId"`
	Mingold int `bun:"mingold"`
	Maxgold int `bun:"maxgold"`
	AIName string `bun:"AIName"`
	MovementType int8 `bun:"MovementType"`
	HoverHeight float64 `bun:"HoverHeight"`
	HealthModifier float64 `bun:"HealthModifier"`
	ManaModifier float64 `bun:"ManaModifier"`
	ArmorModifier float64 `bun:"ArmorModifier"`
	DamageModifier float64 `bun:"DamageModifier"`
	ExperienceModifier float64 `bun:"ExperienceModifier"`
	RacialLeader int8 `bun:"RacialLeader"`
	MovementId int `bun:"movementId"`
	RegenHealth int8 `bun:"RegenHealth"`
	MechanicImmuneMask int `bun:"mechanic_immune_mask"`
	SpellSchoolImmuneMask int `bun:"spell_school_immune_mask"`
	FlagsExtra int `bun:"flags_extra"`
	ScriptName string `bun:"ScriptName"`
	StringId string `bun:"StringId"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBCreatureTemplateSlice []DBCreatureTemplate

func (entry *CreatureTemplate) ToDB() *DBCreatureTemplate {
	if entry == nil {
		return nil
	}
	return &DBCreatureTemplate{
		Entry: entry.Entry,
		DifficultyEntry1: entry.DifficultyEntry1,
		DifficultyEntry2: entry.DifficultyEntry2,
		DifficultyEntry3: entry.DifficultyEntry3,
		KillCredit1: entry.KillCredit1,
		KillCredit2: entry.KillCredit2,
		Modelid1: entry.Modelid1,
		Modelid2: entry.Modelid2,
		Modelid3: entry.Modelid3,
		Modelid4: entry.Modelid4,
		Name: entry.Name,
		Subname: entry.Subname,
		IconName: entry.IconName,
		GossipMenuId: entry.GossipMenuId,
		Minlevel: entry.Minlevel,
		Maxlevel: entry.Maxlevel,
		Exp: entry.Exp,
		Faction: entry.Faction,
		Npcflag: entry.Npcflag,
		SpeedWalk: entry.SpeedWalk,
		SpeedRun: entry.SpeedRun,
		Scale: entry.Scale,
		Rank: entry.Rank,
		Dmgschool: entry.Dmgschool,
		BaseAttackTime: entry.BaseAttackTime,
		RangeAttackTime: entry.RangeAttackTime,
		BaseVariance: entry.BaseVariance,
		RangeVariance: entry.RangeVariance,
		UnitClass: entry.UnitClass,
		UnitFlags: entry.UnitFlags,
		UnitFlags2: entry.UnitFlags2,
		Dynamicflags: entry.Dynamicflags,
		Family: entry.Family,
		Type: entry.Type,
		TypeFlags: entry.TypeFlags,
		Lootid: entry.Lootid,
		Pickpocketloot: entry.Pickpocketloot,
		Skinloot: entry.Skinloot,
		PetSpellDataId: entry.PetSpellDataId,
		VehicleId: entry.VehicleId,
		Mingold: entry.Mingold,
		Maxgold: entry.Maxgold,
		AIName: entry.AIName,
		MovementType: entry.MovementType,
		HoverHeight: entry.HoverHeight,
		HealthModifier: entry.HealthModifier,
		ManaModifier: entry.ManaModifier,
		ArmorModifier: entry.ArmorModifier,
		DamageModifier: entry.DamageModifier,
		ExperienceModifier: entry.ExperienceModifier,
		RacialLeader: entry.RacialLeader,
		MovementId: entry.MovementId,
		RegenHealth: entry.RegenHealth,
		MechanicImmuneMask: entry.MechanicImmuneMask,
		SpellSchoolImmuneMask: entry.SpellSchoolImmuneMask,
		FlagsExtra: entry.FlagsExtra,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBCreatureTemplate) ToWeb() *CreatureTemplate {
	if entry == nil {
		return nil
	}
	return &CreatureTemplate{
		Entry: entry.Entry,
		DifficultyEntry1: entry.DifficultyEntry1,
		DifficultyEntry2: entry.DifficultyEntry2,
		DifficultyEntry3: entry.DifficultyEntry3,
		KillCredit1: entry.KillCredit1,
		KillCredit2: entry.KillCredit2,
		Modelid1: entry.Modelid1,
		Modelid2: entry.Modelid2,
		Modelid3: entry.Modelid3,
		Modelid4: entry.Modelid4,
		Name: entry.Name,
		Subname: entry.Subname,
		IconName: entry.IconName,
		GossipMenuId: entry.GossipMenuId,
		Minlevel: entry.Minlevel,
		Maxlevel: entry.Maxlevel,
		Exp: entry.Exp,
		Faction: entry.Faction,
		Npcflag: entry.Npcflag,
		SpeedWalk: entry.SpeedWalk,
		SpeedRun: entry.SpeedRun,
		Scale: entry.Scale,
		Rank: entry.Rank,
		Dmgschool: entry.Dmgschool,
		BaseAttackTime: entry.BaseAttackTime,
		RangeAttackTime: entry.RangeAttackTime,
		BaseVariance: entry.BaseVariance,
		RangeVariance: entry.RangeVariance,
		UnitClass: entry.UnitClass,
		UnitFlags: entry.UnitFlags,
		UnitFlags2: entry.UnitFlags2,
		Dynamicflags: entry.Dynamicflags,
		Family: entry.Family,
		Type: entry.Type,
		TypeFlags: entry.TypeFlags,
		Lootid: entry.Lootid,
		Pickpocketloot: entry.Pickpocketloot,
		Skinloot: entry.Skinloot,
		PetSpellDataId: entry.PetSpellDataId,
		VehicleId: entry.VehicleId,
		Mingold: entry.Mingold,
		Maxgold: entry.Maxgold,
		AIName: entry.AIName,
		MovementType: entry.MovementType,
		HoverHeight: entry.HoverHeight,
		HealthModifier: entry.HealthModifier,
		ManaModifier: entry.ManaModifier,
		ArmorModifier: entry.ArmorModifier,
		DamageModifier: entry.DamageModifier,
		ExperienceModifier: entry.ExperienceModifier,
		RacialLeader: entry.RacialLeader,
		MovementId: entry.MovementId,
		RegenHealth: entry.RegenHealth,
		MechanicImmuneMask: entry.MechanicImmuneMask,
		SpellSchoolImmuneMask: entry.SpellSchoolImmuneMask,
		FlagsExtra: entry.FlagsExtra,
		ScriptName: entry.ScriptName,
		StringId: entry.StringId,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data CreatureTemplateSlice) ToDB() DBCreatureTemplateSlice {
	result := make(DBCreatureTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCreatureTemplateSlice) ToWeb() CreatureTemplateSlice {
	result := make(CreatureTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
