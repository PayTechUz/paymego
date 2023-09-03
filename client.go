package paymego

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type SubscribeAPI struct {
	headers    xAuthHeaders
	baseURL    string
	httpClient http.Client
	logger     *log.Logger
	timeout    time.Duration
}

type SubsribeAPIOpts struct {
	PaycomID   string
	PaycomKey  string
	Logger     *log.Logger
	HTTPClient http.Client
	BaseURL    string
	Timeout    time.Duration
}

type xAuthHeaders struct {
	paycomID  string
	paycomKey string
}

// NewSubscribeAPI returns new instance of SubscribeAPI
func NewSubscribeAPI(args SubsribeAPIOpts) (SubscribeAPI, error) {
	err := args.validate()
	if err != nil {
		return SubscribeAPI{}, err
	}

	subscribeAPI := SubscribeAPI{
		httpClient: args.HTTPClient,
		baseURL:    args.BaseURL,
		logger:     args.Logger,
		headers:    getXAuthHeaders(args.PaycomID, args.PaycomKey),
		timeout:    args.Timeout,
	}

	return subscribeAPI, nil
}

func (c *SubscribeAPI) sendRequest(
	ctx context.Context,
	requestID, method string,
	params interface{},
	withID bool,
	timeout ...time.Duration,
) (*PaymeResponse, error) {
	var requestTimeout time.Duration

	if len(timeout) > 0 {
		requestTimeout = timeout[0]
	} else {
		requestTimeout = c.timeout
	}

	// Create a context with the specified timeout.
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	data := map[string]interface{}{
		"id":     requestID,
		"method": method,
		"params": params,
	}

	requestBody, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if withID {
		req.Header.Set("X-Auth", c.headers.paycomID)
	} else {
		req.Header.Set("X-Auth", fmt.Sprintf("%s:%s", c.headers.paycomID, c.headers.paycomKey))
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	response, err := client.Do(req)

	timeoutError := errors.Is(err, context.DeadlineExceeded)

	if err != nil {
		if timeoutError {
			return nil, &PaymeTimeoutError{}
		}
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var responseJson PaymeResponse
	err = json.Unmarshal(responseBody, &responseJson)

	if err != nil {
		return nil, err
	}

	// Handle error response with payme specific error codes
	responseJson, err = handleErrorResponse(responseJson)

	if err != nil {
		logrus.Errorf("error response from payme - %v error - %v", responseJson.Error, err)
	}

	return &responseJson, err
}

func (s SubsribeAPIOpts) validate() error {
	if s.PaycomID == "" {
		return &ErrEmptyOrInvalidPaycomID{}
	}
	if s.PaycomKey == "" {
		return &ErrEmptyOrInvalidPaycomKey{}
	}
	return nil
}

func getXAuthHeaders(paycomID, paycomKey string) xAuthHeaders {
	return xAuthHeaders{paycomID: paycomID, paycomKey: paycomKey}
}
