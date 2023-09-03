package paymego

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	PayForOrderReasonID = "6"
	Description         = "Merchant transaction for order - %s "
)

type CardData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type PaymentData struct {
	OrderID  string
	CardData CardData
}

type PaymentDetails struct {
	Client PaymentData
	Driver PaymentData
	Amount float64
}

type Account struct {
	OrderID string `json:"order_id"`
	CardID  string `json:"card_id"`
	Reason  string `json:"reason"`
}

// create merchant check
func (c *SubscribeAPI) createCheck(data PaymentDetails) (createdReceiptsID string, err error) {
	requestID := fmt.Sprintf("ReceiptsCreate:MerchantTransaction:%s", data.Client.OrderID)

	receiptParams := map[string]interface{}{
		"amount": data.Amount,
		"account": Account{
			OrderID: data.Client.OrderID,
			CardID:  data.Client.CardData.ID,
			Reason:  PayForOrderReasonID, // payment for order
		},
		"description": fmt.Sprintf(Description, data.Client.OrderID),
	}

	respCreateReceipt, err := c.sendRequest(requestID, "receipts.create", receiptParams, false)

	if err != nil {
		logrus.Errorf("failed receipts create (error - %v request-id - %s response - %v)", err, requestID, respCreateReceipt)
		return "", err
	}

	createdReceiptsID = respCreateReceipt.Result.Receipt.ID

	logrus.Infof("receipts created for order - %v request-id - %s transaction-id - %s", data.Client.OrderID, requestID, createdReceiptsID)

	return createdReceiptsID, nil
}

// pay merchant check
func (c *SubscribeAPI) payCheck(data PaymentDetails, createdReceiptsID string) (paidReceiptsID string, err error) {
	// Create receipt request-id
	requestID := fmt.Sprintf("ReceiptsPay:%s", data.Client.OrderID)

	receiptParams := map[string]interface{}{
		"id":    createdReceiptsID,          // check id
		"token": data.Client.CardData.Token, // payer card token
	}

	respPayReceipt, err := c.sendRequest(requestID, "receipts.pay", receiptParams, false)

	if err != nil {
		logrus.Errorf("failed receipts pay (error - %v request-id - %s receipts-id %s response - %v)", err, requestID, createdReceiptsID, respPayReceipt.Error)
		return "", err
	}

	paidReceiptsID = respPayReceipt.Result.Receipt.ID

	logrus.Infof("receipts paid for order - %v request-id - %s transaction-id - %s", data.Client.OrderID, requestID, paidReceiptsID)

	return paidReceiptsID, nil
}

// create and pay merchant check
func (c *SubscribeAPI) Pay(data PaymentDetails) (paidReceiptsID string, err error) {
	// create check
	createdReceiptsID, err := c.createCheck(data)

	if err != nil {
		return createdReceiptsID, err
	}

	// pay check
	paidReceiptsID, err = c.payCheck(data, createdReceiptsID)

	if err != nil {
		return createdReceiptsID, err
	}

	return paidReceiptsID, nil
}

// create check p2p
func (c *SubscribeAPI) createCheckP2P(data PaymentDetails) (createdReceiptsID string, err error) {
	// Create receipt request-id
	requestID := fmt.Sprintf("ReceiptsCreate:P2PTransaction:%s", data.Client.OrderID)
	description := fmt.Sprintf("P2PTransaction for order %s", data.Client.OrderID)

	receiptParams := map[string]interface{}{
		"amount":      data.Amount,
		"token":       data.Driver.CardData.Token,
		"description": description,
	}

	respCreateReceipt, err := c.sendRequest(requestID, "receipts.p2p", receiptParams, false)

	if err != nil {
		logrus.Errorf("failed receipts create p2p (error - %v request-id - %s response - %v", err, requestID, respCreateReceipt)
		return "", err
	}

	createdReceiptsID = respCreateReceipt.Result.Receipt.ID

	logrus.Infof("receipts p2p created for order - %v request-id - %s transaction id - %s", data.Client.OrderID, requestID, createdReceiptsID)

	return createdReceiptsID, nil
}

// pay p2p check
func (c *SubscribeAPI) payCheckP2P(data PaymentDetails, createdReceiptsID string) (paidReceiptsID string, err error) {
	// Create receipt request-id
	requestID := fmt.Sprintf("ReceiptsPay:%s", data.Client.OrderID)

	receiptParams := map[string]interface{}{
		"id":    createdReceiptsID,          // check id
		"token": data.Client.CardData.Token, // payer card token
	}

	respPayReceipt, err := c.sendRequest(requestID, "receipts.pay", receiptParams, false)

	if err != nil {
		logrus.Errorf("failed receipts pay (error - %v request-id - %s receipts-id %s response - %v)", err, requestID, createdReceiptsID, respPayReceipt)
		return "", err
	}

	paidReceiptsID = respPayReceipt.Result.Receipt.ID

	logrus.Infof("receipts paid for order - %v request-id - %s transaction id - %s", data.Client.OrderID, requestID, paidReceiptsID)

	return paidReceiptsID, nil
}

// create p2p check and pay p2p check
func (c *SubscribeAPI) PayP2P(data PaymentDetails) (paidReceiptsID string, err error) {
	// create check
	createdReceiptsID, err := c.createCheckP2P(data)

	if err != nil {
		return createdReceiptsID, err
	}

	// pay check
	paidReceiptsID, err = c.payCheckP2P(data, createdReceiptsID)

	if err != nil {
		return createdReceiptsID, err
	}

	return paidReceiptsID, nil
}
