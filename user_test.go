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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "string")
	if err != nil {
		var apierr *terminal.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
