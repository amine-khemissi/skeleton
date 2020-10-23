package write

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	ClientID string `json:"clientID"`
}
