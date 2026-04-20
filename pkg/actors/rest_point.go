package actors

import "gameSrv/pkg/scene"

// RestPoint is a rest point
type RestPoint struct {
	SimpleActor
}

func (r *RestPoint) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewRestPoint creates a new RestPoint
func NewRestPoint() *RestPoint {
	r := &RestPoint{}
	r.SimpleActor = *NewSimpleActor()
	return r
}

// GetActorType returns actor type
func (r *RestPoint) GetActorType() int32 {
	return 17 // EActorType_RestPoint
}
