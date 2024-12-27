package model

import "github.com/uptrace/bun"

type SpellDbc struct {
	Id int `json:"id,omitempty"`
	Dispel int `json:"dispel,omitempty"`
	Mechanic int `json:"mechanic,omitempty"`
	Attributes int `json:"attributes,omitempty"`
	AttributesEx int `json:"attributesex,omitempty"`
	AttributesEx2 int `json:"attributesex2,omitempty"`
	AttributesEx3 int `json:"attributesex3,omitempty"`
	AttributesEx4 int `json:"attributesex4,omitempty"`
	AttributesEx5 int `json:"attributesex5,omitempty"`
	AttributesEx6 int `json:"attributesex6,omitempty"`
	AttributesEx7 int `json:"attributesex7,omitempty"`
	Stances int `json:"stances,omitempty"`
	StancesNot int `json:"stancesnot,omitempty"`
	Targets int `json:"targets,omitempty"`
	CastingTimeIndex int `json:"castingtimeindex,omitempty"`
	AuraInterruptFlags int `json:"aurainterruptflags,omitempty"`
	ProcFlags int `json:"procflags,omitempty"`
	ProcChance int `json:"procchance,omitempty"`
	ProcCharges int `json:"proccharges,omitempty"`
	MaxLevel int `json:"maxlevel,omitempty"`
	BaseLevel int `json:"baselevel,omitempty"`
	SpellLevel int `json:"spelllevel,omitempty"`
	DurationIndex int `json:"durationindex,omitempty"`
	RangeIndex int `json:"rangeindex,omitempty"`
	StackAmount int `json:"stackamount,omitempty"`
	EquippedItemClass int `json:"equippeditemclass,omitempty"`
	EquippedItemSubClassMask int `json:"equippeditemsubclassmask,omitempty"`
	EquippedItemInventoryTypeMask int `json:"equippediteminventorytypemask,omitempty"`
	Effect1 int `json:"effect1,omitempty"`
	Effect2 int `json:"effect2,omitempty"`
	Effect3 int `json:"effect3,omitempty"`
	EffectDieSides1 int `json:"effectdiesides1,omitempty"`
	EffectDieSides2 int `json:"effectdiesides2,omitempty"`
	EffectDieSides3 int `json:"effectdiesides3,omitempty"`
	EffectRealPointsPerLevel1 float64 `json:"effectrealpointsperlevel1,omitempty"`
	EffectRealPointsPerLevel2 float64 `json:"effectrealpointsperlevel2,omitempty"`
	EffectRealPointsPerLevel3 float64 `json:"effectrealpointsperlevel3,omitempty"`
	EffectBasePoints1 int `json:"effectbasepoints1,omitempty"`
	EffectBasePoints2 int `json:"effectbasepoints2,omitempty"`
	EffectBasePoints3 int `json:"effectbasepoints3,omitempty"`
	EffectMechanic1 int `json:"effectmechanic1,omitempty"`
	EffectMechanic2 int `json:"effectmechanic2,omitempty"`
	EffectMechanic3 int `json:"effectmechanic3,omitempty"`
	EffectImplicitTargetA1 int `json:"effectimplicittargeta1,omitempty"`
	EffectImplicitTargetA2 int `json:"effectimplicittargeta2,omitempty"`
	EffectImplicitTargetA3 int `json:"effectimplicittargeta3,omitempty"`
	EffectImplicitTargetB1 int `json:"effectimplicittargetb1,omitempty"`
	EffectImplicitTargetB2 int `json:"effectimplicittargetb2,omitempty"`
	EffectImplicitTargetB3 int `json:"effectimplicittargetb3,omitempty"`
	EffectRadiusIndex1 int `json:"effectradiusindex1,omitempty"`
	EffectRadiusIndex2 int `json:"effectradiusindex2,omitempty"`
	EffectRadiusIndex3 int `json:"effectradiusindex3,omitempty"`
	EffectApplyAuraName1 int `json:"effectapplyauraname1,omitempty"`
	EffectApplyAuraName2 int `json:"effectapplyauraname2,omitempty"`
	EffectApplyAuraName3 int `json:"effectapplyauraname3,omitempty"`
	EffectAmplitude1 int `json:"effectamplitude1,omitempty"`
	EffectAmplitude2 int `json:"effectamplitude2,omitempty"`
	EffectAmplitude3 int `json:"effectamplitude3,omitempty"`
	EffectMultipleValue1 float64 `json:"effectmultiplevalue1,omitempty"`
	EffectMultipleValue2 float64 `json:"effectmultiplevalue2,omitempty"`
	EffectMultipleValue3 float64 `json:"effectmultiplevalue3,omitempty"`
	EffectItemType1 int `json:"effectitemtype1,omitempty"`
	EffectItemType2 int `json:"effectitemtype2,omitempty"`
	EffectItemType3 int `json:"effectitemtype3,omitempty"`
	EffectMiscValue1 int `json:"effectmiscvalue1,omitempty"`
	EffectMiscValue2 int `json:"effectmiscvalue2,omitempty"`
	EffectMiscValue3 int `json:"effectmiscvalue3,omitempty"`
	EffectMiscValueB1 int `json:"effectmiscvalueb1,omitempty"`
	EffectMiscValueB2 int `json:"effectmiscvalueb2,omitempty"`
	EffectMiscValueB3 int `json:"effectmiscvalueb3,omitempty"`
	EffectTriggerSpell1 int `json:"effecttriggerspell1,omitempty"`
	EffectTriggerSpell2 int `json:"effecttriggerspell2,omitempty"`
	EffectTriggerSpell3 int `json:"effecttriggerspell3,omitempty"`
	EffectSpellClassMaskA1 int `json:"effectspellclassmaska1,omitempty"`
	EffectSpellClassMaskA2 int `json:"effectspellclassmaska2,omitempty"`
	EffectSpellClassMaskA3 int `json:"effectspellclassmaska3,omitempty"`
	EffectSpellClassMaskB1 int `json:"effectspellclassmaskb1,omitempty"`
	EffectSpellClassMaskB2 int `json:"effectspellclassmaskb2,omitempty"`
	EffectSpellClassMaskB3 int `json:"effectspellclassmaskb3,omitempty"`
	EffectSpellClassMaskC1 int `json:"effectspellclassmaskc1,omitempty"`
	EffectSpellClassMaskC2 int `json:"effectspellclassmaskc2,omitempty"`
	EffectSpellClassMaskC3 int `json:"effectspellclassmaskc3,omitempty"`
	SpellName string `json:"spellname,omitempty"`
	MaxTargetLevel int `json:"maxtargetlevel,omitempty"`
	SpellFamilyName int `json:"spellfamilyname,omitempty"`
	SpellFamilyFlags1 int `json:"spellfamilyflags1,omitempty"`
	SpellFamilyFlags2 int `json:"spellfamilyflags2,omitempty"`
	SpellFamilyFlags3 int `json:"spellfamilyflags3,omitempty"`
	MaxAffectedTargets int `json:"maxaffectedtargets,omitempty"`
	DmgClass int `json:"dmgclass,omitempty"`
	PreventionType int `json:"preventiontype,omitempty"`
	DmgMultiplier1 float64 `json:"dmgmultiplier1,omitempty"`
	DmgMultiplier2 float64 `json:"dmgmultiplier2,omitempty"`
	DmgMultiplier3 float64 `json:"dmgmultiplier3,omitempty"`
	AreaGroupId int `json:"areagroupid,omitempty"`
	SchoolMask int `json:"schoolmask,omitempty"`
}

type SpellDbcSlice []SpellDbc

type DBSpellDbc struct {
	bun.BaseModel `bun:"table:spell_dbc,alias:spell_dbc"`
	Id int `bun:"Id"`
	Dispel int `bun:"Dispel"`
	Mechanic int `bun:"Mechanic"`
	Attributes int `bun:"Attributes"`
	AttributesEx int `bun:"AttributesEx"`
	AttributesEx2 int `bun:"AttributesEx2"`
	AttributesEx3 int `bun:"AttributesEx3"`
	AttributesEx4 int `bun:"AttributesEx4"`
	AttributesEx5 int `bun:"AttributesEx5"`
	AttributesEx6 int `bun:"AttributesEx6"`
	AttributesEx7 int `bun:"AttributesEx7"`
	Stances int `bun:"Stances"`
	StancesNot int `bun:"StancesNot"`
	Targets int `bun:"Targets"`
	CastingTimeIndex int `bun:"CastingTimeIndex"`
	AuraInterruptFlags int `bun:"AuraInterruptFlags"`
	ProcFlags int `bun:"ProcFlags"`
	ProcChance int `bun:"ProcChance"`
	ProcCharges int `bun:"ProcCharges"`
	MaxLevel int `bun:"MaxLevel"`
	BaseLevel int `bun:"BaseLevel"`
	SpellLevel int `bun:"SpellLevel"`
	DurationIndex int `bun:"DurationIndex"`
	RangeIndex int `bun:"RangeIndex"`
	StackAmount int `bun:"StackAmount"`
	EquippedItemClass int `bun:"EquippedItemClass"`
	EquippedItemSubClassMask int `bun:"EquippedItemSubClassMask"`
	EquippedItemInventoryTypeMask int `bun:"EquippedItemInventoryTypeMask"`
	Effect1 int `bun:"Effect1"`
	Effect2 int `bun:"Effect2"`
	Effect3 int `bun:"Effect3"`
	EffectDieSides1 int `bun:"EffectDieSides1"`
	EffectDieSides2 int `bun:"EffectDieSides2"`
	EffectDieSides3 int `bun:"EffectDieSides3"`
	EffectRealPointsPerLevel1 float64 `bun:"EffectRealPointsPerLevel1"`
	EffectRealPointsPerLevel2 float64 `bun:"EffectRealPointsPerLevel2"`
	EffectRealPointsPerLevel3 float64 `bun:"EffectRealPointsPerLevel3"`
	EffectBasePoints1 int `bun:"EffectBasePoints1"`
	EffectBasePoints2 int `bun:"EffectBasePoints2"`
	EffectBasePoints3 int `bun:"EffectBasePoints3"`
	EffectMechanic1 int `bun:"EffectMechanic1"`
	EffectMechanic2 int `bun:"EffectMechanic2"`
	EffectMechanic3 int `bun:"EffectMechanic3"`
	EffectImplicitTargetA1 int `bun:"EffectImplicitTargetA1"`
	EffectImplicitTargetA2 int `bun:"EffectImplicitTargetA2"`
	EffectImplicitTargetA3 int `bun:"EffectImplicitTargetA3"`
	EffectImplicitTargetB1 int `bun:"EffectImplicitTargetB1"`
	EffectImplicitTargetB2 int `bun:"EffectImplicitTargetB2"`
	EffectImplicitTargetB3 int `bun:"EffectImplicitTargetB3"`
	EffectRadiusIndex1 int `bun:"EffectRadiusIndex1"`
	EffectRadiusIndex2 int `bun:"EffectRadiusIndex2"`
	EffectRadiusIndex3 int `bun:"EffectRadiusIndex3"`
	EffectApplyAuraName1 int `bun:"EffectApplyAuraName1"`
	EffectApplyAuraName2 int `bun:"EffectApplyAuraName2"`
	EffectApplyAuraName3 int `bun:"EffectApplyAuraName3"`
	EffectAmplitude1 int `bun:"EffectAmplitude1"`
	EffectAmplitude2 int `bun:"EffectAmplitude2"`
	EffectAmplitude3 int `bun:"EffectAmplitude3"`
	EffectMultipleValue1 float64 `bun:"EffectMultipleValue1"`
	EffectMultipleValue2 float64 `bun:"EffectMultipleValue2"`
	EffectMultipleValue3 float64 `bun:"EffectMultipleValue3"`
	EffectItemType1 int `bun:"EffectItemType1"`
	EffectItemType2 int `bun:"EffectItemType2"`
	EffectItemType3 int `bun:"EffectItemType3"`
	EffectMiscValue1 int `bun:"EffectMiscValue1"`
	EffectMiscValue2 int `bun:"EffectMiscValue2"`
	EffectMiscValue3 int `bun:"EffectMiscValue3"`
	EffectMiscValueB1 int `bun:"EffectMiscValueB1"`
	EffectMiscValueB2 int `bun:"EffectMiscValueB2"`
	EffectMiscValueB3 int `bun:"EffectMiscValueB3"`
	EffectTriggerSpell1 int `bun:"EffectTriggerSpell1"`
	EffectTriggerSpell2 int `bun:"EffectTriggerSpell2"`
	EffectTriggerSpell3 int `bun:"EffectTriggerSpell3"`
	EffectSpellClassMaskA1 int `bun:"EffectSpellClassMaskA1"`
	EffectSpellClassMaskA2 int `bun:"EffectSpellClassMaskA2"`
	EffectSpellClassMaskA3 int `bun:"EffectSpellClassMaskA3"`
	EffectSpellClassMaskB1 int `bun:"EffectSpellClassMaskB1"`
	EffectSpellClassMaskB2 int `bun:"EffectSpellClassMaskB2"`
	EffectSpellClassMaskB3 int `bun:"EffectSpellClassMaskB3"`
	EffectSpellClassMaskC1 int `bun:"EffectSpellClassMaskC1"`
	EffectSpellClassMaskC2 int `bun:"EffectSpellClassMaskC2"`
	EffectSpellClassMaskC3 int `bun:"EffectSpellClassMaskC3"`
	SpellName string `bun:"SpellName"`
	MaxTargetLevel int `bun:"MaxTargetLevel"`
	SpellFamilyName int `bun:"SpellFamilyName"`
	SpellFamilyFlags1 int `bun:"SpellFamilyFlags1"`
	SpellFamilyFlags2 int `bun:"SpellFamilyFlags2"`
	SpellFamilyFlags3 int `bun:"SpellFamilyFlags3"`
	MaxAffectedTargets int `bun:"MaxAffectedTargets"`
	DmgClass int `bun:"DmgClass"`
	PreventionType int `bun:"PreventionType"`
	DmgMultiplier1 float64 `bun:"DmgMultiplier1"`
	DmgMultiplier2 float64 `bun:"DmgMultiplier2"`
	DmgMultiplier3 float64 `bun:"DmgMultiplier3"`
	AreaGroupId int `bun:"AreaGroupId"`
	SchoolMask int `bun:"SchoolMask"`
}

type DBSpellDbcSlice []DBSpellDbc

func (entry *SpellDbc) ToDB() *DBSpellDbc {
	if entry == nil {
		return nil
	}
	return &DBSpellDbc{
		Id: entry.Id,
		Dispel: entry.Dispel,
		Mechanic: entry.Mechanic,
		Attributes: entry.Attributes,
		AttributesEx: entry.AttributesEx,
		AttributesEx2: entry.AttributesEx2,
		AttributesEx3: entry.AttributesEx3,
		AttributesEx4: entry.AttributesEx4,
		AttributesEx5: entry.AttributesEx5,
		AttributesEx6: entry.AttributesEx6,
		AttributesEx7: entry.AttributesEx7,
		Stances: entry.Stances,
		StancesNot: entry.StancesNot,
		Targets: entry.Targets,
		CastingTimeIndex: entry.CastingTimeIndex,
		AuraInterruptFlags: entry.AuraInterruptFlags,
		ProcFlags: entry.ProcFlags,
		ProcChance: entry.ProcChance,
		ProcCharges: entry.ProcCharges,
		MaxLevel: entry.MaxLevel,
		BaseLevel: entry.BaseLevel,
		SpellLevel: entry.SpellLevel,
		DurationIndex: entry.DurationIndex,
		RangeIndex: entry.RangeIndex,
		StackAmount: entry.StackAmount,
		EquippedItemClass: entry.EquippedItemClass,
		EquippedItemSubClassMask: entry.EquippedItemSubClassMask,
		EquippedItemInventoryTypeMask: entry.EquippedItemInventoryTypeMask,
		Effect1: entry.Effect1,
		Effect2: entry.Effect2,
		Effect3: entry.Effect3,
		EffectDieSides1: entry.EffectDieSides1,
		EffectDieSides2: entry.EffectDieSides2,
		EffectDieSides3: entry.EffectDieSides3,
		EffectRealPointsPerLevel1: entry.EffectRealPointsPerLevel1,
		EffectRealPointsPerLevel2: entry.EffectRealPointsPerLevel2,
		EffectRealPointsPerLevel3: entry.EffectRealPointsPerLevel3,
		EffectBasePoints1: entry.EffectBasePoints1,
		EffectBasePoints2: entry.EffectBasePoints2,
		EffectBasePoints3: entry.EffectBasePoints3,
		EffectMechanic1: entry.EffectMechanic1,
		EffectMechanic2: entry.EffectMechanic2,
		EffectMechanic3: entry.EffectMechanic3,
		EffectImplicitTargetA1: entry.EffectImplicitTargetA1,
		EffectImplicitTargetA2: entry.EffectImplicitTargetA2,
		EffectImplicitTargetA3: entry.EffectImplicitTargetA3,
		EffectImplicitTargetB1: entry.EffectImplicitTargetB1,
		EffectImplicitTargetB2: entry.EffectImplicitTargetB2,
		EffectImplicitTargetB3: entry.EffectImplicitTargetB3,
		EffectRadiusIndex1: entry.EffectRadiusIndex1,
		EffectRadiusIndex2: entry.EffectRadiusIndex2,
		EffectRadiusIndex3: entry.EffectRadiusIndex3,
		EffectApplyAuraName1: entry.EffectApplyAuraName1,
		EffectApplyAuraName2: entry.EffectApplyAuraName2,
		EffectApplyAuraName3: entry.EffectApplyAuraName3,
		EffectAmplitude1: entry.EffectAmplitude1,
		EffectAmplitude2: entry.EffectAmplitude2,
		EffectAmplitude3: entry.EffectAmplitude3,
		EffectMultipleValue1: entry.EffectMultipleValue1,
		EffectMultipleValue2: entry.EffectMultipleValue2,
		EffectMultipleValue3: entry.EffectMultipleValue3,
		EffectItemType1: entry.EffectItemType1,
		EffectItemType2: entry.EffectItemType2,
		EffectItemType3: entry.EffectItemType3,
		EffectMiscValue1: entry.EffectMiscValue1,
		EffectMiscValue2: entry.EffectMiscValue2,
		EffectMiscValue3: entry.EffectMiscValue3,
		EffectMiscValueB1: entry.EffectMiscValueB1,
		EffectMiscValueB2: entry.EffectMiscValueB2,
		EffectMiscValueB3: entry.EffectMiscValueB3,
		EffectTriggerSpell1: entry.EffectTriggerSpell1,
		EffectTriggerSpell2: entry.EffectTriggerSpell2,
		EffectTriggerSpell3: entry.EffectTriggerSpell3,
		EffectSpellClassMaskA1: entry.EffectSpellClassMaskA1,
		EffectSpellClassMaskA2: entry.EffectSpellClassMaskA2,
		EffectSpellClassMaskA3: entry.EffectSpellClassMaskA3,
		EffectSpellClassMaskB1: entry.EffectSpellClassMaskB1,
		EffectSpellClassMaskB2: entry.EffectSpellClassMaskB2,
		EffectSpellClassMaskB3: entry.EffectSpellClassMaskB3,
		EffectSpellClassMaskC1: entry.EffectSpellClassMaskC1,
		EffectSpellClassMaskC2: entry.EffectSpellClassMaskC2,
		EffectSpellClassMaskC3: entry.EffectSpellClassMaskC3,
		SpellName: entry.SpellName,
		MaxTargetLevel: entry.MaxTargetLevel,
		SpellFamilyName: entry.SpellFamilyName,
		SpellFamilyFlags1: entry.SpellFamilyFlags1,
		SpellFamilyFlags2: entry.SpellFamilyFlags2,
		SpellFamilyFlags3: entry.SpellFamilyFlags3,
		MaxAffectedTargets: entry.MaxAffectedTargets,
		DmgClass: entry.DmgClass,
		PreventionType: entry.PreventionType,
		DmgMultiplier1: entry.DmgMultiplier1,
		DmgMultiplier2: entry.DmgMultiplier2,
		DmgMultiplier3: entry.DmgMultiplier3,
		AreaGroupId: entry.AreaGroupId,
		SchoolMask: entry.SchoolMask,
	}
}

func (entry *DBSpellDbc) ToWeb() *SpellDbc {
	if entry == nil {
		return nil
	}
	return &SpellDbc{
		Id: entry.Id,
		Dispel: entry.Dispel,
		Mechanic: entry.Mechanic,
		Attributes: entry.Attributes,
		AttributesEx: entry.AttributesEx,
		AttributesEx2: entry.AttributesEx2,
		AttributesEx3: entry.AttributesEx3,
		AttributesEx4: entry.AttributesEx4,
		AttributesEx5: entry.AttributesEx5,
		AttributesEx6: entry.AttributesEx6,
		AttributesEx7: entry.AttributesEx7,
		Stances: entry.Stances,
		StancesNot: entry.StancesNot,
		Targets: entry.Targets,
		CastingTimeIndex: entry.CastingTimeIndex,
		AuraInterruptFlags: entry.AuraInterruptFlags,
		ProcFlags: entry.ProcFlags,
		ProcChance: entry.ProcChance,
		ProcCharges: entry.ProcCharges,
		MaxLevel: entry.MaxLevel,
		BaseLevel: entry.BaseLevel,
		SpellLevel: entry.SpellLevel,
		DurationIndex: entry.DurationIndex,
		RangeIndex: entry.RangeIndex,
		StackAmount: entry.StackAmount,
		EquippedItemClass: entry.EquippedItemClass,
		EquippedItemSubClassMask: entry.EquippedItemSubClassMask,
		EquippedItemInventoryTypeMask: entry.EquippedItemInventoryTypeMask,
		Effect1: entry.Effect1,
		Effect2: entry.Effect2,
		Effect3: entry.Effect3,
		EffectDieSides1: entry.EffectDieSides1,
		EffectDieSides2: entry.EffectDieSides2,
		EffectDieSides3: entry.EffectDieSides3,
		EffectRealPointsPerLevel1: entry.EffectRealPointsPerLevel1,
		EffectRealPointsPerLevel2: entry.EffectRealPointsPerLevel2,
		EffectRealPointsPerLevel3: entry.EffectRealPointsPerLevel3,
		EffectBasePoints1: entry.EffectBasePoints1,
		EffectBasePoints2: entry.EffectBasePoints2,
		EffectBasePoints3: entry.EffectBasePoints3,
		EffectMechanic1: entry.EffectMechanic1,
		EffectMechanic2: entry.EffectMechanic2,
		EffectMechanic3: entry.EffectMechanic3,
		EffectImplicitTargetA1: entry.EffectImplicitTargetA1,
		EffectImplicitTargetA2: entry.EffectImplicitTargetA2,
		EffectImplicitTargetA3: entry.EffectImplicitTargetA3,
		EffectImplicitTargetB1: entry.EffectImplicitTargetB1,
		EffectImplicitTargetB2: entry.EffectImplicitTargetB2,
		EffectImplicitTargetB3: entry.EffectImplicitTargetB3,
		EffectRadiusIndex1: entry.EffectRadiusIndex1,
		EffectRadiusIndex2: entry.EffectRadiusIndex2,
		EffectRadiusIndex3: entry.EffectRadiusIndex3,
		EffectApplyAuraName1: entry.EffectApplyAuraName1,
		EffectApplyAuraName2: entry.EffectApplyAuraName2,
		EffectApplyAuraName3: entry.EffectApplyAuraName3,
		EffectAmplitude1: entry.EffectAmplitude1,
		EffectAmplitude2: entry.EffectAmplitude2,
		EffectAmplitude3: entry.EffectAmplitude3,
		EffectMultipleValue1: entry.EffectMultipleValue1,
		EffectMultipleValue2: entry.EffectMultipleValue2,
		EffectMultipleValue3: entry.EffectMultipleValue3,
		EffectItemType1: entry.EffectItemType1,
		EffectItemType2: entry.EffectItemType2,
		EffectItemType3: entry.EffectItemType3,
		EffectMiscValue1: entry.EffectMiscValue1,
		EffectMiscValue2: entry.EffectMiscValue2,
		EffectMiscValue3: entry.EffectMiscValue3,
		EffectMiscValueB1: entry.EffectMiscValueB1,
		EffectMiscValueB2: entry.EffectMiscValueB2,
		EffectMiscValueB3: entry.EffectMiscValueB3,
		EffectTriggerSpell1: entry.EffectTriggerSpell1,
		EffectTriggerSpell2: entry.EffectTriggerSpell2,
		EffectTriggerSpell3: entry.EffectTriggerSpell3,
		EffectSpellClassMaskA1: entry.EffectSpellClassMaskA1,
		EffectSpellClassMaskA2: entry.EffectSpellClassMaskA2,
		EffectSpellClassMaskA3: entry.EffectSpellClassMaskA3,
		EffectSpellClassMaskB1: entry.EffectSpellClassMaskB1,
		EffectSpellClassMaskB2: entry.EffectSpellClassMaskB2,
		EffectSpellClassMaskB3: entry.EffectSpellClassMaskB3,
		EffectSpellClassMaskC1: entry.EffectSpellClassMaskC1,
		EffectSpellClassMaskC2: entry.EffectSpellClassMaskC2,
		EffectSpellClassMaskC3: entry.EffectSpellClassMaskC3,
		SpellName: entry.SpellName,
		MaxTargetLevel: entry.MaxTargetLevel,
		SpellFamilyName: entry.SpellFamilyName,
		SpellFamilyFlags1: entry.SpellFamilyFlags1,
		SpellFamilyFlags2: entry.SpellFamilyFlags2,
		SpellFamilyFlags3: entry.SpellFamilyFlags3,
		MaxAffectedTargets: entry.MaxAffectedTargets,
		DmgClass: entry.DmgClass,
		PreventionType: entry.PreventionType,
		DmgMultiplier1: entry.DmgMultiplier1,
		DmgMultiplier2: entry.DmgMultiplier2,
		DmgMultiplier3: entry.DmgMultiplier3,
		AreaGroupId: entry.AreaGroupId,
		SchoolMask: entry.SchoolMask,
	}
}

func (data SpellDbcSlice) ToDB() DBSpellDbcSlice {
	result := make(DBSpellDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBSpellDbcSlice) ToWeb() SpellDbcSlice {
	result := make(SpellDbcSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
