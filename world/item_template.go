package model

import "github.com/uptrace/bun"

type ItemTemplate struct {
	Entry int `json:"entry,omitempty"`
	Class int8 `json:"class,omitempty"`
	Subclass int8 `json:"subclass,omitempty"`
	SoundOverrideSubclass int8 `json:"soundoverridesubclass,omitempty"`
	Name string `json:"name,omitempty"`
	Displayid int `json:"displayid,omitempty"`
	Quality int8 `json:"quality,omitempty"`
	Flags int `json:"flags,omitempty"`
	FlagsExtra int `json:"flagsextra,omitempty"`
	BuyCount int8 `json:"buycount,omitempty"`
	BuyPrice int `json:"buyprice,omitempty"`
	SellPrice int `json:"sellprice,omitempty"`
	InventoryType int8 `json:"inventorytype,omitempty"`
	AllowableClass int `json:"allowableclass,omitempty"`
	AllowableRace int `json:"allowablerace,omitempty"`
	ItemLevel int16 `json:"itemlevel,omitempty"`
	RequiredLevel int8 `json:"requiredlevel,omitempty"`
	RequiredSkill int16 `json:"requiredskill,omitempty"`
	RequiredSkillRank int16 `json:"requiredskillrank,omitempty"`
	Requiredspell int `json:"requiredspell,omitempty"`
	Requiredhonorrank int `json:"requiredhonorrank,omitempty"`
	RequiredCityRank int `json:"requiredcityrank,omitempty"`
	RequiredReputationFaction int16 `json:"requiredreputationfaction,omitempty"`
	RequiredReputationRank int16 `json:"requiredreputationrank,omitempty"`
	Maxcount int `json:"maxcount,omitempty"`
	Stackable int `json:"stackable,omitempty"`
	ContainerSlots int8 `json:"containerslots,omitempty"`
	StatsCount int8 `json:"statscount,omitempty"`
	StatType1 int8 `json:"stat_type1,omitempty"`
	StatValue1 int16 `json:"stat_value1,omitempty"`
	StatType2 int8 `json:"stat_type2,omitempty"`
	StatValue2 int16 `json:"stat_value2,omitempty"`
	StatType3 int8 `json:"stat_type3,omitempty"`
	StatValue3 int16 `json:"stat_value3,omitempty"`
	StatType4 int8 `json:"stat_type4,omitempty"`
	StatValue4 int16 `json:"stat_value4,omitempty"`
	StatType5 int8 `json:"stat_type5,omitempty"`
	StatValue5 int16 `json:"stat_value5,omitempty"`
	StatType6 int8 `json:"stat_type6,omitempty"`
	StatValue6 int16 `json:"stat_value6,omitempty"`
	StatType7 int8 `json:"stat_type7,omitempty"`
	StatValue7 int16 `json:"stat_value7,omitempty"`
	StatType8 int8 `json:"stat_type8,omitempty"`
	StatValue8 int16 `json:"stat_value8,omitempty"`
	StatType9 int8 `json:"stat_type9,omitempty"`
	StatValue9 int16 `json:"stat_value9,omitempty"`
	StatType10 int8 `json:"stat_type10,omitempty"`
	StatValue10 int16 `json:"stat_value10,omitempty"`
	ScalingStatDistribution int16 `json:"scalingstatdistribution,omitempty"`
	ScalingStatValue int `json:"scalingstatvalue,omitempty"`
	DmgMin1 float64 `json:"dmg_min1,omitempty"`
	DmgMax1 float64 `json:"dmg_max1,omitempty"`
	DmgType1 int8 `json:"dmg_type1,omitempty"`
	DmgMin2 float64 `json:"dmg_min2,omitempty"`
	DmgMax2 float64 `json:"dmg_max2,omitempty"`
	DmgType2 int8 `json:"dmg_type2,omitempty"`
	Armor int16 `json:"armor,omitempty"`
	HolyRes int8 `json:"holy_res,omitempty"`
	FireRes int8 `json:"fire_res,omitempty"`
	NatureRes int8 `json:"nature_res,omitempty"`
	FrostRes int8 `json:"frost_res,omitempty"`
	ShadowRes int8 `json:"shadow_res,omitempty"`
	ArcaneRes int8 `json:"arcane_res,omitempty"`
	Delay int16 `json:"delay,omitempty"`
	AmmoType int8 `json:"ammo_type,omitempty"`
	RangedModRange float64 `json:"rangedmodrange,omitempty"`
	Spellid1 int `json:"spellid_1,omitempty"`
	Spelltrigger1 int8 `json:"spelltrigger_1,omitempty"`
	Spellcharges1 int16 `json:"spellcharges_1,omitempty"`
	SpellppmRate1 float64 `json:"spellppmrate_1,omitempty"`
	Spellcooldown1 int `json:"spellcooldown_1,omitempty"`
	Spellcategory1 int16 `json:"spellcategory_1,omitempty"`
	Spellcategorycooldown1 int `json:"spellcategorycooldown_1,omitempty"`
	Spellid2 int `json:"spellid_2,omitempty"`
	Spelltrigger2 int8 `json:"spelltrigger_2,omitempty"`
	Spellcharges2 int16 `json:"spellcharges_2,omitempty"`
	SpellppmRate2 float64 `json:"spellppmrate_2,omitempty"`
	Spellcooldown2 int `json:"spellcooldown_2,omitempty"`
	Spellcategory2 int16 `json:"spellcategory_2,omitempty"`
	Spellcategorycooldown2 int `json:"spellcategorycooldown_2,omitempty"`
	Spellid3 int `json:"spellid_3,omitempty"`
	Spelltrigger3 int8 `json:"spelltrigger_3,omitempty"`
	Spellcharges3 int16 `json:"spellcharges_3,omitempty"`
	SpellppmRate3 float64 `json:"spellppmrate_3,omitempty"`
	Spellcooldown3 int `json:"spellcooldown_3,omitempty"`
	Spellcategory3 int16 `json:"spellcategory_3,omitempty"`
	Spellcategorycooldown3 int `json:"spellcategorycooldown_3,omitempty"`
	Spellid4 int `json:"spellid_4,omitempty"`
	Spelltrigger4 int8 `json:"spelltrigger_4,omitempty"`
	Spellcharges4 int16 `json:"spellcharges_4,omitempty"`
	SpellppmRate4 float64 `json:"spellppmrate_4,omitempty"`
	Spellcooldown4 int `json:"spellcooldown_4,omitempty"`
	Spellcategory4 int16 `json:"spellcategory_4,omitempty"`
	Spellcategorycooldown4 int `json:"spellcategorycooldown_4,omitempty"`
	Spellid5 int `json:"spellid_5,omitempty"`
	Spelltrigger5 int8 `json:"spelltrigger_5,omitempty"`
	Spellcharges5 int16 `json:"spellcharges_5,omitempty"`
	SpellppmRate5 float64 `json:"spellppmrate_5,omitempty"`
	Spellcooldown5 int `json:"spellcooldown_5,omitempty"`
	Spellcategory5 int16 `json:"spellcategory_5,omitempty"`
	Spellcategorycooldown5 int `json:"spellcategorycooldown_5,omitempty"`
	Bonding int8 `json:"bonding,omitempty"`
	Description string `json:"description,omitempty"`
	PageText int `json:"pagetext,omitempty"`
	LanguageID int8 `json:"languageid,omitempty"`
	PageMaterial int8 `json:"pagematerial,omitempty"`
	Startquest int `json:"startquest,omitempty"`
	Lockid int `json:"lockid,omitempty"`
	Material int8 `json:"material,omitempty"`
	Sheath int8 `json:"sheath,omitempty"`
	RandomProperty int `json:"randomproperty,omitempty"`
	RandomSuffix int `json:"randomsuffix,omitempty"`
	Block int `json:"block,omitempty"`
	Itemset int `json:"itemset,omitempty"`
	MaxDurability int16 `json:"maxdurability,omitempty"`
	Area int `json:"area,omitempty"`
	Map int16 `json:"map,omitempty"`
	BagFamily int `json:"bagfamily,omitempty"`
	TotemCategory int `json:"totemcategory,omitempty"`
	SocketColor1 int8 `json:"socketcolor_1,omitempty"`
	SocketContent1 int `json:"socketcontent_1,omitempty"`
	SocketColor2 int8 `json:"socketcolor_2,omitempty"`
	SocketContent2 int `json:"socketcontent_2,omitempty"`
	SocketColor3 int8 `json:"socketcolor_3,omitempty"`
	SocketContent3 int `json:"socketcontent_3,omitempty"`
	SocketBonus int `json:"socketbonus,omitempty"`
	GemProperties int `json:"gemproperties,omitempty"`
	RequiredDisenchantSkill int16 `json:"requireddisenchantskill,omitempty"`
	ArmorDamageModifier float64 `json:"armordamagemodifier,omitempty"`
	Duration int `json:"duration,omitempty"`
	ItemLimitCategory int16 `json:"itemlimitcategory,omitempty"`
	HolidayId int `json:"holidayid,omitempty"`
	ScriptName string `json:"scriptname,omitempty"`
	DisenchantID int `json:"disenchantid,omitempty"`
	FoodType int8 `json:"foodtype,omitempty"`
	MinMoneyLoot int `json:"minmoneyloot,omitempty"`
	MaxMoneyLoot int `json:"maxmoneyloot,omitempty"`
	FlagsCustom int `json:"flagscustom,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type ItemTemplateSlice []ItemTemplate

type DBItemTemplate struct {
	bun.BaseModel `bun:"table:item_template,alias:item_template"`
	Entry int `bun:"entry"`
	Class int8 `bun:"class"`
	Subclass int8 `bun:"subclass"`
	SoundOverrideSubclass int8 `bun:"SoundOverrideSubclass"`
	Name string `bun:"name"`
	Displayid int `bun:"displayid"`
	Quality int8 `bun:"Quality"`
	Flags int `bun:"Flags"`
	FlagsExtra int `bun:"FlagsExtra"`
	BuyCount int8 `bun:"BuyCount"`
	BuyPrice int `bun:"BuyPrice"`
	SellPrice int `bun:"SellPrice"`
	InventoryType int8 `bun:"InventoryType"`
	AllowableClass int `bun:"AllowableClass"`
	AllowableRace int `bun:"AllowableRace"`
	ItemLevel int16 `bun:"ItemLevel"`
	RequiredLevel int8 `bun:"RequiredLevel"`
	RequiredSkill int16 `bun:"RequiredSkill"`
	RequiredSkillRank int16 `bun:"RequiredSkillRank"`
	Requiredspell int `bun:"requiredspell"`
	Requiredhonorrank int `bun:"requiredhonorrank"`
	RequiredCityRank int `bun:"RequiredCityRank"`
	RequiredReputationFaction int16 `bun:"RequiredReputationFaction"`
	RequiredReputationRank int16 `bun:"RequiredReputationRank"`
	Maxcount int `bun:"maxcount"`
	Stackable int `bun:"stackable"`
	ContainerSlots int8 `bun:"ContainerSlots"`
	StatsCount int8 `bun:"StatsCount"`
	StatType1 int8 `bun:"stat_type1"`
	StatValue1 int16 `bun:"stat_value1"`
	StatType2 int8 `bun:"stat_type2"`
	StatValue2 int16 `bun:"stat_value2"`
	StatType3 int8 `bun:"stat_type3"`
	StatValue3 int16 `bun:"stat_value3"`
	StatType4 int8 `bun:"stat_type4"`
	StatValue4 int16 `bun:"stat_value4"`
	StatType5 int8 `bun:"stat_type5"`
	StatValue5 int16 `bun:"stat_value5"`
	StatType6 int8 `bun:"stat_type6"`
	StatValue6 int16 `bun:"stat_value6"`
	StatType7 int8 `bun:"stat_type7"`
	StatValue7 int16 `bun:"stat_value7"`
	StatType8 int8 `bun:"stat_type8"`
	StatValue8 int16 `bun:"stat_value8"`
	StatType9 int8 `bun:"stat_type9"`
	StatValue9 int16 `bun:"stat_value9"`
	StatType10 int8 `bun:"stat_type10"`
	StatValue10 int16 `bun:"stat_value10"`
	ScalingStatDistribution int16 `bun:"ScalingStatDistribution"`
	ScalingStatValue int `bun:"ScalingStatValue"`
	DmgMin1 float64 `bun:"dmg_min1"`
	DmgMax1 float64 `bun:"dmg_max1"`
	DmgType1 int8 `bun:"dmg_type1"`
	DmgMin2 float64 `bun:"dmg_min2"`
	DmgMax2 float64 `bun:"dmg_max2"`
	DmgType2 int8 `bun:"dmg_type2"`
	Armor int16 `bun:"armor"`
	HolyRes int8 `bun:"holy_res"`
	FireRes int8 `bun:"fire_res"`
	NatureRes int8 `bun:"nature_res"`
	FrostRes int8 `bun:"frost_res"`
	ShadowRes int8 `bun:"shadow_res"`
	ArcaneRes int8 `bun:"arcane_res"`
	Delay int16 `bun:"delay"`
	AmmoType int8 `bun:"ammo_type"`
	RangedModRange float64 `bun:"RangedModRange"`
	Spellid1 int `bun:"spellid_1"`
	Spelltrigger1 int8 `bun:"spelltrigger_1"`
	Spellcharges1 int16 `bun:"spellcharges_1"`
	SpellppmRate1 float64 `bun:"spellppmRate_1"`
	Spellcooldown1 int `bun:"spellcooldown_1"`
	Spellcategory1 int16 `bun:"spellcategory_1"`
	Spellcategorycooldown1 int `bun:"spellcategorycooldown_1"`
	Spellid2 int `bun:"spellid_2"`
	Spelltrigger2 int8 `bun:"spelltrigger_2"`
	Spellcharges2 int16 `bun:"spellcharges_2"`
	SpellppmRate2 float64 `bun:"spellppmRate_2"`
	Spellcooldown2 int `bun:"spellcooldown_2"`
	Spellcategory2 int16 `bun:"spellcategory_2"`
	Spellcategorycooldown2 int `bun:"spellcategorycooldown_2"`
	Spellid3 int `bun:"spellid_3"`
	Spelltrigger3 int8 `bun:"spelltrigger_3"`
	Spellcharges3 int16 `bun:"spellcharges_3"`
	SpellppmRate3 float64 `bun:"spellppmRate_3"`
	Spellcooldown3 int `bun:"spellcooldown_3"`
	Spellcategory3 int16 `bun:"spellcategory_3"`
	Spellcategorycooldown3 int `bun:"spellcategorycooldown_3"`
	Spellid4 int `bun:"spellid_4"`
	Spelltrigger4 int8 `bun:"spelltrigger_4"`
	Spellcharges4 int16 `bun:"spellcharges_4"`
	SpellppmRate4 float64 `bun:"spellppmRate_4"`
	Spellcooldown4 int `bun:"spellcooldown_4"`
	Spellcategory4 int16 `bun:"spellcategory_4"`
	Spellcategorycooldown4 int `bun:"spellcategorycooldown_4"`
	Spellid5 int `bun:"spellid_5"`
	Spelltrigger5 int8 `bun:"spelltrigger_5"`
	Spellcharges5 int16 `bun:"spellcharges_5"`
	SpellppmRate5 float64 `bun:"spellppmRate_5"`
	Spellcooldown5 int `bun:"spellcooldown_5"`
	Spellcategory5 int16 `bun:"spellcategory_5"`
	Spellcategorycooldown5 int `bun:"spellcategorycooldown_5"`
	Bonding int8 `bun:"bonding"`
	Description string `bun:"description"`
	PageText int `bun:"PageText"`
	LanguageID int8 `bun:"LanguageID"`
	PageMaterial int8 `bun:"PageMaterial"`
	Startquest int `bun:"startquest"`
	Lockid int `bun:"lockid"`
	Material int8 `bun:"Material"`
	Sheath int8 `bun:"sheath"`
	RandomProperty int `bun:"RandomProperty"`
	RandomSuffix int `bun:"RandomSuffix"`
	Block int `bun:"block"`
	Itemset int `bun:"itemset"`
	MaxDurability int16 `bun:"MaxDurability"`
	Area int `bun:"area"`
	Map int16 `bun:"Map"`
	BagFamily int `bun:"BagFamily"`
	TotemCategory int `bun:"TotemCategory"`
	SocketColor1 int8 `bun:"socketColor_1"`
	SocketContent1 int `bun:"socketContent_1"`
	SocketColor2 int8 `bun:"socketColor_2"`
	SocketContent2 int `bun:"socketContent_2"`
	SocketColor3 int8 `bun:"socketColor_3"`
	SocketContent3 int `bun:"socketContent_3"`
	SocketBonus int `bun:"socketBonus"`
	GemProperties int `bun:"GemProperties"`
	RequiredDisenchantSkill int16 `bun:"RequiredDisenchantSkill"`
	ArmorDamageModifier float64 `bun:"ArmorDamageModifier"`
	Duration int `bun:"duration"`
	ItemLimitCategory int16 `bun:"ItemLimitCategory"`
	HolidayId int `bun:"HolidayId"`
	ScriptName string `bun:"ScriptName"`
	DisenchantID int `bun:"DisenchantID"`
	FoodType int8 `bun:"FoodType"`
	MinMoneyLoot int `bun:"minMoneyLoot"`
	MaxMoneyLoot int `bun:"maxMoneyLoot"`
	FlagsCustom int `bun:"flagsCustom"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBItemTemplateSlice []DBItemTemplate

func (entry *ItemTemplate) ToDB() *DBItemTemplate {
	if entry == nil {
		return nil
	}
	return &DBItemTemplate{
		Entry: entry.Entry,
		Class: entry.Class,
		Subclass: entry.Subclass,
		SoundOverrideSubclass: entry.SoundOverrideSubclass,
		Name: entry.Name,
		Displayid: entry.Displayid,
		Quality: entry.Quality,
		Flags: entry.Flags,
		FlagsExtra: entry.FlagsExtra,
		BuyCount: entry.BuyCount,
		BuyPrice: entry.BuyPrice,
		SellPrice: entry.SellPrice,
		InventoryType: entry.InventoryType,
		AllowableClass: entry.AllowableClass,
		AllowableRace: entry.AllowableRace,
		ItemLevel: entry.ItemLevel,
		RequiredLevel: entry.RequiredLevel,
		RequiredSkill: entry.RequiredSkill,
		RequiredSkillRank: entry.RequiredSkillRank,
		Requiredspell: entry.Requiredspell,
		Requiredhonorrank: entry.Requiredhonorrank,
		RequiredCityRank: entry.RequiredCityRank,
		RequiredReputationFaction: entry.RequiredReputationFaction,
		RequiredReputationRank: entry.RequiredReputationRank,
		Maxcount: entry.Maxcount,
		Stackable: entry.Stackable,
		ContainerSlots: entry.ContainerSlots,
		StatsCount: entry.StatsCount,
		StatType1: entry.StatType1,
		StatValue1: entry.StatValue1,
		StatType2: entry.StatType2,
		StatValue2: entry.StatValue2,
		StatType3: entry.StatType3,
		StatValue3: entry.StatValue3,
		StatType4: entry.StatType4,
		StatValue4: entry.StatValue4,
		StatType5: entry.StatType5,
		StatValue5: entry.StatValue5,
		StatType6: entry.StatType6,
		StatValue6: entry.StatValue6,
		StatType7: entry.StatType7,
		StatValue7: entry.StatValue7,
		StatType8: entry.StatType8,
		StatValue8: entry.StatValue8,
		StatType9: entry.StatType9,
		StatValue9: entry.StatValue9,
		StatType10: entry.StatType10,
		StatValue10: entry.StatValue10,
		ScalingStatDistribution: entry.ScalingStatDistribution,
		ScalingStatValue: entry.ScalingStatValue,
		DmgMin1: entry.DmgMin1,
		DmgMax1: entry.DmgMax1,
		DmgType1: entry.DmgType1,
		DmgMin2: entry.DmgMin2,
		DmgMax2: entry.DmgMax2,
		DmgType2: entry.DmgType2,
		Armor: entry.Armor,
		HolyRes: entry.HolyRes,
		FireRes: entry.FireRes,
		NatureRes: entry.NatureRes,
		FrostRes: entry.FrostRes,
		ShadowRes: entry.ShadowRes,
		ArcaneRes: entry.ArcaneRes,
		Delay: entry.Delay,
		AmmoType: entry.AmmoType,
		RangedModRange: entry.RangedModRange,
		Spellid1: entry.Spellid1,
		Spelltrigger1: entry.Spelltrigger1,
		Spellcharges1: entry.Spellcharges1,
		SpellppmRate1: entry.SpellppmRate1,
		Spellcooldown1: entry.Spellcooldown1,
		Spellcategory1: entry.Spellcategory1,
		Spellcategorycooldown1: entry.Spellcategorycooldown1,
		Spellid2: entry.Spellid2,
		Spelltrigger2: entry.Spelltrigger2,
		Spellcharges2: entry.Spellcharges2,
		SpellppmRate2: entry.SpellppmRate2,
		Spellcooldown2: entry.Spellcooldown2,
		Spellcategory2: entry.Spellcategory2,
		Spellcategorycooldown2: entry.Spellcategorycooldown2,
		Spellid3: entry.Spellid3,
		Spelltrigger3: entry.Spelltrigger3,
		Spellcharges3: entry.Spellcharges3,
		SpellppmRate3: entry.SpellppmRate3,
		Spellcooldown3: entry.Spellcooldown3,
		Spellcategory3: entry.Spellcategory3,
		Spellcategorycooldown3: entry.Spellcategorycooldown3,
		Spellid4: entry.Spellid4,
		Spelltrigger4: entry.Spelltrigger4,
		Spellcharges4: entry.Spellcharges4,
		SpellppmRate4: entry.SpellppmRate4,
		Spellcooldown4: entry.Spellcooldown4,
		Spellcategory4: entry.Spellcategory4,
		Spellcategorycooldown4: entry.Spellcategorycooldown4,
		Spellid5: entry.Spellid5,
		Spelltrigger5: entry.Spelltrigger5,
		Spellcharges5: entry.Spellcharges5,
		SpellppmRate5: entry.SpellppmRate5,
		Spellcooldown5: entry.Spellcooldown5,
		Spellcategory5: entry.Spellcategory5,
		Spellcategorycooldown5: entry.Spellcategorycooldown5,
		Bonding: entry.Bonding,
		Description: entry.Description,
		PageText: entry.PageText,
		LanguageID: entry.LanguageID,
		PageMaterial: entry.PageMaterial,
		Startquest: entry.Startquest,
		Lockid: entry.Lockid,
		Material: entry.Material,
		Sheath: entry.Sheath,
		RandomProperty: entry.RandomProperty,
		RandomSuffix: entry.RandomSuffix,
		Block: entry.Block,
		Itemset: entry.Itemset,
		MaxDurability: entry.MaxDurability,
		Area: entry.Area,
		Map: entry.Map,
		BagFamily: entry.BagFamily,
		TotemCategory: entry.TotemCategory,
		SocketColor1: entry.SocketColor1,
		SocketContent1: entry.SocketContent1,
		SocketColor2: entry.SocketColor2,
		SocketContent2: entry.SocketContent2,
		SocketColor3: entry.SocketColor3,
		SocketContent3: entry.SocketContent3,
		SocketBonus: entry.SocketBonus,
		GemProperties: entry.GemProperties,
		RequiredDisenchantSkill: entry.RequiredDisenchantSkill,
		ArmorDamageModifier: entry.ArmorDamageModifier,
		Duration: entry.Duration,
		ItemLimitCategory: entry.ItemLimitCategory,
		HolidayId: entry.HolidayId,
		ScriptName: entry.ScriptName,
		DisenchantID: entry.DisenchantID,
		FoodType: entry.FoodType,
		MinMoneyLoot: entry.MinMoneyLoot,
		MaxMoneyLoot: entry.MaxMoneyLoot,
		FlagsCustom: entry.FlagsCustom,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBItemTemplate) ToWeb() *ItemTemplate {
	if entry == nil {
		return nil
	}
	return &ItemTemplate{
		Entry: entry.Entry,
		Class: entry.Class,
		Subclass: entry.Subclass,
		SoundOverrideSubclass: entry.SoundOverrideSubclass,
		Name: entry.Name,
		Displayid: entry.Displayid,
		Quality: entry.Quality,
		Flags: entry.Flags,
		FlagsExtra: entry.FlagsExtra,
		BuyCount: entry.BuyCount,
		BuyPrice: entry.BuyPrice,
		SellPrice: entry.SellPrice,
		InventoryType: entry.InventoryType,
		AllowableClass: entry.AllowableClass,
		AllowableRace: entry.AllowableRace,
		ItemLevel: entry.ItemLevel,
		RequiredLevel: entry.RequiredLevel,
		RequiredSkill: entry.RequiredSkill,
		RequiredSkillRank: entry.RequiredSkillRank,
		Requiredspell: entry.Requiredspell,
		Requiredhonorrank: entry.Requiredhonorrank,
		RequiredCityRank: entry.RequiredCityRank,
		RequiredReputationFaction: entry.RequiredReputationFaction,
		RequiredReputationRank: entry.RequiredReputationRank,
		Maxcount: entry.Maxcount,
		Stackable: entry.Stackable,
		ContainerSlots: entry.ContainerSlots,
		StatsCount: entry.StatsCount,
		StatType1: entry.StatType1,
		StatValue1: entry.StatValue1,
		StatType2: entry.StatType2,
		StatValue2: entry.StatValue2,
		StatType3: entry.StatType3,
		StatValue3: entry.StatValue3,
		StatType4: entry.StatType4,
		StatValue4: entry.StatValue4,
		StatType5: entry.StatType5,
		StatValue5: entry.StatValue5,
		StatType6: entry.StatType6,
		StatValue6: entry.StatValue6,
		StatType7: entry.StatType7,
		StatValue7: entry.StatValue7,
		StatType8: entry.StatType8,
		StatValue8: entry.StatValue8,
		StatType9: entry.StatType9,
		StatValue9: entry.StatValue9,
		StatType10: entry.StatType10,
		StatValue10: entry.StatValue10,
		ScalingStatDistribution: entry.ScalingStatDistribution,
		ScalingStatValue: entry.ScalingStatValue,
		DmgMin1: entry.DmgMin1,
		DmgMax1: entry.DmgMax1,
		DmgType1: entry.DmgType1,
		DmgMin2: entry.DmgMin2,
		DmgMax2: entry.DmgMax2,
		DmgType2: entry.DmgType2,
		Armor: entry.Armor,
		HolyRes: entry.HolyRes,
		FireRes: entry.FireRes,
		NatureRes: entry.NatureRes,
		FrostRes: entry.FrostRes,
		ShadowRes: entry.ShadowRes,
		ArcaneRes: entry.ArcaneRes,
		Delay: entry.Delay,
		AmmoType: entry.AmmoType,
		RangedModRange: entry.RangedModRange,
		Spellid1: entry.Spellid1,
		Spelltrigger1: entry.Spelltrigger1,
		Spellcharges1: entry.Spellcharges1,
		SpellppmRate1: entry.SpellppmRate1,
		Spellcooldown1: entry.Spellcooldown1,
		Spellcategory1: entry.Spellcategory1,
		Spellcategorycooldown1: entry.Spellcategorycooldown1,
		Spellid2: entry.Spellid2,
		Spelltrigger2: entry.Spelltrigger2,
		Spellcharges2: entry.Spellcharges2,
		SpellppmRate2: entry.SpellppmRate2,
		Spellcooldown2: entry.Spellcooldown2,
		Spellcategory2: entry.Spellcategory2,
		Spellcategorycooldown2: entry.Spellcategorycooldown2,
		Spellid3: entry.Spellid3,
		Spelltrigger3: entry.Spelltrigger3,
		Spellcharges3: entry.Spellcharges3,
		SpellppmRate3: entry.SpellppmRate3,
		Spellcooldown3: entry.Spellcooldown3,
		Spellcategory3: entry.Spellcategory3,
		Spellcategorycooldown3: entry.Spellcategorycooldown3,
		Spellid4: entry.Spellid4,
		Spelltrigger4: entry.Spelltrigger4,
		Spellcharges4: entry.Spellcharges4,
		SpellppmRate4: entry.SpellppmRate4,
		Spellcooldown4: entry.Spellcooldown4,
		Spellcategory4: entry.Spellcategory4,
		Spellcategorycooldown4: entry.Spellcategorycooldown4,
		Spellid5: entry.Spellid5,
		Spelltrigger5: entry.Spelltrigger5,
		Spellcharges5: entry.Spellcharges5,
		SpellppmRate5: entry.SpellppmRate5,
		Spellcooldown5: entry.Spellcooldown5,
		Spellcategory5: entry.Spellcategory5,
		Spellcategorycooldown5: entry.Spellcategorycooldown5,
		Bonding: entry.Bonding,
		Description: entry.Description,
		PageText: entry.PageText,
		LanguageID: entry.LanguageID,
		PageMaterial: entry.PageMaterial,
		Startquest: entry.Startquest,
		Lockid: entry.Lockid,
		Material: entry.Material,
		Sheath: entry.Sheath,
		RandomProperty: entry.RandomProperty,
		RandomSuffix: entry.RandomSuffix,
		Block: entry.Block,
		Itemset: entry.Itemset,
		MaxDurability: entry.MaxDurability,
		Area: entry.Area,
		Map: entry.Map,
		BagFamily: entry.BagFamily,
		TotemCategory: entry.TotemCategory,
		SocketColor1: entry.SocketColor1,
		SocketContent1: entry.SocketContent1,
		SocketColor2: entry.SocketColor2,
		SocketContent2: entry.SocketContent2,
		SocketColor3: entry.SocketColor3,
		SocketContent3: entry.SocketContent3,
		SocketBonus: entry.SocketBonus,
		GemProperties: entry.GemProperties,
		RequiredDisenchantSkill: entry.RequiredDisenchantSkill,
		ArmorDamageModifier: entry.ArmorDamageModifier,
		Duration: entry.Duration,
		ItemLimitCategory: entry.ItemLimitCategory,
		HolidayId: entry.HolidayId,
		ScriptName: entry.ScriptName,
		DisenchantID: entry.DisenchantID,
		FoodType: entry.FoodType,
		MinMoneyLoot: entry.MinMoneyLoot,
		MaxMoneyLoot: entry.MaxMoneyLoot,
		FlagsCustom: entry.FlagsCustom,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data ItemTemplateSlice) ToDB() DBItemTemplateSlice {
	result := make(DBItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemTemplateSlice) ToWeb() ItemTemplateSlice {
	result := make(ItemTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
