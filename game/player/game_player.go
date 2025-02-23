package player

import (
	"gameSrv/game/modules"
	"gameSrv/pkg/client"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

type GamePlayer struct {
	Pid int64 //player id
	Sid int64 // session id
	modules.ModuleContainer
	innerClientCtx *client.ConnInnerClientContext
}

func NewGamePlayer(playerId int64, innerClientCtx *client.ConnInnerClientContext) *GamePlayer {
	player := &GamePlayer{}
	player.Init(playerId, innerClientCtx)
	return player
}

func (player *GamePlayer) Init(playerId int64, innerClientCtx *client.ConnInnerClientContext) {
	player.innerClientCtx = innerClientCtx
	player.Pid = playerId
	player.InitModules()
}
func (player *GamePlayer) SendMsg(msgId protoGen.ProtoCode, msg proto.Message) {
	player.innerClientCtx.SendMsg(msgId, player.Pid, msg)
}
