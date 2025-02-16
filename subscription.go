// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Create a subscription for the current user.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the subscriptions associated with the current user.
func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a subscription for the current user.
func (r *SubscriptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("subscription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Subscription to a Terminal shop product.
type Subscription struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// ID of the shipping address used for the subscription.
	AddressID string `json:"addressID,required"`
	// ID of the card used for the subscription.
	CardID string `json:"cardID,required"`
	// Frequency of the subscription.
	Frequency SubscriptionFrequency `json:"frequency,required"`
	// ID of the product variant being subscribed to.
	ProductVariantID string `json:"productVariantID,required"`
	// Quantity of the subscription.
	Quantity int64 `json:"quantity,required"`
	// Next shipment and billing date for the subscription.
	Next string `json:"next"`
	// Schedule of the subscription.
	Schedule SubscriptionSchedule `json:"schedule"`
	JSON     subscriptionJSON     `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID               apijson.Field
	AddressID        apijson.Field
	CardID           apijson.Field
	Frequency        apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	Next             apijson.Field
	Schedule         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// Frequency of the subscription.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyFixed   SubscriptionFrequency = "fixed"
	SubscriptionFrequencyDaily   SubscriptionFrequency = "daily"
	SubscriptionFrequencyWeekly  SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly SubscriptionFrequency = "monthly"
	SubscriptionFrequencyYearly  SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyFixed, SubscriptionFrequencyDaily, SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// Schedule of the subscription.
type SubscriptionSchedule struct {
	Type     SubscriptionScheduleType `json:"type,required"`
	Interval int64                    `json:"interval"`
	JSON     subscriptionScheduleJSON `json:"-"`
	union    SubscriptionScheduleUnion
}

// subscriptionScheduleJSON contains the JSON metadata for the struct
// [SubscriptionSchedule]
type subscriptionScheduleJSON struct {
	Type        apijson.Field
	Interval    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r subscriptionScheduleJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionSchedule) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionSchedule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionScheduleUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [SubscriptionScheduleType],
// [SubscriptionScheduleObject].
func (r SubscriptionSchedule) AsUnion() SubscriptionScheduleUnion {
	return r.union
}

// Schedule of the subscription.
//
// Union satisfied by [SubscriptionScheduleType] or [SubscriptionScheduleObject].
type SubscriptionScheduleUnion interface {
	implementsSubscriptionSchedule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionScheduleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionScheduleType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionScheduleObject{}),
		},
	)
}

type SubscriptionScheduleType struct {
	Type SubscriptionScheduleTypeType `json:"type,required"`
	JSON subscriptionScheduleTypeJSON `json:"-"`
}

// subscriptionScheduleTypeJSON contains the JSON metadata for the struct
// [SubscriptionScheduleType]
type subscriptionScheduleTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionScheduleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionScheduleTypeJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionScheduleType) implementsSubscriptionSchedule() {}

type SubscriptionScheduleTypeType string

const (
	SubscriptionScheduleTypeTypeFixed SubscriptionScheduleTypeType = "fixed"
)

func (r SubscriptionScheduleTypeType) IsKnown() bool {
	switch r {
	case SubscriptionScheduleTypeTypeFixed:
		return true
	}
	return false
}

type SubscriptionScheduleObject struct {
	Interval int64                          `json:"interval,required"`
	Type     SubscriptionScheduleObjectType `json:"type,required"`
	JSON     subscriptionScheduleObjectJSON `json:"-"`
}

// subscriptionScheduleObjectJSON contains the JSON metadata for the struct
// [SubscriptionScheduleObject]
type subscriptionScheduleObjectJSON struct {
	Interval    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionScheduleObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionScheduleObjectJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionScheduleObject) implementsSubscriptionSchedule() {}

type SubscriptionScheduleObjectType string

const (
	SubscriptionScheduleObjectTypeWeekly SubscriptionScheduleObjectType = "weekly"
)

func (r SubscriptionScheduleObjectType) IsKnown() bool {
	switch r {
	case SubscriptionScheduleObjectTypeWeekly:
		return true
	}
	return false
}

// Subscription to a Terminal shop product.
type SubscriptionParam struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID param.Field[string] `json:"id,required"`
	// ID of the shipping address used for the subscription.
	AddressID param.Field[string] `json:"addressID,required"`
	// ID of the card used for the subscription.
	CardID param.Field[string] `json:"cardID,required"`
	// Frequency of the subscription.
	Frequency param.Field[SubscriptionFrequency] `json:"frequency,required"`
	// ID of the product variant being subscribed to.
	ProductVariantID param.Field[string] `json:"productVariantID,required"`
	// Quantity of the subscription.
	Quantity param.Field[int64] `json:"quantity,required"`
	// Next shipment and billing date for the subscription.
	Next param.Field[string] `json:"next"`
	// Schedule of the subscription.
	Schedule param.Field[SubscriptionScheduleUnionParam] `json:"schedule"`
}

func (r SubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule of the subscription.
type SubscriptionScheduleParam struct {
	Type     param.Field[SubscriptionScheduleType] `json:"type,required"`
	Interval param.Field[int64]                    `json:"interval"`
}

func (r SubscriptionScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionScheduleParam) implementsSubscriptionScheduleUnionParam() {}

// Schedule of the subscription.
//
// Satisfied by [SubscriptionScheduleTypeParam], [SubscriptionScheduleObjectParam],
// [SubscriptionScheduleParam].
type SubscriptionScheduleUnionParam interface {
	implementsSubscriptionScheduleUnionParam()
}

type SubscriptionScheduleTypeParam struct {
	Type param.Field[SubscriptionScheduleTypeType] `json:"type,required"`
}

func (r SubscriptionScheduleTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionScheduleTypeParam) implementsSubscriptionScheduleUnionParam() {}

type SubscriptionScheduleObjectParam struct {
	Interval param.Field[int64]                          `json:"interval,required"`
	Type     param.Field[SubscriptionScheduleObjectType] `json:"type,required"`
}

func (r SubscriptionScheduleObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionScheduleObjectParam) implementsSubscriptionScheduleUnionParam() {}

type SubscriptionNewResponse struct {
	Data SubscriptionNewResponseData `json:"data,required"`
	JSON subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewResponseData string

const (
	SubscriptionNewResponseDataOk SubscriptionNewResponseData = "ok"
)

func (r SubscriptionNewResponseData) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseDataOk:
		return true
	}
	return false
}

type SubscriptionListResponse struct {
	// List of subscriptions.
	Data []Subscription               `json:"data,required"`
	JSON subscriptionListResponseJSON `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionDeleteResponse struct {
	Data SubscriptionDeleteResponseData `json:"data,required"`
	JSON subscriptionDeleteResponseJSON `json:"-"`
}

// subscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [SubscriptionDeleteResponse]
type subscriptionDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionDeleteResponseData string

const (
	SubscriptionDeleteResponseDataOk SubscriptionDeleteResponseData = "ok"
)

func (r SubscriptionDeleteResponseData) IsKnown() bool {
	switch r {
	case SubscriptionDeleteResponseDataOk:
		return true
	}
	return false
}

type SubscriptionNewParams struct {
	// Subscription to a Terminal shop product.
	Subscription SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}
