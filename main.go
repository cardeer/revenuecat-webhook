package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ApiVersion string `json:"api_version"`
	Event      struct {
		Aliases                  []string          `json:"aliases"`
		AppUserID                string            `json:"app_user_id"`
		CountryCode              string            `json:"country_code"`
		Currency                 string            `json:"currency"`
		EntitlementID            string            `json:"entitlement_id"`
		EntitlementIDs           []string          `json:"entitlement_ids"`
		Environment              string            `json:"environment"`
		EventTimestampMs         int64             `json:"event_timestamp_ms"`
		ExpirationAtMs           int64             `json:"expiration_at_ms"`
		ID                       string            `json:"id"`
		IsFamilyShare            bool              `json:"is_family_share"`
		OriginalAppUserID        string            `json:"original_app_user_id"`
		OriginalTransactionID    string            `json:"original_transaction_id"`
		PeriodType               string            `json:"period_type"`
		PresentedOfferingID      string            `json:"presented_offering_id"`
		Price                    float64           `json:"price"`
		PriceInPurchasedCurrency float64           `json:"price_in_purchased_currency"`
		ProductID                string            `json:"product_id"`
		PurchasedAtMs            int64             `json:"purchased_at_ms"`
		Store                    string            `json:"store"`
		SubscriberAttributes     map[string]string `json:"subscriber_attributes"`
		TakehomePercentage       float64           `json:"takehome_percentage"`
		TransactionID            string            `json:"transaction_id"`
		Type                     string            `json:"type"`
	} `json:"event"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		transaction := new(Transaction)
		if err := c.BodyParser(transaction); err != nil {
			return err
		}
		fmt.Println(transaction)
		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
