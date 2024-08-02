package httpbin

import (
	"auth-service/app/dto"
	"context"
)

type BinClient interface {
	PostMethod(ctx context.Context, requestBody *dto.HttpBin, response *map[string]interface{})
}
