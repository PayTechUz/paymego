package paymego

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// cards check
func (c *SubscribeAPI) CardsCheck(ctx context.Context, requestID, token string, timeout ...time.Duration) (*PaymeResponse, error) {
	var requestTimeout time.Duration

	if len(timeout) > 0 {
		requestTimeout = timeout[0]
	} else {
		requestTimeout = c.timeout
	}

	receiptParams := map[string]interface{}{
		"token": token,
	}

	respCheckCard, err := c.sendRequest(ctx, requestID, "cards.check", receiptParams, true, requestTimeout)

	if err != nil {
		logrus.Errorf("failed cards check(error - %v request-id - %s response - %v)", err, requestID, respCheckCard)
		return &PaymeResponse{}, err
	}
	return respCheckCard, nil
}

// cards remove
func (c *SubscribeAPI) CardsRemove(ctx context.Context, requestID, token string, timeout ...time.Duration) (*PaymeResponse, error) {
	var requestTimeout time.Duration

	if len(timeout) > 0 {
		requestTimeout = timeout[0]
	} else {
		requestTimeout = c.timeout
	}

	receiptParams := map[string]interface{}{
		"token": token,
	}

	respCheckCard, err := c.sendRequest(ctx, requestID, "cards.remove", receiptParams, true, requestTimeout)

	if err != nil {
		logrus.Errorf("failed cards remove (error - %v request-id - %s response - %v)", err, requestID, respCheckCard)
		return &PaymeResponse{}, err
	}
	return respCheckCard, nil
}
