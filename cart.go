// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/terminal-go/internal/apijson"
	"github.com/stainless-sdks/terminal-go/internal/param"
	"github.com/stainless-sdks/terminal-go/internal/requestconfig"
	"github.com/stainless-sdks/terminal-go/option"
	"github.com/stainless-sdks/terminal-go/shared"
)

// CartService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCartService] method instead.
type CartService struct {
	Options []option.RequestOption
}

// NewCartService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCartService(opts ...option.RequestOption) (r *CartService) {
	r = &CartService{}
	r.Options = opts
	return
}

// Get the current user's cart.
func (r *CartService) List(ctx context.Context, opts ...option.RequestOption) (res *CartListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the credit card for the current user's cart.
func (r *CartService) SetCard(ctx context.Context, body CartSetCardParams, opts ...option.RequestOption) (res *CartSetCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Add an item to the current user's cart.
func (r *CartService) SetItem(ctx context.Context, body CartSetItemParams, opts ...option.RequestOption) (res *CartSetItemResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/item"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the shipping address for the current user's cart.
func (r *CartService) SetShipping(ctx context.Context, body CartSetShippingParams, opts ...option.RequestOption) (res *CartSetShippingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CartListResponse struct {
	// The current Terminal shop user's cart.
	Data shared.Cart          `json:"data,required"`
	JSON cartListResponseJSON `json:"-"`
}

// cartListResponseJSON contains the JSON metadata for the struct
// [CartListResponse]
type cartListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartListResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetCardResponse struct {
	Data CartSetCardResponseData `json:"data,required"`
	JSON cartSetCardResponseJSON `json:"-"`
}

// cartSetCardResponseJSON contains the JSON metadata for the struct
// [CartSetCardResponse]
type cartSetCardResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetCardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetCardResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetCardResponseData string

const (
	CartSetCardResponseDataOk CartSetCardResponseData = "ok"
)

func (r CartSetCardResponseData) IsKnown() bool {
	switch r {
	case CartSetCardResponseDataOk:
		return true
	}
	return false
}

type CartSetItemResponse struct {
	// The current Terminal shop user's cart.
	Data shared.Cart             `json:"data,required"`
	JSON cartSetItemResponseJSON `json:"-"`
}

// cartSetItemResponseJSON contains the JSON metadata for the struct
// [CartSetItemResponse]
type cartSetItemResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetItemResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetShippingResponse struct {
	Data CartSetShippingResponseData `json:"data,required"`
	JSON cartSetShippingResponseJSON `json:"-"`
}

// cartSetShippingResponseJSON contains the JSON metadata for the struct
// [CartSetShippingResponse]
type cartSetShippingResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetShippingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetShippingResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetShippingResponseData string

const (
	CartSetShippingResponseDataOk CartSetShippingResponseData = "ok"
)

func (r CartSetShippingResponseData) IsKnown() bool {
	switch r {
	case CartSetShippingResponseDataOk:
		return true
	}
	return false
}

type CartSetCardParams struct {
	// ID of the credit card to set for the current user's cart.
	CardID param.Field[string] `json:"cardID,required"`
}

func (r CartSetCardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CartSetItemParams struct {
	// ID of the product variant to add to the cart.
	ProductVariantID param.Field[string] `json:"productVariantID,required"`
	// Quantity of the item to add to the cart.
	Quantity param.Field[float64] `json:"quantity,required"`
}

func (r CartSetItemParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CartSetShippingParams struct {
	// ID of the shipping address to set for the current user's cart.
	ShippingID param.Field[string] `json:"shippingID,required"`
}

func (r CartSetShippingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
