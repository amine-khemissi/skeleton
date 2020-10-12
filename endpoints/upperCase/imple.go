package upperCase

import (
	"context"
	"github.com/amine-khemissi/skeleton/def/upperCase"
	"strings"
)

func Uppercase(ctx context.Context, req upperCase.Request) (upperCase.Response, error) {

	return upperCase.Response{Value: strings.ToUpper(req.Value)}, nil
}
