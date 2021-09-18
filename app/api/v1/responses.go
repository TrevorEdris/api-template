package v1

type simpleResponse struct {
	Msg string `json:"msg"`
}

type errorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func newSimpleResponse(msg string) simpleResponse {
	return simpleResponse{
		Msg: msg,
	}
}

func newErrorResponse(status, err string) errorResponse {
	return errorResponse{
		Status: status,
		Error:  err,
	}
}
