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

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscription.New(context.TODO(), githubcomterminaldotshopterminalsdkgo.SubscriptionNewParams{
		Subscription: githubcomterminaldotshopterminalsdkgo.SubscriptionParam{
			ID:               githubcomterminaldotshopterminalsdkgo.F("sub_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			AddressID:        githubcomterminaldotshopterminalsdkgo.F("shp_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			CardID:           githubcomterminaldotshopterminalsdkgo.F("crd_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			Created:          githubcomterminaldotshopterminalsdkgo.F("2024-06-29T19:36:19.000Z"),
			Price:            githubcomterminaldotshopterminalsdkgo.F(int64(2200)),
			ProductVariantID: githubcomterminaldotshopterminalsdkgo.F("var_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			Quantity:         githubcomterminaldotshopterminalsdkgo.F(int64(1)),
			Next:             githubcomterminaldotshopterminalsdkgo.F("2025-02-01T19:36:19.000Z"),
			Schedule: githubcomterminaldotshopterminalsdkgo.F[githubcomterminaldotshopterminalsdkgo.SubscriptionScheduleUnionParam](githubcomterminaldotshopterminalsdkgo.SubscriptionScheduleWeeklyParam{
				Interval: githubcomterminaldotshopterminalsdkgo.F(int64(3)),
				Type:     githubcomterminaldotshopterminalsdkgo.F(githubcomterminaldotshopterminalsdkgo.SubscriptionScheduleWeeklyTypeWeekly),
			}),
		},
	})
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscription.Update(
		context.TODO(),
		"sub_XXXXXXXXXXXXXXXXXXXXXXXXX",
		githubcomterminaldotshopterminalsdkgo.SubscriptionUpdateParams{
			AddressID: githubcomterminaldotshopterminalsdkgo.F("shp_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			CardID:    githubcomterminaldotshopterminalsdkgo.F("crd_XXXXXXXXXXXXXXXXXXXXXXXXX"),
			Schedule: githubcomterminaldotshopterminalsdkgo.F[githubcomterminaldotshopterminalsdkgo.SubscriptionUpdateParamsScheduleUnion](githubcomterminaldotshopterminalsdkgo.SubscriptionUpdateParamsScheduleWeekly{
				Interval: githubcomterminaldotshopterminalsdkgo.F(int64(3)),
				Type:     githubcomterminaldotshopterminalsdkgo.F(githubcomterminaldotshopterminalsdkgo.SubscriptionUpdateParamsScheduleWeeklyTypeWeekly),
			}),
		},
	)
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
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
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Subscription.List(context.TODO())
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
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
	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Subscription.Delete(context.TODO(), "sub_XXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionGet(t *testing.T) {
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
	_, err := client.Subscription.Get(context.TODO(), "sub_XXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		var apierr *githubcomterminaldotshopterminalsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
