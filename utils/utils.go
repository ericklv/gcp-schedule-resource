package utils

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func S200(msg string) Response {
	return Response{200, msg}
}

func S4xx(msg string) Response {
	return Response{400, msg}
}

func S5xx(msg string) Response {
	return Response{500, msg}
}
