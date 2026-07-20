package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// FYBuffExecutionServer handles server-side buff execution
type FYBuffExecutionServer struct {
	*FYBuffExecution
}

// NewFYBuffExecutionServer creates a new server-side buff execution
func NewFYBuffExecutionServer(owner *actors.Creature) *FYBuffExecutionServer {
	return &FYBuffExecutionServer{
		FYBuffExecution: NewFYBuffExecution(owner),
	}
}

// AddBuff adds a buff (server-side implementation)
func (s *FYBuffExecutionServer) AddBuff(
	templateId int,
	casterActor, holder *actors.Creature,
	uid int64,
	exParam, layer int,
	life int64,
	bSystem bool,
) *FYBuff {
	newBuff := s.FYBuffExecution.AddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
	if newBuff == nil {
		return nil
	}

	s.startBuffServerTick(newBuff)
	return newBuff
}

// startBuffServerTick starts server-side buff ticking
func (s *FYBuffExecutionServer) startBuffServerTick(fyBuff *FYBuff) {
	if fyBuff.Prop == nil || fyBuff.Prop.TotalTime <= 0 {
		return
	}

	periodicTime := fyBuff.PeriodicTime

	// Non-periodic buff - set timeout
	if periodicTime == 0 {
		periodicTime = int(fyBuff.Prop.TotalTime)
	} else {
		// Periodic buff
		_ = int(fyBuff.Prop.TotalTime) / periodicTime
	}

	// Execute first tick
	s.tickFromServer(fyBuff)

	// TODO: Setup timer task when timer system is implemented
	// fyBuff.TimerTask = ScheduleService.INSTANCE.newTimerTask(
	//     s.tickFromServer,
	//     fyBuff,
	//     nil,
	//     periodicTime,
	//     exeCount,
	// )
	// ScheduleService.INSTANCE.executeTimerTask(fyBuff.TimerTask, periodicTime)
}

// tickFromServer handles server-side buff tick
func (s *FYBuffExecutionServer) tickFromServer(fyBuff *FYBuff) {
	if fyBuff.Prop == nil {
		return
	}

	isTimeOut := fyBuff.ServerTick()
	if isTimeOut {
		s.StopAndRemoveBuff(fyBuff, int(cfg.BuffEndTypeEnum_BUFF_END_TIMEUP))
	}
}

// onBuffAdded is called when buff is added (server-side)
func (s *FYBuffExecutionServer) onBuffAdded(buff *FYBuff) {
	// TODO: Notify ActorBattleModule when implemented
	// s.owner.ActorBattleModule.ActorBuffModule.OnBuffUpdate(buff)
}

// onBuffTriggerEventFinished is called when buff trigger event finishes (server-side)
func (s *FYBuffExecutionServer) onBuffTriggerEventFinished() {
	// TODO: Notify ActorBattleModule when implemented
	// s.owner.ActorBattleModule.onAvatarPropsChangFinish()
}

// StopAndRemoveBuff stops and removes a buff (server-side implementation)
func (s *FYBuffExecutionServer) StopAndRemoveBuff(removedBuff *FYBuff, reason int) bool {
	s.FYBuffExecution.StopAndRemoveBuff(removedBuff, reason)
	// TODO: Notify ActorBattleModule when implemented
	// s.owner.ActorBattleModule.ActorBuffModule.OnBuffRemoved(removedBuff, reason)
	return true
}

// onPlayerLeaveScene is called when player leaves scene (server-side)
func (s *FYBuffExecutionServer) onPlayerLeaveScene() {
	// TODO: Implement server-specific leave scene logic
}

// synBuffToClient syncs buff to client (server-side)
func (s *FYBuffExecutionServer) synBuffToClient() {
	// TODO: Implement server-specific sync logic
}
