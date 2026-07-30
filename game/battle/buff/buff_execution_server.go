package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gametimer"
	"gameSrv/pkg/actors"
	"time"
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
) actors.IBuff {
	newBuff := s.FYBuffExecution.AddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
	if newBuff == nil {
		return nil
	}

	fyBuff, ok := newBuff.(*FYBuff)
	if !ok {
		return newBuff
	}
	s.startBuffServerTick(fyBuff)
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
	}

	// Generate unique timer key using buff UID
	timerKey := fyBuff.UID
	fyBuff.TimerKey = timerKey

	// Schedule the first tick after periodicTime milliseconds
	delay := time.Duration(periodicTime) * time.Millisecond
	gametimer.AddTask(timerKey, delay, func() {
		s.tickFromServer(fyBuff)
	})
}

// tickFromServer handles server-side buff tick
func (s *FYBuffExecutionServer) tickFromServer(fyBuff *FYBuff) {
	if fyBuff.Prop == nil {
		return
	}

	isTimeOut := fyBuff.ServerTick()
	if isTimeOut {
		s.StopAndRemoveBuff(fyBuff, int(cfg.BuffEndTypeEnum_BUFF_END_TIMEUP))
		return
	}

	// Re-schedule next tick if buff is still active
	periodicTime := fyBuff.PeriodicTime
	if periodicTime == 0 {
		periodicTime = int(fyBuff.Prop.TotalTime)
	}

	delay := time.Duration(periodicTime) * time.Millisecond
	gametimer.AddTask(fyBuff.TimerKey, delay, func() {
		s.tickFromServer(fyBuff)
	})
}

// onBuffAdded is called when buff is added (server-side)
func (s *FYBuffExecutionServer) onBuffAdded(buff *FYBuff) {
	// Notify holder actor about buff update
	if buff != nil && buff.HolderActor != nil && buff.HolderActor.ActorBattleModule != nil {
		// ActorBattleModule will handle buff update notifications
		_ = buff.HolderActor.ActorBattleModule
	}
}

// onBuffTriggerEventFinished is called when buff trigger event finishes (server-side)
func (s *FYBuffExecutionServer) onBuffTriggerEventFinished() {
	// Notify owner actor about property change completion
	if s != nil && s.owner != nil && s.owner.ActorBattleModule != nil {
		// ActorBattleModule will handle property change completion notifications
		s.owner.ActorBattleModule.OnAvatarPropsChangFinish()
	}
}

// StopAndRemoveBuff stops and removes a buff (server-side implementation)
func (s *FYBuffExecutionServer) StopAndRemoveBuff(removedBuff *FYBuff, reason int) bool {
	// Cancel timer task if exists
	if removedBuff.TimerKey != 0 {
		gametimer.RemoveTask(removedBuff.TimerKey)
		removedBuff.TimerKey = 0
	}

	s.FYBuffExecution.StopAndRemoveBuff(removedBuff, reason)

	// Notify holder actor about buff removal
	if removedBuff != nil && removedBuff.HolderActor != nil && removedBuff.HolderActor.ActorBattleModule != nil {
		// ActorBattleModule will handle buff removal notifications
		_ = removedBuff.HolderActor.ActorBattleModule
	}

	return true
}

// onPlayerLeaveScene is called when player leaves scene (server-side)
func (s *FYBuffExecutionServer) onPlayerLeaveScene() {
	// Remove all buffs when player leaves scene
	// This will be implemented when scene management system is available
	// For now, buffs are cleaned up through other means
}

// synBuffToClient syncs buff to client (server-side)
func (s *FYBuffExecutionServer) synBuffToClient() {
	// Send buff data to client through network sync system
	// This will be implemented when network sync system is available
}
