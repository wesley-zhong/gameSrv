package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/gameevent"
	"gameSrv/game/modules"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

var questProcess = make(map[int]func(*player.GamePlayer, *modules.Quest, *cfg.QuestContent, event.Event) ProRet)

type ProRet int32

const (
	NONE     ProRet = 0
	CONTINUE ProRet = 1
	FINISH   ProRet = 2
	FAILED   ProRet = 3
)

func finishContentInit() {
	questProcess[cfg.QuestContentType_ROLE_LEVEL_UP] = RoleLvlUpProcess
	questProcess[cfg.QuestContentType_QUEST_FINISHED] = MainQuestFinishProcess

	questProcess[cfg.QuestContentType_KILL_MONSTER] = KillMonsterProcess
	questProcess[cfg.QuestContentType_OBTAIN_ITEM] = ObtainItemProcess
}

func ProcessQuestContentByEvent(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {
	processRet := NONE
	questStepCnf := gamedata.Tables.TbQuestStep.Get(quest.Id)
	if questStepCnf != nil {
		return processRet
	}

	for _, content := range questStepCnf.FinishCond {
		processFuc := questProcess[int(content.Type)]
		if processFuc == nil {
			log.Errorf("finish type =%d not found process function", content.Type)
		}
		processed := processFuc(player, quest, content, ev)
		if processRet == NONE {
			processRet = processed
		}
	}
	quest.Finished = checkQuestFinished(quest, questStepCnf)
	if quest.Finished {
		return FINISH
	}
	return CONTINUE
}

func findOrCreateQuestCndDataByType(quest *modules.Quest, contentType int32) *modules.ContentData {
	for _, questData := range quest.ContentData {
		if questData.ContType == contentType {
			return questData
		}
	}

	newQuestData := &modules.ContentData{
		ContType: contentType,
		CurData:  make([]int64, 0),
	}

	quest.ContentData = append(quest.ContentData, newQuestData)
	return newQuestData
}

func RoleLvlUpProcess(player *player.GamePlayer, quest *modules.Quest, questContent *cfg.QuestContent, ev event.Event) ProRet {
	roleLvlUp := ev.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role lvl up  cur lvl =%d", roleLvlUp.CurLvl)
	questDataByType := findOrCreateQuestCndDataByType(quest, questContent.Type)
	questDataByType.CurData[0] = int64(roleLvlUp.CurLvl)
	if roleLvlUp.CurLvl >= int32(questContent.Param[0]) {
		questDataByType.Finished = true
		return FINISH
	}
	return CONTINUE
}

// 完成主线任务 触发
func MainQuestFinishProcess(player *player.GamePlayer, quest *modules.Quest, questContent *cfg.QuestContent, ev event.Event) ProRet {
	questFinish := ev.(*gameevent.MainQuestFinishEvent)
	log.Infof("MainQuestFinishProcess  main quest id=%d", questFinish.MainQuestId)
	questDataByType := findOrCreateQuestCndDataByType(quest, questContent.Type)
	if questFinish.MainQuestId != questContent.Param[0] {
		return NONE
	}
	questDataByType.CurData[0] = questFinish.MainQuestId
	questDataByType.Finished = true
	return FINISH
}

// 杀死 触发(monster configId [lvl] count
func KillMonsterProcess(player *player.GamePlayer, quest *modules.Quest, questContent *cfg.QuestContent, ev event.Event) ProRet {
	KillMonster := ev.(*gameevent.KillMonsterEvent)
	log.Infof("----- kill monster =%d", KillMonster.MonsterId)

	log.Infof("MainQuestFinishProcess  main quest id=%d", KillMonster.MonsterId)
	questDataByType := findOrCreateQuestCndDataByType(quest, questContent.Type)
	if KillMonster.MonsterId != questContent.Param[0] {
		return NONE
	}
	questDataByType.CurData[0]++
	if questDataByType.CurData[0] < questContent.Param[0] {
		return CONTINUE
	}
	questDataByType.Finished = true
	return FINISH
}

// 获取道具(config id, count) 触发
func ObtainItemProcess(player *player.GamePlayer, quest *modules.Quest, questContent *cfg.QuestContent, ev event.Event) ProRet {
	ObtainItem := ev.(*gameevent.ObtainItemEvent)

	dataChange := NONE
	questDataByType := findOrCreateQuestCndDataByType(quest, questContent.Type)
	// do special check
	targetItemId := questContent.Param[0]
	if targetItemId == 0 || targetItemId == int64(ObtainItem.CnfId) {
		questDataByType.CurData[0]++
		dataChange = CONTINUE
		if questDataByType.CurData[0] >= questContent.Param[1] {
			questDataByType.Finished = true
			dataChange = FINISH
		}
	}
	return dataChange
}

// finish check
func checkQuestFinished(quest *modules.Quest, questStepCnf *cfg.QuestQuestStepCnf) bool {
	ret := true
	for _, cnd := range quest.ContentData {
		ret = ret && cnd.Finished
		// And logic
		if questStepCnf.FinishCondComb == cfg.LogicType_AND {
			if !ret {
				return false
			}
			continue
		}
		// Or logic
		if ret {
			return false
		}
	}
	return true
}
