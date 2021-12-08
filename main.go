package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeRequest(ctx context.Context, route string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad status code %d", resp.StatusCode)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(bs), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*350)
	defer cancel()

	resp, err := makeRequest(ctx, "https://www.google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
