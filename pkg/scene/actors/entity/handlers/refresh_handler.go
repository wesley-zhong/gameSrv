// Copyright Epic Games, Inc. All Rights Reserved.

package handlers

// ActorRefreshType Actor刷新类型
type ActorRefreshType int32

const (
	ActorRefreshTypeNone         ActorRefreshType = 0
	ActorRefreshTypeTimer        ActorRefreshType = 1
	ActorRefreshTypeOnEnter      ActorRefreshType = 2
	ActorRefreshTypeOnEnterDaily ActorRefreshType = 3
	ActorRefreshTypeOnExit       ActorRefreshType = 4
	ActorRefreshTypeSpecial      ActorRefreshType = 5
)

// RefreshHandler 刷新规则管理接口
type RefreshHandler interface {
	// IsRefresh 是否刷新
	IsRefresh(destroyTime int64, ruleCfg interface{}) bool

	// GetActorRefreshType Refresh类型
	GetActorRefreshType() ActorRefreshType
}
