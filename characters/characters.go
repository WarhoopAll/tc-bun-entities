package model

import "github.com/uptrace/bun"

type Characters struct {
	Guid int `json:"guid,omitempty"`
	Account int `json:"account,omitempty"`
	Name string `json:"name,omitempty"`
	Race int8 `json:"race,omitempty"`
	Class int8 `json:"class,omitempty"`
	Gender int8 `json:"gender,omitempty"`
	Level int8 `json:"level,omitempty"`
	Xp int `json:"xp,omitempty"`
	Money int `json:"money,omitempty"`
	Skin int8 `json:"skin,omitempty"`
	Face int8 `json:"face,omitempty"`
	HairStyle int8 `json:"hairstyle,omitempty"`
	HairColor int8 `json:"haircolor,omitempty"`
	FacialStyle int8 `json:"facialstyle,omitempty"`
	BankSlots int8 `json:"bankslots,omitempty"`
	RestState int8 `json:"reststate,omitempty"`
	PlayerFlags int `json:"playerflags,omitempty"`
	PositionX float64 `json:"position_x,omitempty"`
	PositionY float64 `json:"position_y,omitempty"`
	PositionZ float64 `json:"position_z,omitempty"`
	Map int16 `json:"map,omitempty"`
	InstanceId int `json:"instance_id,omitempty"`
	InstanceModeMask int8 `json:"instance_mode_mask,omitempty"`
	Orientation float64 `json:"orientation,omitempty"`
	Taximask string `json:"taximask,omitempty"`
	Online int8 `json:"online,omitempty"`
	Cinematic int8 `json:"cinematic,omitempty"`
	Totaltime int `json:"totaltime,omitempty"`
	Leveltime int `json:"leveltime,omitempty"`
	LogoutTime int `json:"logout_time,omitempty"`
	IsLogoutResting int8 `json:"is_logout_resting,omitempty"`
	RestBonus float64 `json:"rest_bonus,omitempty"`
	ResettalentsCost int `json:"resettalents_cost,omitempty"`
	ResettalentsTime int `json:"resettalents_time,omitempty"`
	TransX float64 `json:"trans_x,omitempty"`
	TransY float64 `json:"trans_y,omitempty"`
	TransZ float64 `json:"trans_z,omitempty"`
	TransO float64 `json:"trans_o,omitempty"`
	Transguid int32 `json:"transguid,omitempty"`
	ExtraFlags int16 `json:"extra_flags,omitempty"`
	StableSlots int8 `json:"stable_slots,omitempty"`
	AtLogin int16 `json:"at_login,omitempty"`
	Zone int16 `json:"zone,omitempty"`
	DeathExpireTime int `json:"death_expire_time,omitempty"`
	TaxiPath string `json:"taxi_path,omitempty"`
	ArenaPoints int `json:"arenapoints,omitempty"`
	TotalHonorPoints int `json:"totalhonorpoints,omitempty"`
	TodayHonorPoints int `json:"todayhonorpoints,omitempty"`
	YesterdayHonorPoints int `json:"yesterdayhonorpoints,omitempty"`
	TotalKills int `json:"totalkills,omitempty"`
	TodayKills int16 `json:"todaykills,omitempty"`
	YesterdayKills int16 `json:"yesterdaykills,omitempty"`
	ChosenTitle int `json:"chosentitle,omitempty"`
	KnownCurrencies int `json:"knowncurrencies,omitempty"`
	WatchedFaction int `json:"watchedfaction,omitempty"`
	Drunk int8 `json:"drunk,omitempty"`
	Health int `json:"health,omitempty"`
	Power1 int `json:"power1,omitempty"`
	Power2 int `json:"power2,omitempty"`
	Power3 int `json:"power3,omitempty"`
	Power4 int `json:"power4,omitempty"`
	Power5 int `json:"power5,omitempty"`
	Power6 int `json:"power6,omitempty"`
	Power7 int `json:"power7,omitempty"`
	Latency int32 `json:"latency,omitempty"`
	TalentGroupsCount int8 `json:"talentgroupscount,omitempty"`
	ActiveTalentGroup int8 `json:"activetalentgroup,omitempty"`
	ExploredZones string `json:"exploredzones,omitempty"`
	EquipmentCache string `json:"equipmentcache,omitempty"`
	AmmoId int `json:"ammoid,omitempty"`
	KnownTitles string `json:"knowntitles,omitempty"`
	ActionBars int8 `json:"actionbars,omitempty"`
	GrantableLevels int8 `json:"grantablelevels,omitempty"`
	DeleteInfosAccount int `json:"deleteinfos_account,omitempty"`
	DeleteInfosName string `json:"deleteinfos_name,omitempty"`
	DeleteDate int `json:"deletedate,omitempty"`
}

type CharactersSlice []Characters

type DBCharacters struct {
	bun.BaseModel `bun:"table:characters,alias:characters"`
	Guid int `bun:"guid"`
	Account int `bun:"account"`
	Name string `bun:"name"`
	Race int8 `bun:"race"`
	Class int8 `bun:"class"`
	Gender int8 `bun:"gender"`
	Level int8 `bun:"level"`
	Xp int `bun:"xp"`
	Money int `bun:"money"`
	Skin int8 `bun:"skin"`
	Face int8 `bun:"face"`
	HairStyle int8 `bun:"hairStyle"`
	HairColor int8 `bun:"hairColor"`
	FacialStyle int8 `bun:"facialStyle"`
	BankSlots int8 `bun:"bankSlots"`
	RestState int8 `bun:"restState"`
	PlayerFlags int `bun:"playerFlags"`
	PositionX float64 `bun:"position_x"`
	PositionY float64 `bun:"position_y"`
	PositionZ float64 `bun:"position_z"`
	Map int16 `bun:"map"`
	InstanceId int `bun:"instance_id"`
	InstanceModeMask int8 `bun:"instance_mode_mask"`
	Orientation float64 `bun:"orientation"`
	Taximask string `bun:"taximask"`
	Online int8 `bun:"online"`
	Cinematic int8 `bun:"cinematic"`
	Totaltime int `bun:"totaltime"`
	Leveltime int `bun:"leveltime"`
	LogoutTime int `bun:"logout_time"`
	IsLogoutResting int8 `bun:"is_logout_resting"`
	RestBonus float64 `bun:"rest_bonus"`
	ResettalentsCost int `bun:"resettalents_cost"`
	ResettalentsTime int `bun:"resettalents_time"`
	TransX float64 `bun:"trans_x"`
	TransY float64 `bun:"trans_y"`
	TransZ float64 `bun:"trans_z"`
	TransO float64 `bun:"trans_o"`
	Transguid int32 `bun:"transguid"`
	ExtraFlags int16 `bun:"extra_flags"`
	StableSlots int8 `bun:"stable_slots"`
	AtLogin int16 `bun:"at_login"`
	Zone int16 `bun:"zone"`
	DeathExpireTime int `bun:"death_expire_time"`
	TaxiPath string `bun:"taxi_path"`
	ArenaPoints int `bun:"arenaPoints"`
	TotalHonorPoints int `bun:"totalHonorPoints"`
	TodayHonorPoints int `bun:"todayHonorPoints"`
	YesterdayHonorPoints int `bun:"yesterdayHonorPoints"`
	TotalKills int `bun:"totalKills"`
	TodayKills int16 `bun:"todayKills"`
	YesterdayKills int16 `bun:"yesterdayKills"`
	ChosenTitle int `bun:"chosenTitle"`
	KnownCurrencies int `bun:"knownCurrencies"`
	WatchedFaction int `bun:"watchedFaction"`
	Drunk int8 `bun:"drunk"`
	Health int `bun:"health"`
	Power1 int `bun:"power1"`
	Power2 int `bun:"power2"`
	Power3 int `bun:"power3"`
	Power4 int `bun:"power4"`
	Power5 int `bun:"power5"`
	Power6 int `bun:"power6"`
	Power7 int `bun:"power7"`
	Latency int32 `bun:"latency"`
	TalentGroupsCount int8 `bun:"talentGroupsCount"`
	ActiveTalentGroup int8 `bun:"activeTalentGroup"`
	ExploredZones string `bun:"exploredZones"`
	EquipmentCache string `bun:"equipmentCache"`
	AmmoId int `bun:"ammoId"`
	KnownTitles string `bun:"knownTitles"`
	ActionBars int8 `bun:"actionBars"`
	GrantableLevels int8 `bun:"grantableLevels"`
	DeleteInfosAccount int `bun:"deleteInfos_Account"`
	DeleteInfosName string `bun:"deleteInfos_Name"`
	DeleteDate int `bun:"deleteDate"`
}

type DBCharactersSlice []DBCharacters

func (entry *Characters) ToDB() *DBCharacters {
	if entry == nil {
		return nil
	}
	return &DBCharacters{
		Guid: entry.Guid,
		Account: entry.Account,
		Name: entry.Name,
		Race: entry.Race,
		Class: entry.Class,
		Gender: entry.Gender,
		Level: entry.Level,
		Xp: entry.Xp,
		Money: entry.Money,
		Skin: entry.Skin,
		Face: entry.Face,
		HairStyle: entry.HairStyle,
		HairColor: entry.HairColor,
		FacialStyle: entry.FacialStyle,
		BankSlots: entry.BankSlots,
		RestState: entry.RestState,
		PlayerFlags: entry.PlayerFlags,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Map: entry.Map,
		InstanceId: entry.InstanceId,
		InstanceModeMask: entry.InstanceModeMask,
		Orientation: entry.Orientation,
		Taximask: entry.Taximask,
		Online: entry.Online,
		Cinematic: entry.Cinematic,
		Totaltime: entry.Totaltime,
		Leveltime: entry.Leveltime,
		LogoutTime: entry.LogoutTime,
		IsLogoutResting: entry.IsLogoutResting,
		RestBonus: entry.RestBonus,
		ResettalentsCost: entry.ResettalentsCost,
		ResettalentsTime: entry.ResettalentsTime,
		TransX: entry.TransX,
		TransY: entry.TransY,
		TransZ: entry.TransZ,
		TransO: entry.TransO,
		Transguid: entry.Transguid,
		ExtraFlags: entry.ExtraFlags,
		StableSlots: entry.StableSlots,
		AtLogin: entry.AtLogin,
		Zone: entry.Zone,
		DeathExpireTime: entry.DeathExpireTime,
		TaxiPath: entry.TaxiPath,
		ArenaPoints: entry.ArenaPoints,
		TotalHonorPoints: entry.TotalHonorPoints,
		TodayHonorPoints: entry.TodayHonorPoints,
		YesterdayHonorPoints: entry.YesterdayHonorPoints,
		TotalKills: entry.TotalKills,
		TodayKills: entry.TodayKills,
		YesterdayKills: entry.YesterdayKills,
		ChosenTitle: entry.ChosenTitle,
		KnownCurrencies: entry.KnownCurrencies,
		WatchedFaction: entry.WatchedFaction,
		Drunk: entry.Drunk,
		Health: entry.Health,
		Power1: entry.Power1,
		Power2: entry.Power2,
		Power3: entry.Power3,
		Power4: entry.Power4,
		Power5: entry.Power5,
		Power6: entry.Power6,
		Power7: entry.Power7,
		Latency: entry.Latency,
		TalentGroupsCount: entry.TalentGroupsCount,
		ActiveTalentGroup: entry.ActiveTalentGroup,
		ExploredZones: entry.ExploredZones,
		EquipmentCache: entry.EquipmentCache,
		AmmoId: entry.AmmoId,
		KnownTitles: entry.KnownTitles,
		ActionBars: entry.ActionBars,
		GrantableLevels: entry.GrantableLevels,
		DeleteInfosAccount: entry.DeleteInfosAccount,
		DeleteInfosName: entry.DeleteInfosName,
		DeleteDate: entry.DeleteDate,
	}
}

func (entry *DBCharacters) ToWeb() *Characters {
	if entry == nil {
		return nil
	}
	return &Characters{
		Guid: entry.Guid,
		Account: entry.Account,
		Name: entry.Name,
		Race: entry.Race,
		Class: entry.Class,
		Gender: entry.Gender,
		Level: entry.Level,
		Xp: entry.Xp,
		Money: entry.Money,
		Skin: entry.Skin,
		Face: entry.Face,
		HairStyle: entry.HairStyle,
		HairColor: entry.HairColor,
		FacialStyle: entry.FacialStyle,
		BankSlots: entry.BankSlots,
		RestState: entry.RestState,
		PlayerFlags: entry.PlayerFlags,
		PositionX: entry.PositionX,
		PositionY: entry.PositionY,
		PositionZ: entry.PositionZ,
		Map: entry.Map,
		InstanceId: entry.InstanceId,
		InstanceModeMask: entry.InstanceModeMask,
		Orientation: entry.Orientation,
		Taximask: entry.Taximask,
		Online: entry.Online,
		Cinematic: entry.Cinematic,
		Totaltime: entry.Totaltime,
		Leveltime: entry.Leveltime,
		LogoutTime: entry.LogoutTime,
		IsLogoutResting: entry.IsLogoutResting,
		RestBonus: entry.RestBonus,
		ResettalentsCost: entry.ResettalentsCost,
		ResettalentsTime: entry.ResettalentsTime,
		TransX: entry.TransX,
		TransY: entry.TransY,
		TransZ: entry.TransZ,
		TransO: entry.TransO,
		Transguid: entry.Transguid,
		ExtraFlags: entry.ExtraFlags,
		StableSlots: entry.StableSlots,
		AtLogin: entry.AtLogin,
		Zone: entry.Zone,
		DeathExpireTime: entry.DeathExpireTime,
		TaxiPath: entry.TaxiPath,
		ArenaPoints: entry.ArenaPoints,
		TotalHonorPoints: entry.TotalHonorPoints,
		TodayHonorPoints: entry.TodayHonorPoints,
		YesterdayHonorPoints: entry.YesterdayHonorPoints,
		TotalKills: entry.TotalKills,
		TodayKills: entry.TodayKills,
		YesterdayKills: entry.YesterdayKills,
		ChosenTitle: entry.ChosenTitle,
		KnownCurrencies: entry.KnownCurrencies,
		WatchedFaction: entry.WatchedFaction,
		Drunk: entry.Drunk,
		Health: entry.Health,
		Power1: entry.Power1,
		Power2: entry.Power2,
		Power3: entry.Power3,
		Power4: entry.Power4,
		Power5: entry.Power5,
		Power6: entry.Power6,
		Power7: entry.Power7,
		Latency: entry.Latency,
		TalentGroupsCount: entry.TalentGroupsCount,
		ActiveTalentGroup: entry.ActiveTalentGroup,
		ExploredZones: entry.ExploredZones,
		EquipmentCache: entry.EquipmentCache,
		AmmoId: entry.AmmoId,
		KnownTitles: entry.KnownTitles,
		ActionBars: entry.ActionBars,
		GrantableLevels: entry.GrantableLevels,
		DeleteInfosAccount: entry.DeleteInfosAccount,
		DeleteInfosName: entry.DeleteInfosName,
		DeleteDate: entry.DeleteDate,
	}
}

func (data CharactersSlice) ToDB() DBCharactersSlice {
	result := make(DBCharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharactersSlice) ToWeb() CharactersSlice {
	result := make(CharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
