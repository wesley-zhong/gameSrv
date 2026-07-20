package buff

import (
	"gameSrv/pkg/actors"
)

// FYBuffExecutionClient handles client-side buff execution
type FYBuffExecutionClient struct {
	*FYBuffExecution
}

// NewFYBuffExecutionClient creates a new client-side buff execution
func NewFYBuffExecutionClient(owner *actors.Creature) *FYBuffExecutionClient {
	return &FYBuffExecutionClient{
		FYBuffExecution: NewFYBuffExecution(owner),
	}
}

// onBuffAdded is called when buff is added (client-side)
func (c *FYBuffExecutionClient) onBuffAdded(buff *FYBuff) {
	// Client-side implementation - empty for now
}

// onBuffTriggerEventFinished is called when buff trigger event finishes (client-side)
func (c *FYBuffExecutionClient) onBuffTriggerEventFinished() {
	// Client-side implementation - empty for now
}

// onPlayerLeaveScene is called when player leaves scene (client-side)
func (c *FYBuffExecutionClient) onPlayerLeaveScene() {
	// Client-side implementation - empty for now
}

// synBuffToClient syncs buff to client (client-side)
func (c *FYBuffExecutionClient) synBuffToClient() {
	// Client-side implementation - empty for now
}
