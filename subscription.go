// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terminal-sdk-go/shared"
)

// SubscriptionService contains methods and other services that help with
// interacting with the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SubscriptionNewResponse struct {
	Result bool                        `json:"result,required"`
	JSON   subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionListResponse struct {
	Result []shared.Subscription        `json:"result,required"`
	JSON   subscriptionListResponseJSON `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	CardID           param.Field[string]                         `json:"cardID,required"`
	Frequency        param.Field[SubscriptionNewParamsFrequency] `json:"frequency,required"`
	ProductVariantID param.Field[string]                         `json:"productVariantID,required"`
	Quantity         param.Field[int64]                          `json:"quantity,required"`
	ShippingID       param.Field[string]                         `json:"shippingID,required"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsFrequency string

const (
	SubscriptionNewParamsFrequencyFixed   SubscriptionNewParamsFrequency = "fixed"
	SubscriptionNewParamsFrequencyDaily   SubscriptionNewParamsFrequency = "daily"
	SubscriptionNewParamsFrequencyWeekly  SubscriptionNewParamsFrequency = "weekly"
	SubscriptionNewParamsFrequencyMonthly SubscriptionNewParamsFrequency = "monthly"
	SubscriptionNewParamsFrequencyYearly  SubscriptionNewParamsFrequency = "yearly"
)

func (r SubscriptionNewParamsFrequency) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsFrequencyFixed, SubscriptionNewParamsFrequencyDaily, SubscriptionNewParamsFrequencyWeekly, SubscriptionNewParamsFrequencyMonthly, SubscriptionNewParamsFrequencyYearly:
		return true
	}
	return false
}
