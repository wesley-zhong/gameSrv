package excuteEvents

// Copyright Epic Games, Inc. All Rights Reserved.

// InteractExecuteEventHandler 交互后处理事件处理器接口
type InteractExecuteEventHandler interface {
	// GetType 获取处理器对应的事件类型
	GetType() int

	// Execute 执行后处理事件
	// interactMan: 交互者
	// targetActor: 被交互的目标Actor
	// eventConfig: 事件配置
	Execute(interactMan Actor, targetActor Actor, eventConfig InteractExecuteEventsConfig) error
}
