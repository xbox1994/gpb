package request

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
