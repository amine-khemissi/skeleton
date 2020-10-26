package header

import (
	"context"
	"sync"
)

type headersType int

const (
	headersKey headersType = iota
)

type headers struct {
	locker *sync.Mutex
	values map[string]string
}

func (h *headers) Add(k string, v string) {
	h.locker.Lock()
	defer h.locker.Unlock()
	h.values[k] = v
}
func newHeaders() *headers {
	return &headers{
		locker: &sync.Mutex{},
		values: map[string]string{},
	}
}

func Add(ctx context.Context, k string, v string) context.Context {
	h, isHeaders := ctx.Value(headersKey).(*headers)
	if h == nil || !isHeaders {
		h = newHeaders()
	}
	h.Add(k, v)
	return context.WithValue(ctx, headersKey, h)
}

func Get(ctx context.Context, k string) string {
	h, isHeaders := ctx.Value(headersKey).(*headers)
	if h == nil || !isHeaders {
		h = newHeaders()
	}
	return h.values[k]
}

func GetAll(ctx context.Context) map[string]string {
	h, isHeaders := ctx.Value(headersKey).(*headers)
	if h == nil || !isHeaders {
		h = newHeaders()
	}
	return h.values

}
