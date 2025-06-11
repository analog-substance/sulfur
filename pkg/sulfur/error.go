package sulfur

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Identity struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"identity"`
	} `json:"data"`
}
