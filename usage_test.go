// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/terminal-go"
	"github.com/stainless-sdks/terminal-go/internal/testutil"
	"github.com/stainless-sdks/terminal-go/option"
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
	productGetResponse, err := client.Product.Get(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", productGetResponse.Result)
}
