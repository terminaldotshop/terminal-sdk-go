// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

// ProductService contains methods and other services that help with interacting
// with the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options []option.RequestOption
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	return
}

func (r *ProductService) List(ctx context.Context, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProductListResponse struct {
	Result []ProductListResponseResult `json:"result,required"`
	JSON   productListResponseJSON     `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseJSON) RawJSON() string {
	return r.raw
}

type ProductListResponseResult struct {
	ID          string                             `json:"id,required"`
	Description string                             `json:"description,required"`
	Name        string                             `json:"name,required"`
	Variants    []ProductListResponseResultVariant `json:"variants,required"`
	JSON        productListResponseResultJSON      `json:"-"`
}

// productListResponseResultJSON contains the JSON metadata for the struct
// [ProductListResponseResult]
type productListResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ProductListResponseResultVariant struct {
	ID    string                               `json:"id,required"`
	Name  string                               `json:"name,required"`
	Price int64                                `json:"price,required"`
	JSON  productListResponseResultVariantJSON `json:"-"`
}

// productListResponseResultVariantJSON contains the JSON metadata for the struct
// [ProductListResponseResultVariant]
type productListResponseResultVariantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseResultVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseResultVariantJSON) RawJSON() string {
	return r.raw
}
