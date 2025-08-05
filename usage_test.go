// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo_test

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
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	products, err := client.Product.List(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", products.Data)
}
