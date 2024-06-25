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

func (r *CartService) List(ctx context.Context, opts ...option.RequestOption) (res *CartListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *CartService) SetCard(ctx context.Context, body CartSetCardParams, opts ...option.RequestOption) (res *CartSetCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *CartService) SetItem(ctx context.Context, body CartSetItemParams, opts ...option.RequestOption) (res *CartSetItemResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/item"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *CartService) SetShipping(ctx context.Context, body CartSetShippingParams, opts ...option.RequestOption) (res *CartSetShippingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CartListResponse struct {
	Result shared.Cart          `json:"result,required"`
	JSON   cartListResponseJSON `json:"-"`
}

// cartListResponseJSON contains the JSON metadata for the struct
// [CartListResponse]
type cartListResponseJSON struct {
	Result      apijson.Field
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
	Result CartSetCardResponseResult `json:"result,required"`
	JSON   cartSetCardResponseJSON   `json:"-"`
}

// cartSetCardResponseJSON contains the JSON metadata for the struct
// [CartSetCardResponse]
type cartSetCardResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetCardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetCardResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetCardResponseResult string

const (
	CartSetCardResponseResultOk CartSetCardResponseResult = "ok"
)

func (r CartSetCardResponseResult) IsKnown() bool {
	switch r {
	case CartSetCardResponseResultOk:
		return true
	}
	return false
}

type CartSetItemResponse struct {
	Result shared.Cart             `json:"result,required"`
	JSON   cartSetItemResponseJSON `json:"-"`
}

// cartSetItemResponseJSON contains the JSON metadata for the struct
// [CartSetItemResponse]
type cartSetItemResponseJSON struct {
	Result      apijson.Field
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
	Result CartSetShippingResponseResult `json:"result,required"`
	JSON   cartSetShippingResponseJSON   `json:"-"`
}

// cartSetShippingResponseJSON contains the JSON metadata for the struct
// [CartSetShippingResponse]
type cartSetShippingResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartSetShippingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetShippingResponseJSON) RawJSON() string {
	return r.raw
}

type CartSetShippingResponseResult string

const (
	CartSetShippingResponseResultOk CartSetShippingResponseResult = "ok"
)

func (r CartSetShippingResponseResult) IsKnown() bool {
	switch r {
	case CartSetShippingResponseResultOk:
		return true
	}
	return false
}

type CartSetCardParams struct {
	CardID param.Field[string] `json:"cardID,required"`
}

func (r CartSetCardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CartSetItemParams struct {
	ProductVariantID param.Field[string] `json:"productVariantID,required"`
	Quantity         param.Field[int64]  `json:"quantity,required"`
}

func (r CartSetItemParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CartSetShippingParams struct {
	ShippingID param.Field[string] `json:"shippingID,required"`
}

func (r CartSetShippingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
