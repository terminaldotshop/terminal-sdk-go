// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

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

func (r *CartService) List(ctx context.Context, opts ...option.RequestOption) (res *CartListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *CartService) SetItem(ctx context.Context, body CartSetItemParams, opts ...option.RequestOption) (res *CartSetItemResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CartListResponse struct {
	Result []CartListResponseResult `json:"result,required"`
	JSON   cartListResponseJSON     `json:"-"`
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

type CartListResponseResult struct {
	ID               string                     `json:"id,required"`
	ProductVariantID string                     `json:"productVariantID,required"`
	Quantity         int64                      `json:"quantity,required"`
	Subtotal         float64                    `json:"subtotal,required"`
	JSON             cartListResponseResultJSON `json:"-"`
}

// cartListResponseResultJSON contains the JSON metadata for the struct
// [CartListResponseResult]
type cartListResponseResultJSON struct {
	ID               apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	Subtotal         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CartListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartListResponseResultJSON) RawJSON() string {
	return r.raw
}

type CartSetItemResponse struct {
	Result []CartSetItemResponseResult `json:"result,required"`
	JSON   cartSetItemResponseJSON     `json:"-"`
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

type CartSetItemResponseResult struct {
	ID               string                        `json:"id,required"`
	ProductVariantID string                        `json:"productVariantID,required"`
	Quantity         int64                         `json:"quantity,required"`
	Subtotal         float64                       `json:"subtotal,required"`
	JSON             cartSetItemResponseResultJSON `json:"-"`
}

// cartSetItemResponseResultJSON contains the JSON metadata for the struct
// [CartSetItemResponseResult]
type cartSetItemResponseResultJSON struct {
	ID               apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	Subtotal         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CartSetItemResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartSetItemResponseResultJSON) RawJSON() string {
	return r.raw
}

type CartSetItemParams struct {
	ProductVariantID param.Field[string] `json:"productVariantID,required"`
	Quantity         param.Field[int64]  `json:"quantity,required"`
}

func (r CartSetItemParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
