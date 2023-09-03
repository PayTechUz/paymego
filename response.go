package paymego

type Receipt struct {
	ID string `json:"_id"`
}

type Card struct {
	Number     string `json:"number"`
	Expire     string `json:"expire"`
	Token      string `json:"token"`
	Recurrent  bool   `json:"recurrent"`
	Verify     bool   `json:"verify"`
	Type       string `json:"type"`
	NumberHash string `json:"number_hash"`
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
		Card    Card    `json:"card"`
	} `json:"result"`
	Error Error `json:"error"`
}
