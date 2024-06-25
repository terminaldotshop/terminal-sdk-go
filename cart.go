// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
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
