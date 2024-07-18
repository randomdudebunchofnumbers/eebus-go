package main

import (
	"context"
	"fmt"

	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

type UseCaseId string
type UseCaseTypeType string

const (
	UseCaseTypeLPC UseCaseTypeType = "LPC"
)

type UseCaseBuilder func(spineapi.EntityLocalInterface, api.EntityEventCallback) api.UseCaseInterface

func (r *Remote) RegisterUseCase(entityType model.EntityTypeType, usecaseId string, builder UseCaseBuilder) {
	// entityType/uc
	var identifier UseCaseId = UseCaseId(fmt.Sprintf("%s/%s", entityType, usecaseId))

	localInterface := r.service.LocalDevice().EntityForType(entityType)
	uc := builder(localInterface, func(
		ski string,
		device spineapi.DeviceRemoteInterface,
		entity spineapi.EntityRemoteInterface,
		event api.EventType,
	) {
		r.PropagateEvent(identifier, ski, device, entity, event)
	})
	r.service.AddUseCase(uc)

	r.RegisterMethods(usecaseId, uc)
}

func (r *Remote) PropagateEvent(
	id UseCaseId,
	ski string,
	device spineapi.DeviceRemoteInterface,
	entity spineapi.EntityRemoteInterface,
	event api.EventType,
) {
	// convert between EntityRemoteInterface and EntityAddressType
	address := entity.Address()
	r.entityInterfaces[fmt.Sprintf("%s", address)] = entity

	params := make(map[string]interface{}, 2)
	params["device"] = device.Ski()
	params["entity"] = address
	for _, conn := range r.connections {
		conn.Notify(context.Background(), string(event), params)
	}
}
