package count

type Request struct {
	Value string `json:"value"`
}
type Response struct {
	Count int `json:"count"`
}
