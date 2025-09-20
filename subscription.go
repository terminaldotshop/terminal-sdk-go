// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

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
	opts = slices.Concat(r.Options, opts)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update card, address, or interval for an existing subscription.
func (r *SubscriptionService) Update(ctx context.Context, id string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("subscription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List the subscriptions associated with the current user.
func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a subscription for the current user.
func (r *SubscriptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SubscriptionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("subscription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the subscription with the given ID.
func (r *SubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SubscriptionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("subscription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	// Date the subscription was created.
	Created string `json:"created,required"`
	// Price of the subscription in cents (USD).
	Price int64 `json:"price,required"`
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
	Created          apijson.Field
	Price            apijson.Field
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
// Possible runtime types of the union are [SubscriptionScheduleFixed],
// [SubscriptionScheduleWeekly].
func (r SubscriptionSchedule) AsUnion() SubscriptionScheduleUnion {
	return r.union
}

// Schedule of the subscription.
//
// Union satisfied by [SubscriptionScheduleFixed] or [SubscriptionScheduleWeekly].
type SubscriptionScheduleUnion interface {
	implementsSubscriptionSchedule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionScheduleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionScheduleFixed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionScheduleWeekly{}),
		},
	)
}

type SubscriptionScheduleFixed struct {
	Type SubscriptionScheduleFixedType `json:"type,required"`
	JSON subscriptionScheduleFixedJSON `json:"-"`
}

// subscriptionScheduleFixedJSON contains the JSON metadata for the struct
// [SubscriptionScheduleFixed]
type subscriptionScheduleFixedJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionScheduleFixed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionScheduleFixedJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionScheduleFixed) implementsSubscriptionSchedule() {}

type SubscriptionScheduleFixedType string

const (
	SubscriptionScheduleFixedTypeFixed SubscriptionScheduleFixedType = "fixed"
)

func (r SubscriptionScheduleFixedType) IsKnown() bool {
	switch r {
	case SubscriptionScheduleFixedTypeFixed:
		return true
	}
	return false
}

type SubscriptionScheduleWeekly struct {
	Interval int64                          `json:"interval,required"`
	Type     SubscriptionScheduleWeeklyType `json:"type,required"`
	JSON     subscriptionScheduleWeeklyJSON `json:"-"`
}

// subscriptionScheduleWeeklyJSON contains the JSON metadata for the struct
// [SubscriptionScheduleWeekly]
type subscriptionScheduleWeeklyJSON struct {
	Interval    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionScheduleWeekly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionScheduleWeeklyJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionScheduleWeekly) implementsSubscriptionSchedule() {}

type SubscriptionScheduleWeeklyType string

const (
	SubscriptionScheduleWeeklyTypeWeekly SubscriptionScheduleWeeklyType = "weekly"
)

func (r SubscriptionScheduleWeeklyType) IsKnown() bool {
	switch r {
	case SubscriptionScheduleWeeklyTypeWeekly:
		return true
	}
	return false
}

type SubscriptionScheduleType string

const (
	SubscriptionScheduleTypeFixed  SubscriptionScheduleType = "fixed"
	SubscriptionScheduleTypeWeekly SubscriptionScheduleType = "weekly"
)

func (r SubscriptionScheduleType) IsKnown() bool {
	switch r {
	case SubscriptionScheduleTypeFixed, SubscriptionScheduleTypeWeekly:
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
	// Date the subscription was created.
	Created param.Field[string] `json:"created,required"`
	// Price of the subscription in cents (USD).
	Price param.Field[int64] `json:"price,required"`
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
// Satisfied by [SubscriptionScheduleFixedParam],
// [SubscriptionScheduleWeeklyParam], [SubscriptionScheduleParam].
type SubscriptionScheduleUnionParam interface {
	implementsSubscriptionScheduleUnionParam()
}

type SubscriptionScheduleFixedParam struct {
	Type param.Field[SubscriptionScheduleFixedType] `json:"type,required"`
}

func (r SubscriptionScheduleFixedParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionScheduleFixedParam) implementsSubscriptionScheduleUnionParam() {}

type SubscriptionScheduleWeeklyParam struct {
	Interval param.Field[int64]                          `json:"interval,required"`
	Type     param.Field[SubscriptionScheduleWeeklyType] `json:"type,required"`
}

func (r SubscriptionScheduleWeeklyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionScheduleWeeklyParam) implementsSubscriptionScheduleUnionParam() {}

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

type SubscriptionUpdateResponse struct {
	// Subscription to a Terminal shop product.
	Data Subscription                   `json:"data,required"`
	JSON subscriptionUpdateResponseJSON `json:"-"`
}

// subscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponse]
type subscriptionUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseJSON) RawJSON() string {
	return r.raw
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

type SubscriptionGetResponse struct {
	// Subscription to a Terminal shop product.
	Data Subscription                `json:"data,required"`
	JSON subscriptionGetResponseJSON `json:"-"`
}

// subscriptionGetResponseJSON contains the JSON metadata for the struct
// [SubscriptionGetResponse]
type subscriptionGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	// Subscription to a Terminal shop product.
	Subscription SubscriptionParam `json:"subscription"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateParams struct {
	// New shipping address ID for the subscription.
	AddressID param.Field[string] `json:"addressID"`
	// New payment method ID for the subscription.
	CardID param.Field[string] `json:"cardID"`
	// New schedule for the subscription.
	Schedule param.Field[SubscriptionUpdateParamsScheduleUnion] `json:"schedule"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New schedule for the subscription.
type SubscriptionUpdateParamsSchedule struct {
	Type     param.Field[SubscriptionUpdateParamsScheduleType] `json:"type,required"`
	Interval param.Field[int64]                                `json:"interval"`
}

func (r SubscriptionUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdateParamsSchedule) implementsSubscriptionUpdateParamsScheduleUnion() {}

// New schedule for the subscription.
//
// Satisfied by [SubscriptionUpdateParamsScheduleFixed],
// [SubscriptionUpdateParamsScheduleWeekly], [SubscriptionUpdateParamsSchedule].
type SubscriptionUpdateParamsScheduleUnion interface {
	implementsSubscriptionUpdateParamsScheduleUnion()
}

type SubscriptionUpdateParamsScheduleFixed struct {
	Type param.Field[SubscriptionUpdateParamsScheduleFixedType] `json:"type,required"`
}

func (r SubscriptionUpdateParamsScheduleFixed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdateParamsScheduleFixed) implementsSubscriptionUpdateParamsScheduleUnion() {}

type SubscriptionUpdateParamsScheduleFixedType string

const (
	SubscriptionUpdateParamsScheduleFixedTypeFixed SubscriptionUpdateParamsScheduleFixedType = "fixed"
)

func (r SubscriptionUpdateParamsScheduleFixedType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsScheduleFixedTypeFixed:
		return true
	}
	return false
}

type SubscriptionUpdateParamsScheduleWeekly struct {
	Interval param.Field[int64]                                      `json:"interval,required"`
	Type     param.Field[SubscriptionUpdateParamsScheduleWeeklyType] `json:"type,required"`
}

func (r SubscriptionUpdateParamsScheduleWeekly) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdateParamsScheduleWeekly) implementsSubscriptionUpdateParamsScheduleUnion() {}

type SubscriptionUpdateParamsScheduleWeeklyType string

const (
	SubscriptionUpdateParamsScheduleWeeklyTypeWeekly SubscriptionUpdateParamsScheduleWeeklyType = "weekly"
)

func (r SubscriptionUpdateParamsScheduleWeeklyType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsScheduleWeeklyTypeWeekly:
		return true
	}
	return false
}

type SubscriptionUpdateParamsScheduleType string

const (
	SubscriptionUpdateParamsScheduleTypeFixed  SubscriptionUpdateParamsScheduleType = "fixed"
	SubscriptionUpdateParamsScheduleTypeWeekly SubscriptionUpdateParamsScheduleType = "weekly"
)

func (r SubscriptionUpdateParamsScheduleType) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsScheduleTypeFixed, SubscriptionUpdateParamsScheduleTypeWeekly:
		return true
	}
	return false
}
