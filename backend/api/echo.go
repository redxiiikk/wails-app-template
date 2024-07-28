package api

type EchoApi struct{}

func NewEchoApi() *EchoApi {
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
