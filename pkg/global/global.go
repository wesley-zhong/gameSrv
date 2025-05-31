package global

type GameServerType int32

// ConnInnerClientContext -------------server inner aresTcpClient ---------------
const (
	GATE_WAY GameServerType = 1
	GAME     GameServerType = 2
	ROUTER   GameServerType = 3
	LOGIN    GameServerType = 4
)

var SelfServerType GameServerType

var SelfServiceId string
