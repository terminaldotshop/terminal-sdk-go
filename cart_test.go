// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/internal/testutil"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

func TestCartClear(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.Clear(context.TODO())
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartConvert(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.Convert(context.TODO())
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.Get(context.TODO())
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCartSetAddress(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.SetAddress(context.TODO(), githubcomterminaldotshopterminalsdkgo.CartSetAddressParams{
		AddressID: githubcomterminaldotshopterminalsdkgo.F("shp_XXXXXXXXXXXXXXXXXXXXXXXXX"),
	})
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
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
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.SetCard(context.TODO(), githubcomterminaldotshopterminalsdkgo.CartSetCardParams{
		CardID: githubcomterminaldotshopterminalsdkgo.F("crd_XXXXXXXXXXXXXXXXXXXXXXXXX"),
	})
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
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
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Cart.SetItem(context.TODO(), githubcomterminaldotshopterminalsdkgo.CartSetItemParams{
		ProductVariantID: githubcomterminaldotshopterminalsdkgo.F("var_XXXXXXXXXXXXXXXXXXXXXXXXX"),
		Quantity:         githubcomterminaldotshopterminalsdkgo.F(int64(2)),
	})
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
