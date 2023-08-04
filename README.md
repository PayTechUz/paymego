# Payme Implementation

This MVP project helps for implementing <a href="https://developer.help.paycom.uz">payme-doc</a>.

# Installation
1 - `go get -v github.com/paytechuz/paymego`

## Example
```go

package main

import (
	"context"
	"fmt"
	"github.com/paytechuz/paymego"
)

func main() {
	ctx := context.Background()
	s, err := paymego.NewSubscribeAPI(paymego.SubsribeAPIOpts{
		PaycomID:  "paycom-id",
		PaycomKey: "paycom-key",
		BaseURL:   "paycom-api-url",
	})
	if err != nil {
		panic(err)
	}
	r, err := s.CreateReceipt(ctx, "2134567879654321", "blabl", "awdaw", 2500, paymego.Account{
		OrderID:  "2383",
		CardID:   "129812",
		ReasonID: "12",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("response %+v", r)
}
```
