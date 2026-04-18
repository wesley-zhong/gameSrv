// Copyright Epic Games, Inc. All Rights Reserved.

package extData

import "reflect"

// IExtDataCreator is an abstract creator for ext data instances
type IExtDataCreator interface {
	// SubType returns the monster sub type
	SubType() int
	// ObjClass returns the reflect type of the ext data
	ObjClass() reflect.Type
	// NewInstance creates a new instance of the ext data
	NewInstance() interface{}
}

// BaseExtDataCreator provides a base implementation for IExtDataCreator
type BaseExtDataCreator struct {
	subType  int
	objClass reflect.Type
}

// NewBaseExtDataCreator creates a new BaseExtDataCreator
func NewBaseExtDataCreator(subType int, objClass reflect.Type) *BaseExtDataCreator {
	return &BaseExtDataCreator{
		subType:  subType,
		objClass: objClass,
	}
}

// SubType returns the monster sub type
func (c *BaseExtDataCreator) SubType() int {
	return c.subType
}

// ObjClass returns the reflect type of the ext data
func (c *BaseExtDataCreator) ObjClass() reflect.Type {
	return c.objClass
}
