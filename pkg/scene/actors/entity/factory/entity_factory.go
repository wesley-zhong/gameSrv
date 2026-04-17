package factory

import "gameSrv/pkg/scene"

type IEntityFactory[T any] interface {
	CreateEntityFromConfig(cnfId int32) *T
	CreateEntityFromDO() *T
	GetEntityType() int32
}

type EntityFactory[T any] struct {
}

func (factory *EntityFactory[T]) CreateEntityFromConfig(cnfId int32) *T {
	obj := factory.createEntity()
	if obj != nil {
		factory.initFromConfig(obj, cnfId)
		return obj
	}
	return nil
}

func (factory *EntityFactory[T]) CreateEntityFromDO() *T {
	obj := factory.createEntity()
	if obj != nil {
		factory.initFromDO(obj)
		return obj
	}
	return nil
}

func (factory *EntityFactory[T]) createEntity() *T {
	return nil
}

func (factory *EntityFactory[T]) initFromConfig(obj *T, cnfId int32) {

}

func (factory *EntityFactory[T]) initFromDO(obj *T) {

}

// Global factory instances
var (
	monsterFactory          *MonsterFactory
	gadgetFactory           *GadgetFactory
	summonFactory           *SummonFactory
	chestFactory            *ChestFactory
	buffVolumeFactory       *BuffVolumeFactory
	killZVolumeFactory      *KillZVolumeFactory
	ladderFactory           *LadderFactory
	invisibleWallFactory    *InvisibleWallFactory
	interactiveActorFactory *InteractiveActorFactory
	chestDestructionFactory *ChestDestructionFactory
	commonSwitchFactory     *CommonSwitchFactory
	restPointFactory        *RestPointFactory
	teleporterFactory       *TeleporterFactory
	gameFlowActorFactory    *GameFlowActorFactory
)

// InitFactories initializes all global factory instances
func InitFactories() {
	monsterFactory = NewMonsterFactory()
	gadgetFactory = NewGadgetFactory()
	summonFactory = NewSummonFactory()
	chestFactory = NewChestFactory()
	buffVolumeFactory = NewBuffVolumeFactory()
	killZVolumeFactory = NewKillZVolumeFactory()
	ladderFactory = NewLadderFactory()
	invisibleWallFactory = NewInvisibleWallFactory()
	interactiveActorFactory = NewInteractiveActorFactory()
	chestDestructionFactory = NewChestDestructionFactory()
	commonSwitchFactory = NewCommonSwitchFactory()
	restPointFactory = NewRestPointFactory()
	teleporterFactory = NewTeleporterFactory()
	gameFlowActorFactory = NewGameFlowActorFactory()
}

// CreateEntityFromConf creates entity from configuration using global factory instances
func CreateEntityFromConf(actorType int32, cnfId int32) scene.IEntity {
	switch actorType {
	case 1: // Monster
		return monsterFactory.CreateEntityFromConfig(cnfId)
	case 3: // Gadget
		return gadgetFactory.CreateEntityFromConfig(cnfId)
	case 4: // Summon
		return summonFactory.CreateEntityFromConfig(cnfId)
	case 5: // Chest
		return chestFactory.CreateEntityFromConfig(cnfId)
	case 6: // BuffVolume
		return buffVolumeFactory.CreateEntityFromConfig(cnfId)
	case 7: // KillZVolume
		return killZVolumeFactory.CreateEntityFromConfig(cnfId)
	case 8: // Ladder
		return ladderFactory.CreateEntityFromConfig(cnfId)
	case 9: // InvisibleWall
		return invisibleWallFactory.CreateEntityFromConfig(cnfId)
	case 10: // InteractiveActor
		return interactiveActorFactory.CreateEntityFromConfig(cnfId)
	case 11: // ChestDestruction
		return chestDestructionFactory.CreateEntityFromConfig(cnfId)
	case 12: // CommonSwitch
		return commonSwitchFactory.CreateEntityFromConfig(cnfId)
	case 13: // RestPoint
		return restPointFactory.CreateEntityFromConfig(cnfId)
	case 14: // Teleporter
		return teleporterFactory.CreateEntityFromConfig(cnfId)
	case 15: // GameFlow
		return gameFlowActorFactory.CreateEntityFromConfig(cnfId)
	default:
		return nil
	}
}
