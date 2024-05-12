// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/terminal-go/internal/apijson"
	"github.com/stainless-sdks/terminal-go/internal/requestconfig"
	"github.com/stainless-sdks/terminal-go/option"
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

func (r *ProductService) Get(ctx context.Context, opts ...option.RequestOption) (res *ProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProductGetResponse struct {
	Result []ProductGetResponseResult `json:"result,required"`
	JSON   productGetResponseJSON     `json:"-"`
}

// productGetResponseJSON contains the JSON metadata for the struct
// [ProductGetResponse]
type productGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProductGetResponseResult struct {
	ID          string                            `json:"id,required"`
	Description string                            `json:"description,required"`
	Name        string                            `json:"name,required"`
	Variants    []ProductGetResponseResultVariant `json:"variants,required"`
	JSON        productGetResponseResultJSON      `json:"-"`
}

// productGetResponseResultJSON contains the JSON metadata for the struct
// [ProductGetResponseResult]
type productGetResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type ProductGetResponseResultVariant struct {
	ID        string                              `json:"id,required"`
	Name      string                              `json:"name,required"`
	Price     int64                               `json:"price,required"`
	ProductID string                              `json:"productID,required"`
	JSON      productGetResponseResultVariantJSON `json:"-"`
}

// productGetResponseResultVariantJSON contains the JSON metadata for the struct
// [ProductGetResponseResultVariant]
type productGetResponseResultVariantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	ProductID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductGetResponseResultVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponseResultVariantJSON) RawJSON() string {
	return r.raw
}
