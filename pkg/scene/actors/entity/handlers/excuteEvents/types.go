package excuteEvents

// Copyright Epic Games, Inc. All Rights Reserved.

// Actor Actor接口定义
type Actor interface {
	GetOwner() Owner
}

// Owner 所有者接口定义
type Owner interface {
	GetUid() uint64
}

// InteractExecuteEventsConfig 交互执行事件配置接口
type InteractExecuteEventsConfig interface {
	GetProperty() InteractExecuteEventsProperty
}

// InteractExecuteEventsProperty 交互执行事件属性接口
type InteractExecuteEventsProperty interface {
	GetExecuteType() int
}
