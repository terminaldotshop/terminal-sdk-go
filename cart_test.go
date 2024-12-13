// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/terminal-go"
	"github.com/stainless-sdks/terminal-go/internal/testutil"
	"github.com/stainless-sdks/terminal-go/option"
)

func TestCartList(t *testing.T) {
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
	_, err := client.Cart.List(context.TODO())
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartSetCard(t *testing.T) {
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
	_, err := client.Cart.SetCard(context.TODO(), terminal.CartSetCardParams{
		CardID: terminal.F("crd_XXXXXXXXXXXXXXXXXXXXXXXXX"),
	})
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartSetItem(t *testing.T) {
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
	_, err := client.Cart.SetItem(context.TODO(), terminal.CartSetItemParams{
		ProductVariantID: terminal.F("var_XXXXXXXXXXXXXXXXXXXXXXXXX"),
		Quantity:         terminal.F(2.000000),
	})
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartSetShipping(t *testing.T) {
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
	_, err := client.Cart.SetShipping(context.TODO(), terminal.CartSetShippingParams{
		ShippingID: terminal.F("shp_XXXXXXXXXXXXXXXXXXXXXXXXX"),
	})
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
