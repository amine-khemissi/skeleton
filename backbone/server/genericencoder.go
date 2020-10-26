package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/header"

	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
)

func encodeHeaders(ctx context.Context, w http.ResponseWriter) {
	for k, v := range header.GetAll(ctx) {
		w.Header().Add(k, v)
	}
}

func genericEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	encodeHeaders(ctx, w)
	return json.NewEncoder(w).Encode(response)
}

func genericErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	encodeHeaders(ctx, w)
	skltErr, isSkltErr := err.(*errorsklt.InternalErr)
	if isSkltErr {
		w.WriteHeader(skltErr.Code)
		bts, _ := json.Marshal(skltErr.Stack)
		w.Write(bts)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	bts, _ := json.Marshal(err)
	w.Write(bts)
}
