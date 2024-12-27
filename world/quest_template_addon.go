package model

import "github.com/uptrace/bun"

type QuestTemplateAddon struct {
	ID int `json:"id,omitempty"`
	MaxLevel int8 `json:"maxlevel,omitempty"`
	AllowableClasses int `json:"allowableclasses,omitempty"`
	SourceSpellID int `json:"sourcespellid,omitempty"`
	PrevQuestID int `json:"prevquestid,omitempty"`
	NextQuestID int `json:"nextquestid,omitempty"`
	ExclusiveGroup int `json:"exclusivegroup,omitempty"`
	BreadcrumbForQuestId int `json:"breadcrumbforquestid,omitempty"`
	RewardMailTemplateID int `json:"rewardmailtemplateid,omitempty"`
	RewardMailDelay int `json:"rewardmaildelay,omitempty"`
	RequiredSkillID int16 `json:"requiredskillid,omitempty"`
	RequiredSkillPoints int16 `json:"requiredskillpoints,omitempty"`
	RequiredMinRepFaction int16 `json:"requiredminrepfaction,omitempty"`
	RequiredMaxRepFaction int16 `json:"requiredmaxrepfaction,omitempty"`
	RequiredMinRepValue int `json:"requiredminrepvalue,omitempty"`
	RequiredMaxRepValue int `json:"requiredmaxrepvalue,omitempty"`
	ProvidedItemCount int8 `json:"provideditemcount,omitempty"`
	SpecialFlags int8 `json:"specialflags,omitempty"`
}

type QuestTemplateAddonSlice []QuestTemplateAddon

type DBQuestTemplateAddon struct {
	bun.BaseModel `bun:"table:quest_template_addon,alias:quest_template_addon"`
	ID int `bun:"ID"`
	MaxLevel int8 `bun:"MaxLevel"`
	AllowableClasses int `bun:"AllowableClasses"`
	SourceSpellID int `bun:"SourceSpellID"`
	PrevQuestID int `bun:"PrevQuestID"`
	NextQuestID int `bun:"NextQuestID"`
	ExclusiveGroup int `bun:"ExclusiveGroup"`
	BreadcrumbForQuestId int `bun:"BreadcrumbForQuestId"`
	RewardMailTemplateID int `bun:"RewardMailTemplateID"`
	RewardMailDelay int `bun:"RewardMailDelay"`
	RequiredSkillID int16 `bun:"RequiredSkillID"`
	RequiredSkillPoints int16 `bun:"RequiredSkillPoints"`
	RequiredMinRepFaction int16 `bun:"RequiredMinRepFaction"`
	RequiredMaxRepFaction int16 `bun:"RequiredMaxRepFaction"`
	RequiredMinRepValue int `bun:"RequiredMinRepValue"`
	RequiredMaxRepValue int `bun:"RequiredMaxRepValue"`
	ProvidedItemCount int8 `bun:"ProvidedItemCount"`
	SpecialFlags int8 `bun:"SpecialFlags"`
}

type DBQuestTemplateAddonSlice []DBQuestTemplateAddon

func (entry *QuestTemplateAddon) ToDB() *DBQuestTemplateAddon {
	if entry == nil {
		return nil
	}
	return &DBQuestTemplateAddon{
		ID: entry.ID,
		MaxLevel: entry.MaxLevel,
		AllowableClasses: entry.AllowableClasses,
		SourceSpellID: entry.SourceSpellID,
		PrevQuestID: entry.PrevQuestID,
		NextQuestID: entry.NextQuestID,
		ExclusiveGroup: entry.ExclusiveGroup,
		BreadcrumbForQuestId: entry.BreadcrumbForQuestId,
		RewardMailTemplateID: entry.RewardMailTemplateID,
		RewardMailDelay: entry.RewardMailDelay,
		RequiredSkillID: entry.RequiredSkillID,
		RequiredSkillPoints: entry.RequiredSkillPoints,
		RequiredMinRepFaction: entry.RequiredMinRepFaction,
		RequiredMaxRepFaction: entry.RequiredMaxRepFaction,
		RequiredMinRepValue: entry.RequiredMinRepValue,
		RequiredMaxRepValue: entry.RequiredMaxRepValue,
		ProvidedItemCount: entry.ProvidedItemCount,
		SpecialFlags: entry.SpecialFlags,
	}
}

func (entry *DBQuestTemplateAddon) ToWeb() *QuestTemplateAddon {
	if entry == nil {
		return nil
	}
	return &QuestTemplateAddon{
		ID: entry.ID,
		MaxLevel: entry.MaxLevel,
		AllowableClasses: entry.AllowableClasses,
		SourceSpellID: entry.SourceSpellID,
		PrevQuestID: entry.PrevQuestID,
		NextQuestID: entry.NextQuestID,
		ExclusiveGroup: entry.ExclusiveGroup,
		BreadcrumbForQuestId: entry.BreadcrumbForQuestId,
		RewardMailTemplateID: entry.RewardMailTemplateID,
		RewardMailDelay: entry.RewardMailDelay,
		RequiredSkillID: entry.RequiredSkillID,
		RequiredSkillPoints: entry.RequiredSkillPoints,
		RequiredMinRepFaction: entry.RequiredMinRepFaction,
		RequiredMaxRepFaction: entry.RequiredMaxRepFaction,
		RequiredMinRepValue: entry.RequiredMinRepValue,
		RequiredMaxRepValue: entry.RequiredMaxRepValue,
		ProvidedItemCount: entry.ProvidedItemCount,
		SpecialFlags: entry.SpecialFlags,
	}
}

func (data QuestTemplateAddonSlice) ToDB() DBQuestTemplateAddonSlice {
	result := make(DBQuestTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestTemplateAddonSlice) ToWeb() QuestTemplateAddonSlice {
	result := make(QuestTemplateAddonSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
