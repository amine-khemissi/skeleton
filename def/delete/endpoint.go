package delete

type Request struct {
	ClientID string `json:"-" path:"clientID"`
}

type Response struct {
}
