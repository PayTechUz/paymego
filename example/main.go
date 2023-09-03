package main

import (
	"context"
	"fmt"
	"time"

	"github.com/paytechuz/paymego"
)

func main() {

	ctx := context.Background()

	s, err := paymego.NewSubscribeAPI(paymego.SubsribeAPIOpts{
		PaycomID:  "paycom-id",
		PaycomKey: "paycom-key",
		BaseURL:   "https://checkout.test.paycom.uz/api/",
		Timeout:   5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	resp, err := s.CardsCheck(ctx, "uuid4", "card-token")

	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", resp.Result)
}
