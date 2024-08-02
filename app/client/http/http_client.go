package http

import (
	"auth-service/app/common/exceptions"
	"auth-service/app/common/logs"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"time"
)

type (
	Header struct {
		Key   string
		Value string
	}

	ClientComponent[T any, E any] struct {
		HttpMethod     string
		UrlApi         string
		ConnectTimeOut int32
		ActiveTimeOut  int32
		Headers        []Header
		RequestBody    *T
		ResponseBody   *E
	}
)

func (c ClientComponent[T, E]) Execute(ctx context.Context) error {
	client := &http.Client{
		Timeout: time.Duration(rand.Int31n(c.ActiveTimeOut)) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout: time.Second * 5,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(rand.Int31n(c.ConnectTimeOut))*time.Millisecond)
			},
		},
	}
	var request *http.Request
	var response *http.Response
	var err error = nil

	// Set request body
	if reflect.ValueOf(c.RequestBody).IsZero() || c.RequestBody == nil {
		request, err = http.NewRequest(c.HttpMethod, c.UrlApi, nil)
	}
	requestBody, err := json.Marshal(c.ResponseBody)
	exceptions.PanicLogging(err)

	// Log request body
	logs.NewLogger().Info("Request Body ", string(requestBody))
	requestBodyByte := bytes.NewBuffer(requestBody)
	request, err = http.NewRequestWithContext(ctx, c.HttpMethod, c.UrlApi, requestBodyByte)
	exceptions.PanicLogging(err)

	// Set header
	request.Header.Set("Content-Type", "application/json")
	for _, header := range c.Headers {
		request.Header.Set(header.Key, header.Value)
	}

	// Logging before
	logs.NewLogger().Info("Request url ", c.UrlApi)
	logs.NewLogger().Info("Request method ", c.HttpMethod)
	logs.NewLogger().Info("Request header ", request.Header)

	// time
	start := time.Now()

	response, err = client.Do(request)
	if err != nil {
		return err
	}

	elapsed := time.Now().Sub(start)
	responseBody, err := io.ReadAll(response.Body)
	exceptions.PanicLogging(err)

	err = json.Unmarshal(responseBody, &c.ResponseBody)
	exceptions.PanicLogging(err)

	logs.NewLogger().Info("Received response for ", c.UrlApi, "in", elapsed.Microseconds(), "ms")
	logs.NewLogger().Info("Response Header ", response.Header)
	logs.NewLogger().Info("Response Status ", response.Status)
	logs.NewLogger().Info("Response http version ", response.Proto)
	logs.NewLogger().Info("Response body ", string(responseBody))

	return nil
}
