// Copyright Epic Games, Inc. All Rights Reserved.

package extData

// MuSpiteExtData represents ext data for MuSprite type monsters
type MuSpiteExtData struct {
	GroupID int `json:"groupId"`
}

// GetGroupID returns the group id
func (m *MuSpiteExtData) GetGroupID() int {
	return m.GroupID
}

// SetGroupID sets the group id
func (m *MuSpiteExtData) SetGroupID(groupID int) {
	m.GroupID = groupID
}

// MuSpiteExtDataCreator creates MuSpiteExtData instances
type MuSpiteExtDataCreator struct {
	*BaseExtDataCreator
}

// NewInstance creates a new MuSpiteExtData instance
func (c *MuSpiteExtDataCreator) NewInstance() interface{} {
	return &MuSpiteExtData{}
}
