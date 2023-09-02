package paymego

type Receipt struct {
	ID string `json:"_id"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Origin  string `json:"origin"`
}

type PaymeResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  struct {
		Receipt Receipt `json:"receipt"`
	} `json:"result"`
	Error Error `json:"error"`
}
