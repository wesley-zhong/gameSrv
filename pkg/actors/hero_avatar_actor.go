package actors

// HeroAvatarActor is a hero avatar actor
type HeroAvatarActor struct {
	AvatarActor
}

// NewHeroAvatarActor creates a new HeroAvatarActor
func NewHeroAvatarActor() *HeroAvatarActor {
	h := &HeroAvatarActor{}
	return h
}
