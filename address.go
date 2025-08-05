// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

// AddressService contains methods and other services that help with interacting
// with the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r *AddressService) {
	r = &AddressService{}
	r.Options = opts
	return
}

// Create and add a shipping address to the current user.
func (r *AddressService) New(ctx context.Context, body AddressNewParams, opts ...option.RequestOption) (res *AddressNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the shipping addresses associated with the current user.
func (r *AddressService) List(ctx context.Context, opts ...option.RequestOption) (res *AddressListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "address"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a shipping address from the current user.
func (r *AddressService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("address/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the shipping address with the given ID.
func (r *AddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("address/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Physical address associated with a Terminal shop user.
type Address struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// City of the address.
	City string `json:"city,required"`
	// ISO 3166-1 alpha-2 country code of the address.
	Country string `json:"country,required"`
	// Date the address was created.
	Created string `json:"created,required"`
	// The recipient's name.
	Name string `json:"name,required"`
	// Street of the address.
	Street1 string `json:"street1,required"`
	// Zip code of the address.
	Zip string `json:"zip,required"`
	// Phone number of the recipient.
	Phone string `json:"phone"`
	// Province or state of the address.
	Province string `json:"province"`
	// Apartment, suite, etc. of the address.
	Street2 string      `json:"street2"`
	JSON    addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	ID          apijson.Field
	City        apijson.Field
	Country     apijson.Field
	Created     apijson.Field
	Name        apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
	Phone       apijson.Field
	Province    apijson.Field
	Street2     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressNewResponse struct {
	// Shipping address ID.
	Data string                 `json:"data,required"`
	JSON addressNewResponseJSON `json:"-"`
}

// addressNewResponseJSON contains the JSON metadata for the struct
// [AddressNewResponse]
type addressNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseJSON) RawJSON() string {
	return r.raw
}

type AddressListResponse struct {
	// Shipping addresses.
	Data []Address               `json:"data,required"`
	JSON addressListResponseJSON `json:"-"`
}

// addressListResponseJSON contains the JSON metadata for the struct
// [AddressListResponse]
type addressListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressListResponseJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponse struct {
	Data AddressDeleteResponseData `json:"data,required"`
	JSON addressDeleteResponseJSON `json:"-"`
}

// addressDeleteResponseJSON contains the JSON metadata for the struct
// [AddressDeleteResponse]
type addressDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponseData string

const (
	AddressDeleteResponseDataOk AddressDeleteResponseData = "ok"
)

func (r AddressDeleteResponseData) IsKnown() bool {
	switch r {
	case AddressDeleteResponseDataOk:
		return true
	}
	return false
}

type AddressGetResponse struct {
	// Physical address associated with a Terminal shop user.
	Data Address                `json:"data,required"`
	JSON addressGetResponseJSON `json:"-"`
}

// addressGetResponseJSON contains the JSON metadata for the struct
// [AddressGetResponse]
type addressGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseJSON) RawJSON() string {
	return r.raw
}

type AddressNewParams struct {
	// City of the address.
	City param.Field[string] `json:"city,required"`
	// ISO 3166-1 alpha-2 country code of the address.
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

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
