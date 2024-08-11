package api

import "github.com/redxiiikk/wails-app-template/backend/utils"

type EchoApi struct{}

func NewEchoApi() *EchoApi {
	utils.Logger.Info("[API] create new echo api instance")
	return &EchoApi{}
}

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string `json:"message"`
}

func (echoApi *EchoApi) Echo(request EchoRequest) EchoResponse {
	return EchoResponse(request)
}
