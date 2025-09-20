// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
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

// Clear the current user's cart.
func (r *CartService) Clear(ctx context.Context, opts ...option.RequestOption) (res *CartClearResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Convert the current user's cart to an order.
func (r *CartService) Convert(ctx context.Context, opts ...option.RequestOption) (res *CartConvertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get the current user's cart.
func (r *CartService) Get(ctx context.Context, opts ...option.RequestOption) (res *CartGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the shipping address for the current user's cart.
func (r *CartService) SetAddress(ctx context.Context, body CartSetAddressParams, opts ...option.RequestOption) (res *CartSetAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart/address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the credit card for the current user's cart.
func (r *CartService) SetCard(ctx context.Context, body CartSetCardParams, opts ...option.RequestOption) (res *CartSetCardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Add an item to the current user's cart.
func (r *CartService) SetItem(ctx context.Context, body CartSetItemParams, opts ...option.RequestOption) (res *CartSetItemResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cart/item"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The current Terminal shop user's cart.
type Cart struct {
	// The subtotal and shipping amounts for the current user's cart.
	Amount CartAmount `json:"amount,required"`
	// An array of items in the current user's cart.
	Items []CartItem `json:"items,required"`
	// The subtotal of all items in the current user's cart, in cents (USD).
	Subtotal int64 `json:"subtotal,required"`
	// ID of the shipping address selected on the current user's cart.
	AddressID string `json:"addressID"`
	// ID of the card selected on the current user's cart.
	CardID string `json:"cardID"`
	// Shipping information for the current user's cart.
	Shipping CartShipping `json:"shipping"`
	JSON     cartJSON     `json:"-"`
}

// cartJSON contains the JSON metadata for the struct [Cart]
type cartJSON struct {
	Amount      apijson.Field
	Items       apijson.Field
	Subtotal    apijson.Field
	AddressID   apijson.Field
	CardID      apijson.Field
	Shipping    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartJSON) RawJSON() string {
	return r.raw
}

// The subtotal and shipping amounts for the current user's cart.
type CartAmount struct {
	// Subtotal of the current user's cart, in cents (USD).
	Subtotal int64 `json:"subtotal,required"`
	// Shipping amount of the current user's cart, in cents (USD).
	Shipping int64 `json:"shipping"`
	// Total amount after any discounts, in cents (USD).
	Total int64          `json:"total"`
	JSON  cartAmountJSON `json:"-"`
}

// cartAmountJSON contains the JSON metadata for the struct [CartAmount]
type cartAmountJSON struct {
	Subtotal    apijson.Field
	Shipping    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartAmountJSON) RawJSON() string {
	return r.raw
}

// An item in the current Terminal shop user's cart.
type CartItem struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// ID of the product variant for this item in the current user's cart.
	ProductVariantID string `json:"productVariantID,required"`
	// Quantity of the item in the current user's cart.
	Quantity int64 `json:"quantity,required"`
	// Subtotal of the item in the current user's cart, in cents (USD).
	Subtotal int64        `json:"subtotal,required"`
	JSON     cartItemJSON `json:"-"`
}

// cartItemJSON contains the JSON metadata for the struct [CartItem]
type cartItemJSON struct {
	ID               apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	Subtotal         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CartItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartItemJSON) RawJSON() string {
	return r.raw
}

// Shipping information for the current user's cart.
type CartShipping struct {
	// Shipping service name.
	Service string `json:"service"`
	// Shipping timeframe provided by the shipping carrier.
	Timeframe string           `json:"timeframe"`
	JSON      cartShippingJSON `json:"-"`
}

// cartShippingJSON contains the JSON metadata for the struct [CartShipping]
type cartShippingJSON struct {
	Service     apijson.Field
	Timeframe   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartShipping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartShippingJSON) RawJSON() string {
	return r.raw
}

type CartClearResponse struct {
	Data CartClearResponseData `json:"data,required"`
	JSON cartClearResponseJSON `json:"-"`
}

// cartClearResponseJSON contains the JSON metadata for the struct
// [CartClearResponse]
type cartClearResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartClearResponseJSON) RawJSON() string {
	return r.raw
}

type CartClearResponseData string

const (
	CartClearResponseDataOk CartClearResponseData = "ok"
)

func (r CartClearResponseData) IsKnown() bool {
	switch r {
	case CartClearResponseDataOk:
		return true
	}
	return false
}

type CartConvertResponse struct {
	// An order from the Terminal shop.
	Data Order                   `json:"data,required"`
	JSON cartConvertResponseJSON `json:"-"`
}

// cartConvertResponseJSON contains the JSON metadata for the struct
// [CartConvertResponse]
type cartConvertResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartConvertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartConvertResponseJSON) RawJSON() string {
	return r.raw
}

type CartGetResponse struct {
	// The current Terminal shop user's cart.
	Data Cart                `json:"data,required"`
	JSON cartGetResponseJSON `json:"-"`
}

// cartGetResponseJSON contains the JSON metadata for the struct [CartGetResponse]
type cartGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartGetResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetAddressResponse struct {
	Data CartSetAddressResponseData `json:"data,required"`
	JSON cartSetAddressResponseJSON `json:"-"`
}

// cartSetAddressResponseJSON contains the JSON metadata for the struct
// [CartSetAddressResponse]
type cartSetAddressResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetAddressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetAddressResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetAddressResponseData string

const (
	CartSetAddressResponseDataOk CartSetAddressResponseData = "ok"
)

func (r CartSetAddressResponseData) IsKnown() bool {
	switch r {
	case CartSetAddressResponseDataOk:
		return true
	}
	return false
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
	Data Cart                    `json:"data,required"`
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

type CartSetAddressParams struct {
	// ID of the shipping address to set for the current user's cart.
	AddressID param.Field[string] `json:"addressID,required"`
}

func (r CartSetAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r CartSetItemParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
