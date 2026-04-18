package actors

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

// AvatarActor is a player-controlled creature
type AvatarActor struct {
	PlayerCreature
	IsPlayerControl bool // Whether player is controlling this avatar
}

func (a *AvatarActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

func (a *AvatarActor) GetOnlyVisionUid() int64 {
	return 0
}

func (a *AvatarActor) GetPhasingId() int64 {
	return a.PhasingID
}

// NewAvatarActor creates a new AvatarActor
func NewAvatarActor(p scene.IGamePlayer) *AvatarActor {
	a := &AvatarActor{
		PlayerCreature: *NewPlayerCreature(p),
	}
	return a
}

// GetActorType returns actor type
func (a *AvatarActor) GetActorType() int32 {
	return 3 // EActorType_Avatar
}

// OnBeforeEnterScene is called before entering scene
func (a *AvatarActor) OnBeforeEnterScene(scn scene.IScene, context interface{}) {
	// TODO: implement
}

// InFrontStage is called when entering front stage
func (a *AvatarActor) InFrontStage() {
	// TODO: implement
}

// HandleDead handles death
func (a *AvatarActor) HandleDead(killerActor *Entity) bool {
	// TODO: implement
	return true
}

// RefreshValidTransform refreshes valid transform
func (a *AvatarActor) RefreshValidTransform(pos, rot *math.Vector3) {
	// TODO: implement
}
