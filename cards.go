package paymego

import (
	"context"

	"github.com/sirupsen/logrus"
)

// cards check
func (c *SubscribeAPI) CardsCheck(ctx context.Context, requestID, token string) (*PaymeResponse, error) {
	receiptParams := map[string]interface{}{
		"token": token,
	}

	respCheckCard, err := c.sendRequest(ctx, requestID, "cards.check", receiptParams, true)

	if err != nil {
		logrus.Errorf("failed cards check(error - %v request-id - %s response - %v)", err, requestID, respCheckCard)
		return &PaymeResponse{}, err
	}
	return respCheckCard, nil
}

// cards remove
func (c *SubscribeAPI) CardsRemove(ctx context.Context, requestID, token string) (*PaymeResponse, error) {
	receiptParams := map[string]interface{}{
		"token": token,
	}

	respCheckCard, err := c.sendRequest(ctx, requestID, "cards.remove", receiptParams, true)

	if err != nil {
		logrus.Errorf("failed cards remove (error - %v request-id - %s response - %v)", err, requestID, respCheckCard)
		return &PaymeResponse{}, err
	}
	return respCheckCard, nil
}
