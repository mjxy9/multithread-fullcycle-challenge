package httpclient

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
}

func (hc *HttpClient) GetAddress(url string) ([]byte, error) {

	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	defer cancel()

	if err != nil {
		log.Printf("could not generate http request, error: %v", err)
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil || response.StatusCode != 200 {
		log.Printf("request failed, error: %v", err)
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
