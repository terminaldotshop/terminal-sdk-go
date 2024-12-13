// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/terminal-go/internal/apijson"
	"github.com/stainless-sdks/terminal-go/internal/requestconfig"
	"github.com/stainless-sdks/terminal-go/option"
	"github.com/stainless-sdks/terminal-go/shared"
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
	path := "product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProductListResponse struct {
	// A list of products.
	Data []shared.Product        `json:"data,required"`
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
