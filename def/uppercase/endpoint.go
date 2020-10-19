package uppercase

type Request struct {
	Value string `json:"-" qs:"value"`
}
type Response struct {
	Value string `json:"value"`
}
