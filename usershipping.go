// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/terminal-go/internal/apijson"
	"github.com/stainless-sdks/terminal-go/internal/param"
	"github.com/stainless-sdks/terminal-go/internal/requestconfig"
	"github.com/stainless-sdks/terminal-go/option"
	"github.com/stainless-sdks/terminal-go/shared"
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

// Create and add a shipping address to the current user.
func (r *UserShippingService) New(ctx context.Context, body UserShippingNewParams, opts ...option.RequestOption) (res *UserShippingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the shipping addresses associated with the current user.
func (r *UserShippingService) List(ctx context.Context, opts ...option.RequestOption) (res *UserShippingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/shipping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a shipping address from the current user.
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
	// Shipping address ID.
	Data string                      `json:"data,required"`
	JSON userShippingNewResponseJSON `json:"-"`
}

// userShippingNewResponseJSON contains the JSON metadata for the struct
// [UserShippingNewResponse]
type userShippingNewResponseJSON struct {
	Data        apijson.Field
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
	// Shipping addresses.
	Data []UserShippingListResponseData `json:"data,required"`
	JSON userShippingListResponseJSON   `json:"-"`
}

// userShippingListResponseJSON contains the JSON metadata for the struct
// [UserShippingListResponse]
type userShippingListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingListResponseJSON) RawJSON() string {
	return r.raw
}

// Shipping address associated with a Terminal shop user.
type UserShippingListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A physical address for shipping that sweet, sweet coffee to people's doorstep.
	Address shared.Address                   `json:"address,required"`
	JSON    userShippingListResponseDataJSON `json:"-"`
}

// userShippingListResponseDataJSON contains the JSON metadata for the struct
// [UserShippingListResponseData]
type userShippingListResponseDataJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingListResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserShippingDeleteResponse struct {
	Data UserShippingDeleteResponseData `json:"data,required"`
	JSON userShippingDeleteResponseJSON `json:"-"`
}

// userShippingDeleteResponseJSON contains the JSON metadata for the struct
// [UserShippingDeleteResponse]
type userShippingDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserShippingDeleteResponseData string

const (
	UserShippingDeleteResponseDataOk UserShippingDeleteResponseData = "ok"
)

func (r UserShippingDeleteResponseData) IsKnown() bool {
	switch r {
	case UserShippingDeleteResponseDataOk:
		return true
	}
	return false
}

type UserShippingNewParams struct {
	// City of the address.
	City    param.Field[string] `json:"city,required"`
	Country param.Field[string] `json:"country,required"`
	// The recipient's name.
	Name param.Field[string] `json:"name,required"`
	// Street of the address.
	Street1 param.Field[string] `json:"street1,required"`
	// Zip code of the address.
	Zip param.Field[string] `json:"zip,required"`
	// Phone number of the recipient.
	Phone param.Field[string] `json:"phone"`
	// Province or state of the address.
	Province param.Field[string] `json:"province"`
	// Apartment, suite, etc. of the address.
	Street2 param.Field[string] `json:"street2"`
}

func (r UserShippingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
