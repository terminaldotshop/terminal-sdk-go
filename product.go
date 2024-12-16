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

// List all products for sale in the Terminal shop.
func (r *ProductService) List(ctx context.Context, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Product sold in the Terminal shop.
type Product struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Description of the product.
	Description string `json:"description,required"`
	// Name of the product.
	Name string `json:"name,required"`
	// List of variants of the product.
	Variants []ProductVariant `json:"variants,required"`
	// Order of the product used when displaying a sorted list of products.
	Order int64 `json:"order"`
	// Whether the product must be or can be subscribed to.
	Subscription ProductSubscription `json:"subscription"`
	// Tags for the product.
	Tags map[string]string `json:"tags"`
	JSON productJSON       `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Variants     apijson.Field
	Order        apijson.Field
	Subscription apijson.Field
	Tags         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

// Whether the product must be or can be subscribed to.
type ProductSubscription string

const (
	ProductSubscriptionAllowed  ProductSubscription = "allowed"
	ProductSubscriptionRequired ProductSubscription = "required"
)

func (r ProductSubscription) IsKnown() bool {
	switch r {
	case ProductSubscriptionAllowed, ProductSubscriptionRequired:
		return true
	}
	return false
}

// Variant of a product in the Terminal shop.
type ProductVariant struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Name of the product variant.
	Name string `json:"name,required"`
	// Price of the product variant in cents (USD).
	Price int64              `json:"price,required"`
	JSON  productVariantJSON `json:"-"`
}

// productVariantJSON contains the JSON metadata for the struct [ProductVariant]
type productVariantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productVariantJSON) RawJSON() string {
	return r.raw
}

type ProductListResponse struct {
	// A list of products.
	Data []Product               `json:"data,required"`
	JSON productListResponseJSON `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseJSON) RawJSON() string {
	return r.raw
}
