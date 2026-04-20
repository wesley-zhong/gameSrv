// Copyright Epic Games, Inc. All Rights Reserved.

package scene

import (
	"errors"
)

var (
	// ErrViewMgrNotFound 视野管理器未找到错误
	ErrViewMgrNotFound = errors.New("viewMgr is null")
)

// PlayerViewMgr 玩家视野管理器
type PlayerViewMgr struct {
	scene   *Scene                // 所属的scene
	player  IGamePlayer           // 所属的玩家
	viewMap map[int64]interface{} // 正在查看的entity集合
}

// NewPlayerViewMgr 创建新的PlayerViewMgr
func NewPlayerViewMgr(scene *Scene, plyr IGamePlayer) *PlayerViewMgr {
	return &PlayerViewMgr{
		scene:   scene,
		player:  plyr,
		viewMap: make(map[int64]interface{}),
	}
}

// GetPlayerUid 获取玩家UID
func (p *PlayerViewMgr) GetPlayerUid() int64 {
	return p.player.GetUid()
}

// GetPlayer 获取玩家
func (p *PlayerViewMgr) GetPlayer() IGamePlayer {
	return p.player
}

// GetScene 获取场景
func (p *PlayerViewMgr) GetScene() *Scene {
	return p.scene
}

// ResetPlayerViewMgr 重置玩家视野管理器
func (p *PlayerViewMgr) ResetPlayerViewMgr() {
	p.viewMap = make(map[int64]interface{})
}

// IsContainEntityInView 检查实体是否在视野中
func (p *PlayerViewMgr) IsContainEntityInView(act interface{}) bool {
	// Check if entity implements GetEntityId method
	if entity, ok := act.(interface{ GetEntityId() int64 }); ok {
		entityId := entity.GetEntityId()
		_, exists := p.viewMap[entityId]
		return exists
	}
	return false
}

// AddEntityInView 添加实体到视野
func (p *PlayerViewMgr) AddEntityInView(act interface{}) {
	// Check if entity implements GetEntityId method
	if entity, ok := act.(interface{ GetEntityId() int64 }); ok {
		entityId := entity.GetEntityId()
		p.viewMap[entityId] = act
	}
}

// DelEntityInView 从视野中删除实体
func (p *PlayerViewMgr) DelEntityInView(act interface{}) {
	// Check if entity implements GetEntityId method
	if entity, ok := act.(interface{ GetEntityId() int64 }); ok {
		entityId := entity.GetEntityId()
		delete(p.viewMap, entityId)
	}
}

// GetEntitiesInView 获取视野中的实体
func (p *PlayerViewMgr) GetEntitiesInView() []interface{} {
	result := make([]interface{}, 0, len(p.viewMap))
	for _, act := range p.viewMap {
		result = append(result, act)
	}
	return result
}

// CopyEntitiesInView 复制视野中的实体
func (p *PlayerViewMgr) CopyEntitiesInView() []interface{} {
	return p.GetEntitiesInView()
}

// GetPlayerSetByViewEntity 根据查看的实体获取玩家集合
func (p *PlayerViewMgr) GetPlayerSetByViewEntity(entities []interface{}, excludeUid int64) []IGamePlayer {
	// TODO: implement get player set by view entity logic
	return []IGamePlayer{}
}
