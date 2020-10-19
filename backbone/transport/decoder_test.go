package transport

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/require"
)

type w struct {
}

func (w w) Header() http.Header {
	return map[string][]string{}
}
func (w w) Write(out []byte) (int, error) {
	fmt.Println(string(out))
	return 0, nil
}

func (w w) WriteHeader(statusCode int) {
	fmt.Println("statusCode", statusCode)
}

type h struct {
	t *testing.T
}

func (h *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := h.t
	type TestStruct struct {
		Id      uint64 `json:"-"  path:"id"`
		Male    bool   `json:"-" qs:"male"`
		Age     int    `json:"age"`
		Name    string `json:"name"`
		Session string `json:"-" header:"session-id"`
	}
	decodedReq, err := decodeRequest(&TestStruct{}, r)
	assert.NoError(t, err)
	tStruct := decodedReq.(*TestStruct)
	assert.Equal(t, "john", tStruct.Name)
	assert.Equal(t, 18, tStruct.Age)
	assert.Equal(t, true, tStruct.Male)
	assert.Equal(t, "123456789", tStruct.Session)
	assert.Equal(t, uint64(42), tStruct.Id)

}
func TestDecodeRequest(t *testing.T) {

	b := bytes.NewBufferString(`{"name":"john","age":18}`)
	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/persons/{id}").Headers("session-id", "").Handler(&h{t})

	r, err := http.NewRequest(http.MethodGet, "http://what.com/persons/42?male=true", b)

	require.NoError(t, err)
	r.Header.Add("session-id", "123456789")
	router.ServeHTTP(&w{}, r)

}
