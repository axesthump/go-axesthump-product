package clientwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendRequest[Req any, Res any](ctx context.Context, req Req, url string) (Res, error) {
	buff := new(bytes.Buffer)
	err := json.NewEncoder(buff).Encode(req)
	var response Res

	if err != nil {
		return response, fmt.Errorf("marshaling json: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buff)
	if err != nil {
		return response, fmt.Errorf("creating http request: %w", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, fmt.Errorf("calling http: %w", err)
	}
	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("decode json: %w", err)
	}

	return response, nil
}
