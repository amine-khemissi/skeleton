package update

type Request struct {
	ClientID string `json:"-" path:"clientID"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type Response struct {
}
