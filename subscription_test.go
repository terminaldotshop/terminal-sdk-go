// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/internal/testutil"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

func TestSubscriptionNew(t *testing.T) {
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
	_, err := client.Subscriptions.New(context.TODO(), terminal.SubscriptionNewParams{
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
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionList(t *testing.T) {
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
	_, err := client.Subscriptions.List(context.TODO())
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionDelete(t *testing.T) {
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
	_, err := client.Subscriptions.Delete(context.TODO(), "sub_XXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
