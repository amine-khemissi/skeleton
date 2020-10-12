package uppercase

import (
	"context"
	"strings"

	"github.com/amine-khemissi/skeleton/def/uppercase"
)

func Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error) {

	return uppercase.Response{Value: strings.ToUpper(req.Value)}, nil
}
