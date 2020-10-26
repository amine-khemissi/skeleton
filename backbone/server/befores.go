package server

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/header"
	"github.com/google/uuid"
)

func addContextID(ctx context.Context, r *http.Request) context.Context {
	return header.Add(ctx, header.ContextID, uuid.New().String())
}
