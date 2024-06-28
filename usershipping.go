// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terminal-sdk-go/shared"
)

// UserShippingService contains methods and other services that help with
// interacting with the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserShippingService] method instead.
type UserShippingService struct {
	Options []option.RequestOption
}

// NewUserShippingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserShippingService(opts ...option.RequestOption) (r *UserShippingService) {
	r = &UserShippingService{}
	r.Options = opts
	return
}

func (r *UserShippingService) New(ctx context.Context, body UserShippingNewParams, opts ...option.RequestOption) (res *UserShippingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *UserShippingService) List(ctx context.Context, opts ...option.RequestOption) (res *UserShippingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *UserShippingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *UserShippingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("user/shipping/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type UserShippingNewResponse struct {
	Result []shared.Shipping           `json:"result,required"`
	JSON   userShippingNewResponseJSON `json:"-"`
}

// userShippingNewResponseJSON contains the JSON metadata for the struct
// [UserShippingNewResponse]
type userShippingNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingNewResponseJSON) RawJSON() string {
	return r.raw
}

type UserShippingListResponse struct {
	Result []shared.Shipping            `json:"result,required"`
	JSON   userShippingListResponseJSON `json:"-"`
}

// userShippingListResponseJSON contains the JSON metadata for the struct
// [UserShippingListResponse]
type userShippingListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingListResponseJSON) RawJSON() string {
	return r.raw
}

type UserShippingDeleteResponse struct {
	Result []shared.Shipping              `json:"result,required"`
	JSON   userShippingDeleteResponseJSON `json:"-"`
}

// userShippingDeleteResponseJSON contains the JSON metadata for the struct
// [UserShippingDeleteResponse]
type userShippingDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserShippingNewParams struct {
	Address shared.AddressParam `json:"address,required"`
}

func (r UserShippingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Address)
}
