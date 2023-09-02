# Payme Implementation

This MVP project helps for implementing <a href="https://developer.help.paycom.uz">payme-doc</a>.

# Installation
1 - `go get -v github.com/paytechuz/paymego`

## Example
```go
package main

import (
	"fmt"

	"github.com/paytechuz/paymego"
)

func main() {
	s, err := paymego.NewSubscribeAPI(paymego.SubsribeAPIOpts{
		PaycomID:  "paycom-id",
		PaycomKey: "paycom-key",
		BaseURL:   "https://checkout.test.paycom.uz/api/",
	})

	if err != nil {
		panic(err)
	}

	// Example usage:
	cardClient := paymego.CardData{
		ID:    "123456789",
		Token: "card-token",
	}

	cardDriver := paymego.CardData{
		ID:    "123456789",
		Token: "card-token",
	}

	paymentDataClient := paymego.PaymentData{
		OrderID:  "order123",
		CardData: cardClient,
	}

	paymentDataDriver := paymego.PaymentData{
		OrderID:  "order123",
		CardData: cardDriver,
	}

	paymentDetails := paymego.PaymentDetails{
		Client: paymentDataClient,
		Driver: paymentDataDriver,
		Amount: 100000,
	}

	resp, err := s.Pay(paymentDetails)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", resp)

}
```
