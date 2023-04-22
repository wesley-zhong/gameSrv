package message

const (
	TypeGateway               = 2
	TypeGame                  = 3
	INNER_PROTO_LOGIN_REQUEST = -5
	INNER_PROTO_ADD_SERVER    = -1
)

func BuildServerUid(serverType, serverId int) int64 {
	return ((int64(serverType)) << 32) | int64(serverId)
}
