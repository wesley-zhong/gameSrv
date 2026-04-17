// Copyright Epic Games, Inc. All Rights Reserved.

package scene

import (
	"gameSrv/pkg/math"
)

// CheckRegionType 区域检查类型
type CheckRegionType int

const (
	CheckRegionTypeNone CheckRegionType = iota
	CheckRegionTypeMove  CheckRegionType = iota
	CheckRegionTypeBorn CheckRegionType = iota
	CheckRegionTypeLeave CheckRegionType = iota
)

// ScenePlayerPreSlotInfo 场景玩家预进入信息
type ScenePlayerPreSlotInfo struct {
	UID          int64
	PreEnterTime int64
	Nickname     string
	Transform   *math.Vector3
}

// NewScenePlayerPreSlotInfo creates a new ScenePlayerPreSlotInfo
func NewScenePlayerPreSlotInfo() *ScenePlayerPreSlotInfo {
	return &ScenePlayerPreSlotInfo{}
}

// AbstractScene represents an abstract scene base class
type AbstractScene struct {
	OwnerUID          int64
	SceneUID           string
	SceneCnfId         int32
	BeginTime          int64
	AllowOptions        int32
}

// NewAbstractScene creates a new AbstractScene
func NewAbstractScene(sceneUID string, sceneCnfId int32) *AbstractScene {
	return &AbstractScene{
		SceneUID:     sceneUID,
		SceneCnfId:   sceneCnfId,
		BeginTime:    0,
		AllowOptions:  0,
	}
}

// GetSightModule returns the vision module (abstract method)
func (s *AbstractScene) GetSightModule() interface{} {
	return nil // To be implemented by subclasses
}

// GetSceneType returns scene type (abstract method)
func (s *AbstractScene) GetSceneType() int32 {
	return 0 // To be implemented by subclasses
}

// CanAcceptQuest checks if a quest type can be accepted (abstract method)
func (s *AbstractScene) CanAcceptQuest(questType int32) bool {
	return false // To be implemented by subclasses
}