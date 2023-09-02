package paymego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type SubscribeAPI struct {
	headers    xAuthHeaders
	baseURL    string
	httpClient http.Client
	logger     *log.Logger
}

type SubsribeAPIOpts struct {
	PaycomID   string
	PaycomKey  string
	Logger     *log.Logger
	HTTPClient http.Client
	BaseURL    string
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
	}

	return subscribeAPI, nil
}

func (c *SubscribeAPI) sendRequest(requestID, method string, params interface{}, withID bool) (*PaymeResponse, error) {
	data := map[string]interface{}{
		"id":     requestID,
		"method": method,
		"params": params,
	}

	requestBody, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if withID {
		request.Header.Set("X-Auth", c.headers.paycomID)
	} else {
		request.Header.Set("X-Auth", fmt.Sprintf("%s:%s", c.headers.paycomID, c.headers.paycomKey))
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
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

	// handle error case with payme specific error codes
	responseJson, err = handleErrorResponse(responseJson)

	if err != nil {
		logrus.Errorf("error response from payme response - %v error - %v params - %s", responseJson.Error, err, params)
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
