package model

import "github.com/uptrace/bun"

type QuestTemplate struct {
	ID int `json:"id,omitempty"`
	QuestType int8 `json:"questtype,omitempty"`
	QuestLevel int16 `json:"questlevel,omitempty"`
	MinLevel int8 `json:"minlevel,omitempty"`
	QuestSortID int16 `json:"questsortid,omitempty"`
	QuestInfoID int16 `json:"questinfoid,omitempty"`
	SuggestedGroupNum int8 `json:"suggestedgroupnum,omitempty"`
	RequiredFactionId1 int16 `json:"requiredfactionid1,omitempty"`
	RequiredFactionId2 int16 `json:"requiredfactionid2,omitempty"`
	RequiredFactionValue1 int `json:"requiredfactionvalue1,omitempty"`
	RequiredFactionValue2 int `json:"requiredfactionvalue2,omitempty"`
	RewardNextQuest int `json:"rewardnextquest,omitempty"`
	RewardXPDifficulty int8 `json:"rewardxpdifficulty,omitempty"`
	RewardMoney int `json:"rewardmoney,omitempty"`
	RewardBonusMoney int `json:"rewardbonusmoney,omitempty"`
	RewardDisplaySpell int `json:"rewarddisplayspell,omitempty"`
	RewardSpell int `json:"rewardspell,omitempty"`
	RewardHonor int `json:"rewardhonor,omitempty"`
	RewardKillHonor float64 `json:"rewardkillhonor,omitempty"`
	StartItem int `json:"startitem,omitempty"`
	Flags int `json:"flags,omitempty"`
	RequiredPlayerKills int8 `json:"requiredplayerkills,omitempty"`
	RewardItem1 int `json:"rewarditem1,omitempty"`
	RewardAmount1 int16 `json:"rewardamount1,omitempty"`
	RewardItem2 int `json:"rewarditem2,omitempty"`
	RewardAmount2 int16 `json:"rewardamount2,omitempty"`
	RewardItem3 int `json:"rewarditem3,omitempty"`
	RewardAmount3 int16 `json:"rewardamount3,omitempty"`
	RewardItem4 int `json:"rewarditem4,omitempty"`
	RewardAmount4 int16 `json:"rewardamount4,omitempty"`
	ItemDrop1 int `json:"itemdrop1,omitempty"`
	ItemDropQuantity1 int16 `json:"itemdropquantity1,omitempty"`
	ItemDrop2 int `json:"itemdrop2,omitempty"`
	ItemDropQuantity2 int16 `json:"itemdropquantity2,omitempty"`
	ItemDrop3 int `json:"itemdrop3,omitempty"`
	ItemDropQuantity3 int16 `json:"itemdropquantity3,omitempty"`
	ItemDrop4 int `json:"itemdrop4,omitempty"`
	ItemDropQuantity4 int16 `json:"itemdropquantity4,omitempty"`
	RewardChoiceItemID1 int `json:"rewardchoiceitemid1,omitempty"`
	RewardChoiceItemQuantity1 int16 `json:"rewardchoiceitemquantity1,omitempty"`
	RewardChoiceItemID2 int `json:"rewardchoiceitemid2,omitempty"`
	RewardChoiceItemQuantity2 int16 `json:"rewardchoiceitemquantity2,omitempty"`
	RewardChoiceItemID3 int `json:"rewardchoiceitemid3,omitempty"`
	RewardChoiceItemQuantity3 int16 `json:"rewardchoiceitemquantity3,omitempty"`
	RewardChoiceItemID4 int `json:"rewardchoiceitemid4,omitempty"`
	RewardChoiceItemQuantity4 int16 `json:"rewardchoiceitemquantity4,omitempty"`
	RewardChoiceItemID5 int `json:"rewardchoiceitemid5,omitempty"`
	RewardChoiceItemQuantity5 int16 `json:"rewardchoiceitemquantity5,omitempty"`
	RewardChoiceItemID6 int `json:"rewardchoiceitemid6,omitempty"`
	RewardChoiceItemQuantity6 int16 `json:"rewardchoiceitemquantity6,omitempty"`
	POIContinent int16 `json:"poicontinent,omitempty"`
	POIx float64 `json:"poix,omitempty"`
	POIy float64 `json:"poiy,omitempty"`
	POIPriority int `json:"poipriority,omitempty"`
	RewardTitle int8 `json:"rewardtitle,omitempty"`
	RewardTalents int8 `json:"rewardtalents,omitempty"`
	RewardArenaPoints int16 `json:"rewardarenapoints,omitempty"`
	RewardFactionID1 int16 `json:"rewardfactionid1,omitempty"`
	RewardFactionValue1 int `json:"rewardfactionvalue1,omitempty"`
	RewardFactionOverride1 int `json:"rewardfactionoverride1,omitempty"`
	RewardFactionID2 int16 `json:"rewardfactionid2,omitempty"`
	RewardFactionValue2 int `json:"rewardfactionvalue2,omitempty"`
	RewardFactionOverride2 int `json:"rewardfactionoverride2,omitempty"`
	RewardFactionID3 int16 `json:"rewardfactionid3,omitempty"`
	RewardFactionValue3 int `json:"rewardfactionvalue3,omitempty"`
	RewardFactionOverride3 int `json:"rewardfactionoverride3,omitempty"`
	RewardFactionID4 int16 `json:"rewardfactionid4,omitempty"`
	RewardFactionValue4 int `json:"rewardfactionvalue4,omitempty"`
	RewardFactionOverride4 int `json:"rewardfactionoverride4,omitempty"`
	RewardFactionID5 int16 `json:"rewardfactionid5,omitempty"`
	RewardFactionValue5 int `json:"rewardfactionvalue5,omitempty"`
	RewardFactionOverride5 int `json:"rewardfactionoverride5,omitempty"`
	TimeAllowed int `json:"timeallowed,omitempty"`
	AllowableRaces int16 `json:"allowableraces,omitempty"`
	LogTitle string `json:"logtitle,omitempty"`
	LogDescription string `json:"logdescription,omitempty"`
	QuestDescription string `json:"questdescription,omitempty"`
	AreaDescription string `json:"areadescription,omitempty"`
	QuestCompletionLog string `json:"questcompletionlog,omitempty"`
	RequiredNpcOrGo1 int `json:"requirednpcorgo1,omitempty"`
	RequiredNpcOrGo2 int `json:"requirednpcorgo2,omitempty"`
	RequiredNpcOrGo3 int `json:"requirednpcorgo3,omitempty"`
	RequiredNpcOrGo4 int `json:"requirednpcorgo4,omitempty"`
	RequiredNpcOrGoCount1 int16 `json:"requirednpcorgocount1,omitempty"`
	RequiredNpcOrGoCount2 int16 `json:"requirednpcorgocount2,omitempty"`
	RequiredNpcOrGoCount3 int16 `json:"requirednpcorgocount3,omitempty"`
	RequiredNpcOrGoCount4 int16 `json:"requirednpcorgocount4,omitempty"`
	RequiredItemId1 int `json:"requireditemid1,omitempty"`
	RequiredItemId2 int `json:"requireditemid2,omitempty"`
	RequiredItemId3 int `json:"requireditemid3,omitempty"`
	RequiredItemId4 int `json:"requireditemid4,omitempty"`
	RequiredItemId5 int `json:"requireditemid5,omitempty"`
	RequiredItemId6 int `json:"requireditemid6,omitempty"`
	RequiredItemCount1 int16 `json:"requireditemcount1,omitempty"`
	RequiredItemCount2 int16 `json:"requireditemcount2,omitempty"`
	RequiredItemCount3 int16 `json:"requireditemcount3,omitempty"`
	RequiredItemCount4 int16 `json:"requireditemcount4,omitempty"`
	RequiredItemCount5 int16 `json:"requireditemcount5,omitempty"`
	RequiredItemCount6 int16 `json:"requireditemcount6,omitempty"`
	Unknown0 int8 `json:"unknown0,omitempty"`
	ObjectiveText1 string `json:"objectivetext1,omitempty"`
	ObjectiveText2 string `json:"objectivetext2,omitempty"`
	ObjectiveText3 string `json:"objectivetext3,omitempty"`
	ObjectiveText4 string `json:"objectivetext4,omitempty"`
	VerifiedBuild int `json:"verifiedbuild,omitempty"`
}

type QuestTemplateSlice []QuestTemplate

type DBQuestTemplate struct {
	bun.BaseModel `bun:"table:quest_template,alias:quest_template"`
	ID int `bun:"ID"`
	QuestType int8 `bun:"QuestType"`
	QuestLevel int16 `bun:"QuestLevel"`
	MinLevel int8 `bun:"MinLevel"`
	QuestSortID int16 `bun:"QuestSortID"`
	QuestInfoID int16 `bun:"QuestInfoID"`
	SuggestedGroupNum int8 `bun:"SuggestedGroupNum"`
	RequiredFactionId1 int16 `bun:"RequiredFactionId1"`
	RequiredFactionId2 int16 `bun:"RequiredFactionId2"`
	RequiredFactionValue1 int `bun:"RequiredFactionValue1"`
	RequiredFactionValue2 int `bun:"RequiredFactionValue2"`
	RewardNextQuest int `bun:"RewardNextQuest"`
	RewardXPDifficulty int8 `bun:"RewardXPDifficulty"`
	RewardMoney int `bun:"RewardMoney"`
	RewardBonusMoney int `bun:"RewardBonusMoney"`
	RewardDisplaySpell int `bun:"RewardDisplaySpell"`
	RewardSpell int `bun:"RewardSpell"`
	RewardHonor int `bun:"RewardHonor"`
	RewardKillHonor float64 `bun:"RewardKillHonor"`
	StartItem int `bun:"StartItem"`
	Flags int `bun:"Flags"`
	RequiredPlayerKills int8 `bun:"RequiredPlayerKills"`
	RewardItem1 int `bun:"RewardItem1"`
	RewardAmount1 int16 `bun:"RewardAmount1"`
	RewardItem2 int `bun:"RewardItem2"`
	RewardAmount2 int16 `bun:"RewardAmount2"`
	RewardItem3 int `bun:"RewardItem3"`
	RewardAmount3 int16 `bun:"RewardAmount3"`
	RewardItem4 int `bun:"RewardItem4"`
	RewardAmount4 int16 `bun:"RewardAmount4"`
	ItemDrop1 int `bun:"ItemDrop1"`
	ItemDropQuantity1 int16 `bun:"ItemDropQuantity1"`
	ItemDrop2 int `bun:"ItemDrop2"`
	ItemDropQuantity2 int16 `bun:"ItemDropQuantity2"`
	ItemDrop3 int `bun:"ItemDrop3"`
	ItemDropQuantity3 int16 `bun:"ItemDropQuantity3"`
	ItemDrop4 int `bun:"ItemDrop4"`
	ItemDropQuantity4 int16 `bun:"ItemDropQuantity4"`
	RewardChoiceItemID1 int `bun:"RewardChoiceItemID1"`
	RewardChoiceItemQuantity1 int16 `bun:"RewardChoiceItemQuantity1"`
	RewardChoiceItemID2 int `bun:"RewardChoiceItemID2"`
	RewardChoiceItemQuantity2 int16 `bun:"RewardChoiceItemQuantity2"`
	RewardChoiceItemID3 int `bun:"RewardChoiceItemID3"`
	RewardChoiceItemQuantity3 int16 `bun:"RewardChoiceItemQuantity3"`
	RewardChoiceItemID4 int `bun:"RewardChoiceItemID4"`
	RewardChoiceItemQuantity4 int16 `bun:"RewardChoiceItemQuantity4"`
	RewardChoiceItemID5 int `bun:"RewardChoiceItemID5"`
	RewardChoiceItemQuantity5 int16 `bun:"RewardChoiceItemQuantity5"`
	RewardChoiceItemID6 int `bun:"RewardChoiceItemID6"`
	RewardChoiceItemQuantity6 int16 `bun:"RewardChoiceItemQuantity6"`
	POIContinent int16 `bun:"POIContinent"`
	POIx float64 `bun:"POIx"`
	POIy float64 `bun:"POIy"`
	POIPriority int `bun:"POIPriority"`
	RewardTitle int8 `bun:"RewardTitle"`
	RewardTalents int8 `bun:"RewardTalents"`
	RewardArenaPoints int16 `bun:"RewardArenaPoints"`
	RewardFactionID1 int16 `bun:"RewardFactionID1"`
	RewardFactionValue1 int `bun:"RewardFactionValue1"`
	RewardFactionOverride1 int `bun:"RewardFactionOverride1"`
	RewardFactionID2 int16 `bun:"RewardFactionID2"`
	RewardFactionValue2 int `bun:"RewardFactionValue2"`
	RewardFactionOverride2 int `bun:"RewardFactionOverride2"`
	RewardFactionID3 int16 `bun:"RewardFactionID3"`
	RewardFactionValue3 int `bun:"RewardFactionValue3"`
	RewardFactionOverride3 int `bun:"RewardFactionOverride3"`
	RewardFactionID4 int16 `bun:"RewardFactionID4"`
	RewardFactionValue4 int `bun:"RewardFactionValue4"`
	RewardFactionOverride4 int `bun:"RewardFactionOverride4"`
	RewardFactionID5 int16 `bun:"RewardFactionID5"`
	RewardFactionValue5 int `bun:"RewardFactionValue5"`
	RewardFactionOverride5 int `bun:"RewardFactionOverride5"`
	TimeAllowed int `bun:"TimeAllowed"`
	AllowableRaces int16 `bun:"AllowableRaces"`
	LogTitle string `bun:"LogTitle"`
	LogDescription string `bun:"LogDescription"`
	QuestDescription string `bun:"QuestDescription"`
	AreaDescription string `bun:"AreaDescription"`
	QuestCompletionLog string `bun:"QuestCompletionLog"`
	RequiredNpcOrGo1 int `bun:"RequiredNpcOrGo1"`
	RequiredNpcOrGo2 int `bun:"RequiredNpcOrGo2"`
	RequiredNpcOrGo3 int `bun:"RequiredNpcOrGo3"`
	RequiredNpcOrGo4 int `bun:"RequiredNpcOrGo4"`
	RequiredNpcOrGoCount1 int16 `bun:"RequiredNpcOrGoCount1"`
	RequiredNpcOrGoCount2 int16 `bun:"RequiredNpcOrGoCount2"`
	RequiredNpcOrGoCount3 int16 `bun:"RequiredNpcOrGoCount3"`
	RequiredNpcOrGoCount4 int16 `bun:"RequiredNpcOrGoCount4"`
	RequiredItemId1 int `bun:"RequiredItemId1"`
	RequiredItemId2 int `bun:"RequiredItemId2"`
	RequiredItemId3 int `bun:"RequiredItemId3"`
	RequiredItemId4 int `bun:"RequiredItemId4"`
	RequiredItemId5 int `bun:"RequiredItemId5"`
	RequiredItemId6 int `bun:"RequiredItemId6"`
	RequiredItemCount1 int16 `bun:"RequiredItemCount1"`
	RequiredItemCount2 int16 `bun:"RequiredItemCount2"`
	RequiredItemCount3 int16 `bun:"RequiredItemCount3"`
	RequiredItemCount4 int16 `bun:"RequiredItemCount4"`
	RequiredItemCount5 int16 `bun:"RequiredItemCount5"`
	RequiredItemCount6 int16 `bun:"RequiredItemCount6"`
	Unknown0 int8 `bun:"Unknown0"`
	ObjectiveText1 string `bun:"ObjectiveText1"`
	ObjectiveText2 string `bun:"ObjectiveText2"`
	ObjectiveText3 string `bun:"ObjectiveText3"`
	ObjectiveText4 string `bun:"ObjectiveText4"`
	VerifiedBuild int `bun:"VerifiedBuild"`
}

type DBQuestTemplateSlice []DBQuestTemplate

func (entry *QuestTemplate) ToDB() *DBQuestTemplate {
	if entry == nil {
		return nil
	}
	return &DBQuestTemplate{
		ID: entry.ID,
		QuestType: entry.QuestType,
		QuestLevel: entry.QuestLevel,
		MinLevel: entry.MinLevel,
		QuestSortID: entry.QuestSortID,
		QuestInfoID: entry.QuestInfoID,
		SuggestedGroupNum: entry.SuggestedGroupNum,
		RequiredFactionId1: entry.RequiredFactionId1,
		RequiredFactionId2: entry.RequiredFactionId2,
		RequiredFactionValue1: entry.RequiredFactionValue1,
		RequiredFactionValue2: entry.RequiredFactionValue2,
		RewardNextQuest: entry.RewardNextQuest,
		RewardXPDifficulty: entry.RewardXPDifficulty,
		RewardMoney: entry.RewardMoney,
		RewardBonusMoney: entry.RewardBonusMoney,
		RewardDisplaySpell: entry.RewardDisplaySpell,
		RewardSpell: entry.RewardSpell,
		RewardHonor: entry.RewardHonor,
		RewardKillHonor: entry.RewardKillHonor,
		StartItem: entry.StartItem,
		Flags: entry.Flags,
		RequiredPlayerKills: entry.RequiredPlayerKills,
		RewardItem1: entry.RewardItem1,
		RewardAmount1: entry.RewardAmount1,
		RewardItem2: entry.RewardItem2,
		RewardAmount2: entry.RewardAmount2,
		RewardItem3: entry.RewardItem3,
		RewardAmount3: entry.RewardAmount3,
		RewardItem4: entry.RewardItem4,
		RewardAmount4: entry.RewardAmount4,
		ItemDrop1: entry.ItemDrop1,
		ItemDropQuantity1: entry.ItemDropQuantity1,
		ItemDrop2: entry.ItemDrop2,
		ItemDropQuantity2: entry.ItemDropQuantity2,
		ItemDrop3: entry.ItemDrop3,
		ItemDropQuantity3: entry.ItemDropQuantity3,
		ItemDrop4: entry.ItemDrop4,
		ItemDropQuantity4: entry.ItemDropQuantity4,
		RewardChoiceItemID1: entry.RewardChoiceItemID1,
		RewardChoiceItemQuantity1: entry.RewardChoiceItemQuantity1,
		RewardChoiceItemID2: entry.RewardChoiceItemID2,
		RewardChoiceItemQuantity2: entry.RewardChoiceItemQuantity2,
		RewardChoiceItemID3: entry.RewardChoiceItemID3,
		RewardChoiceItemQuantity3: entry.RewardChoiceItemQuantity3,
		RewardChoiceItemID4: entry.RewardChoiceItemID4,
		RewardChoiceItemQuantity4: entry.RewardChoiceItemQuantity4,
		RewardChoiceItemID5: entry.RewardChoiceItemID5,
		RewardChoiceItemQuantity5: entry.RewardChoiceItemQuantity5,
		RewardChoiceItemID6: entry.RewardChoiceItemID6,
		RewardChoiceItemQuantity6: entry.RewardChoiceItemQuantity6,
		POIContinent: entry.POIContinent,
		POIx: entry.POIx,
		POIy: entry.POIy,
		POIPriority: entry.POIPriority,
		RewardTitle: entry.RewardTitle,
		RewardTalents: entry.RewardTalents,
		RewardArenaPoints: entry.RewardArenaPoints,
		RewardFactionID1: entry.RewardFactionID1,
		RewardFactionValue1: entry.RewardFactionValue1,
		RewardFactionOverride1: entry.RewardFactionOverride1,
		RewardFactionID2: entry.RewardFactionID2,
		RewardFactionValue2: entry.RewardFactionValue2,
		RewardFactionOverride2: entry.RewardFactionOverride2,
		RewardFactionID3: entry.RewardFactionID3,
		RewardFactionValue3: entry.RewardFactionValue3,
		RewardFactionOverride3: entry.RewardFactionOverride3,
		RewardFactionID4: entry.RewardFactionID4,
		RewardFactionValue4: entry.RewardFactionValue4,
		RewardFactionOverride4: entry.RewardFactionOverride4,
		RewardFactionID5: entry.RewardFactionID5,
		RewardFactionValue5: entry.RewardFactionValue5,
		RewardFactionOverride5: entry.RewardFactionOverride5,
		TimeAllowed: entry.TimeAllowed,
		AllowableRaces: entry.AllowableRaces,
		LogTitle: entry.LogTitle,
		LogDescription: entry.LogDescription,
		QuestDescription: entry.QuestDescription,
		AreaDescription: entry.AreaDescription,
		QuestCompletionLog: entry.QuestCompletionLog,
		RequiredNpcOrGo1: entry.RequiredNpcOrGo1,
		RequiredNpcOrGo2: entry.RequiredNpcOrGo2,
		RequiredNpcOrGo3: entry.RequiredNpcOrGo3,
		RequiredNpcOrGo4: entry.RequiredNpcOrGo4,
		RequiredNpcOrGoCount1: entry.RequiredNpcOrGoCount1,
		RequiredNpcOrGoCount2: entry.RequiredNpcOrGoCount2,
		RequiredNpcOrGoCount3: entry.RequiredNpcOrGoCount3,
		RequiredNpcOrGoCount4: entry.RequiredNpcOrGoCount4,
		RequiredItemId1: entry.RequiredItemId1,
		RequiredItemId2: entry.RequiredItemId2,
		RequiredItemId3: entry.RequiredItemId3,
		RequiredItemId4: entry.RequiredItemId4,
		RequiredItemId5: entry.RequiredItemId5,
		RequiredItemId6: entry.RequiredItemId6,
		RequiredItemCount1: entry.RequiredItemCount1,
		RequiredItemCount2: entry.RequiredItemCount2,
		RequiredItemCount3: entry.RequiredItemCount3,
		RequiredItemCount4: entry.RequiredItemCount4,
		RequiredItemCount5: entry.RequiredItemCount5,
		RequiredItemCount6: entry.RequiredItemCount6,
		Unknown0: entry.Unknown0,
		ObjectiveText1: entry.ObjectiveText1,
		ObjectiveText2: entry.ObjectiveText2,
		ObjectiveText3: entry.ObjectiveText3,
		ObjectiveText4: entry.ObjectiveText4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (entry *DBQuestTemplate) ToWeb() *QuestTemplate {
	if entry == nil {
		return nil
	}
	return &QuestTemplate{
		ID: entry.ID,
		QuestType: entry.QuestType,
		QuestLevel: entry.QuestLevel,
		MinLevel: entry.MinLevel,
		QuestSortID: entry.QuestSortID,
		QuestInfoID: entry.QuestInfoID,
		SuggestedGroupNum: entry.SuggestedGroupNum,
		RequiredFactionId1: entry.RequiredFactionId1,
		RequiredFactionId2: entry.RequiredFactionId2,
		RequiredFactionValue1: entry.RequiredFactionValue1,
		RequiredFactionValue2: entry.RequiredFactionValue2,
		RewardNextQuest: entry.RewardNextQuest,
		RewardXPDifficulty: entry.RewardXPDifficulty,
		RewardMoney: entry.RewardMoney,
		RewardBonusMoney: entry.RewardBonusMoney,
		RewardDisplaySpell: entry.RewardDisplaySpell,
		RewardSpell: entry.RewardSpell,
		RewardHonor: entry.RewardHonor,
		RewardKillHonor: entry.RewardKillHonor,
		StartItem: entry.StartItem,
		Flags: entry.Flags,
		RequiredPlayerKills: entry.RequiredPlayerKills,
		RewardItem1: entry.RewardItem1,
		RewardAmount1: entry.RewardAmount1,
		RewardItem2: entry.RewardItem2,
		RewardAmount2: entry.RewardAmount2,
		RewardItem3: entry.RewardItem3,
		RewardAmount3: entry.RewardAmount3,
		RewardItem4: entry.RewardItem4,
		RewardAmount4: entry.RewardAmount4,
		ItemDrop1: entry.ItemDrop1,
		ItemDropQuantity1: entry.ItemDropQuantity1,
		ItemDrop2: entry.ItemDrop2,
		ItemDropQuantity2: entry.ItemDropQuantity2,
		ItemDrop3: entry.ItemDrop3,
		ItemDropQuantity3: entry.ItemDropQuantity3,
		ItemDrop4: entry.ItemDrop4,
		ItemDropQuantity4: entry.ItemDropQuantity4,
		RewardChoiceItemID1: entry.RewardChoiceItemID1,
		RewardChoiceItemQuantity1: entry.RewardChoiceItemQuantity1,
		RewardChoiceItemID2: entry.RewardChoiceItemID2,
		RewardChoiceItemQuantity2: entry.RewardChoiceItemQuantity2,
		RewardChoiceItemID3: entry.RewardChoiceItemID3,
		RewardChoiceItemQuantity3: entry.RewardChoiceItemQuantity3,
		RewardChoiceItemID4: entry.RewardChoiceItemID4,
		RewardChoiceItemQuantity4: entry.RewardChoiceItemQuantity4,
		RewardChoiceItemID5: entry.RewardChoiceItemID5,
		RewardChoiceItemQuantity5: entry.RewardChoiceItemQuantity5,
		RewardChoiceItemID6: entry.RewardChoiceItemID6,
		RewardChoiceItemQuantity6: entry.RewardChoiceItemQuantity6,
		POIContinent: entry.POIContinent,
		POIx: entry.POIx,
		POIy: entry.POIy,
		POIPriority: entry.POIPriority,
		RewardTitle: entry.RewardTitle,
		RewardTalents: entry.RewardTalents,
		RewardArenaPoints: entry.RewardArenaPoints,
		RewardFactionID1: entry.RewardFactionID1,
		RewardFactionValue1: entry.RewardFactionValue1,
		RewardFactionOverride1: entry.RewardFactionOverride1,
		RewardFactionID2: entry.RewardFactionID2,
		RewardFactionValue2: entry.RewardFactionValue2,
		RewardFactionOverride2: entry.RewardFactionOverride2,
		RewardFactionID3: entry.RewardFactionID3,
		RewardFactionValue3: entry.RewardFactionValue3,
		RewardFactionOverride3: entry.RewardFactionOverride3,
		RewardFactionID4: entry.RewardFactionID4,
		RewardFactionValue4: entry.RewardFactionValue4,
		RewardFactionOverride4: entry.RewardFactionOverride4,
		RewardFactionID5: entry.RewardFactionID5,
		RewardFactionValue5: entry.RewardFactionValue5,
		RewardFactionOverride5: entry.RewardFactionOverride5,
		TimeAllowed: entry.TimeAllowed,
		AllowableRaces: entry.AllowableRaces,
		LogTitle: entry.LogTitle,
		LogDescription: entry.LogDescription,
		QuestDescription: entry.QuestDescription,
		AreaDescription: entry.AreaDescription,
		QuestCompletionLog: entry.QuestCompletionLog,
		RequiredNpcOrGo1: entry.RequiredNpcOrGo1,
		RequiredNpcOrGo2: entry.RequiredNpcOrGo2,
		RequiredNpcOrGo3: entry.RequiredNpcOrGo3,
		RequiredNpcOrGo4: entry.RequiredNpcOrGo4,
		RequiredNpcOrGoCount1: entry.RequiredNpcOrGoCount1,
		RequiredNpcOrGoCount2: entry.RequiredNpcOrGoCount2,
		RequiredNpcOrGoCount3: entry.RequiredNpcOrGoCount3,
		RequiredNpcOrGoCount4: entry.RequiredNpcOrGoCount4,
		RequiredItemId1: entry.RequiredItemId1,
		RequiredItemId2: entry.RequiredItemId2,
		RequiredItemId3: entry.RequiredItemId3,
		RequiredItemId4: entry.RequiredItemId4,
		RequiredItemId5: entry.RequiredItemId5,
		RequiredItemId6: entry.RequiredItemId6,
		RequiredItemCount1: entry.RequiredItemCount1,
		RequiredItemCount2: entry.RequiredItemCount2,
		RequiredItemCount3: entry.RequiredItemCount3,
		RequiredItemCount4: entry.RequiredItemCount4,
		RequiredItemCount5: entry.RequiredItemCount5,
		RequiredItemCount6: entry.RequiredItemCount6,
		Unknown0: entry.Unknown0,
		ObjectiveText1: entry.ObjectiveText1,
		ObjectiveText2: entry.ObjectiveText2,
		ObjectiveText3: entry.ObjectiveText3,
		ObjectiveText4: entry.ObjectiveText4,
		VerifiedBuild: entry.VerifiedBuild,
	}
}

func (data QuestTemplateSlice) ToDB() DBQuestTemplateSlice {
	result := make(DBQuestTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBQuestTemplateSlice) ToWeb() QuestTemplateSlice {
	result := make(QuestTemplateSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
