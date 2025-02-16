// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal_test

import (
	"context"
	"os"
	"testing"

	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/internal/testutil"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := terminal.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	subscription, err := client.Subscription.New(context.TODO(), terminal.SubscriptionNewParams{
		Subscription: terminal.SubscriptionParam{
			ID:               terminal.F("sub_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			AddressID:        terminal.F("shp_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			CardID:           terminal.F("crd_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			Frequency:        terminal.F(terminal.SubscriptionFrequencyFixed),
			ProductVariantID: terminal.F("var_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			Quantity:         terminal.F(int64(1)),
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", subscription.Data)
}
