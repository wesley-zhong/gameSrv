package player

import (
	"gameSrv/game/modules"
	"gameSrv/pkg/aresTcpClient"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

type GamePlayer struct {
	modules.ModuleContainer
	Pid            int64 //player id
	Sid            int64 // session id
	innerClientCtx *aresTcpClient.ConnInnerClientContext
}

func NewGamePlayer(playerId int64, innerClientCtx *aresTcpClient.ConnInnerClientContext) *GamePlayer {
	player := &GamePlayer{}
	player.GamePlayer = player
	player.Init(playerId, innerClientCtx)
	return player
}

func (player *GamePlayer) Init(playerId int64, innerClientCtx *aresTcpClient.ConnInnerClientContext) {
	player.innerClientCtx = innerClientCtx
	player.Modules = make([]modules.AresModule, modules.MAX_ITEM_MODULES)
	player.GamePlayer = player
	player.Pid = playerId
	player.InitModules()
}
func (player *GamePlayer) SendMsg(msgId protoGen.ProtoCode, msg proto.Message) {
	player.innerClientCtx.SendMsg(msgId, player.Pid, msg)
}
func (player *GamePlayer) GetPlayerId() int64 {
	return player.Pid
}
