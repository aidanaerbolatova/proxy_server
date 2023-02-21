package proxy

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"proxy/internal/cache"
	"proxy/models"
)

type ProxyServer struct {
	cache *cache.Cache
}

func NewProxyServer(cache *cache.Cache) *ProxyServer {
	return &ProxyServer{
		cache: cache,
	}
}

func (p *ProxyServer) ProxyRequest(req models.Request) (models.Response, error) {
	if req.Method != "GET" {
		fmt.Println(req.Method)
		return models.Response{}, errors.New("you can use only method: GET")
	}
	_, err := url.Parse(req.URL)
	if err != nil {
		return models.Response{}, errors.New("invalid origin server URl")
	}
	// convert json to key for cache
	cacheKey, err := p.cache.ConvertCacheKey(req)
	if err != nil {
		return models.Response{}, err
	}
	// get pesponse from proxy
	response, ok := p.cache.Get(cacheKey)
	if ok {
		return response, nil
	}
	// if response does not exists, create response
	newRequest, err := createRequest(req)
	if err != nil {
		return models.Response{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()
	proxyResponse := createResponse(resp)
	//add response to cache
	p.cache.Set(cacheKey, proxyResponse)
	return proxyResponse, nil
}

func createRequest(request models.Request) (*http.Request, error) {
	newRequest, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return nil, errors.New("error with create new  proxy request")
	}
	for key, value := range request.Headers {
		newRequest.Header.Set(key, value)
	}
	return newRequest, nil
}

func createResponse(resp *http.Response) models.Response {
	var response models.Response
	response.Headers = map[string][]string{}
	for key, value := range resp.Header {
		response.Headers[key] = value
	}
	return models.Response{
		Id:      uuid.New().String(),
		Status:  resp.Status,
		Headers: response.Headers,
		Length:  int(resp.ContentLength),
	}
}
