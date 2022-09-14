package client

import (
	"github.com/narawichsaphimarn/handlercontrol/interfaces"
)

type callClassRoomListener struct {
	HttpRequest interfaces.HttpRequest
}

func NewCallClassRoomListener(service interfaces.HttpRequest) *callClassRoomListener {
	return &callClassRoomListener{
		HttpRequest: service,
	}
}

func CallClient()  {
	
}
