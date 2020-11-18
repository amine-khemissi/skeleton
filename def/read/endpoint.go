package read

type Request struct {
	ClientID string `json:"-" path:"clientID"`
}

type Response struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
