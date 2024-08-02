package httpbin

import (
	"auth-service/app/client/http"
	"auth-service/app/common/exceptions"
	"auth-service/app/dto"
	"context"
)

func NewHttpBinRestClient() BinClient {
	return &BinRestClient{}
}

type BinRestClient struct {
}

func (h BinRestClient) PostMethod(ctx context.Context, requestBody *dto.HttpBin, response *map[string]interface{}) {
	var headers []http.Header
	headers = append(headers, http.Header{Key: "X-Key", Value: "123456"})
	httpClient := http.ClientComponent[dto.HttpBin, map[string]interface{}]{
		HttpMethod:     "POST",
		UrlApi:         "https://httpbin.org/post",
		RequestBody:    requestBody,
		ResponseBody:   response,
		Headers:        headers,
		ConnectTimeOut: 30000,
		ActiveTimeOut:  30000,
	}

	err := httpClient.Execute(ctx)
	exceptions.PanicLogging(err)
}
