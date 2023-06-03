package common

func BuildServerUid(serverType, serverId int) int64 {
	return ((int64(serverType)) << 32) | int64(serverId)
}
