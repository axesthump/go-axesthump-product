package clientwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func SendRequest[Req any, Res any](ctx context.Context, req Req, url string) (Res, error) {
	rawJSON, err := json.Marshal(req)
	var response Res

	if err != nil {
		return response, errors.Wrap(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(rawJSON))
	if err != nil {
		return response, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return response, errors.Wrap(err, "decode json")
	}

	return response, nil
}
