package services

import (
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func CreateCheckoutSession(directorySlug string, priceCents int64, userEmail string) (string, error) {
	successURL := os.Getenv("APP_DOMAIN") + "/payment-success/" + directorySlug
	cancelURL := os.Getenv("APP_DOMAIN") + "/submit"

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		CustomerEmail:      stripe.String(userEmail),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Directory Listing"),
					},
					UnitAmount: stripe.Int64(priceCents),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
	}

	s, err := session.New(params)
	if err != nil {
		return "", err
	}
	return s.URL, nil
}
