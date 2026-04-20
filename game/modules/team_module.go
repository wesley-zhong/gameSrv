package modules

type Team struct {
	AvatarIds []int32
}
type TeamDO struct {
	Id        int64
	CurTeamId int64
	Teams     map[int32]*Team
}
type TeamModule struct {
	GameModule[TeamDO]
	CurTeam *Team
}
