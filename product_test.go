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

func TestProductList(t *testing.T) {
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
	_, err := client.Product.List(context.TODO())
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductGet(t *testing.T) {
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
	_, err := client.Product.Get(context.TODO(), "prd_XXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
