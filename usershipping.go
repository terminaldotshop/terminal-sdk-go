// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
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

func (r *UserShippingService) Delete(ctx context.Context, id string, body UserShippingDeleteParams, opts ...option.RequestOption) (res *UserShippingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/shipping/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type UserShippingNewResponse struct {
	Result []UserShippingNewResponseResult `json:"result,required"`
	JSON   userShippingNewResponseJSON     `json:"-"`
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

type UserShippingNewResponseResult struct {
	ID      string                               `json:"id,required"`
	Address UserShippingNewResponseResultAddress `json:"address,required"`
	JSON    userShippingNewResponseResultJSON    `json:"-"`
}

// userShippingNewResponseResultJSON contains the JSON metadata for the struct
// [UserShippingNewResponseResult]
type userShippingNewResponseResultJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserShippingNewResponseResultAddress struct {
	City     string                                   `json:"city,required"`
	Country  string                                   `json:"country,required"`
	Name     string                                   `json:"name,required"`
	Province string                                   `json:"province,required"`
	Street1  string                                   `json:"street1,required"`
	Zip      string                                   `json:"zip,required"`
	Street2  string                                   `json:"street2"`
	JSON     userShippingNewResponseResultAddressJSON `json:"-"`
}

// userShippingNewResponseResultAddressJSON contains the JSON metadata for the
// struct [UserShippingNewResponseResultAddress]
type userShippingNewResponseResultAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Province    apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
	Street2     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingNewResponseResultAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingNewResponseResultAddressJSON) RawJSON() string {
	return r.raw
}

type UserShippingListResponse struct {
	Result []UserShippingListResponseResult `json:"result,required"`
	JSON   userShippingListResponseJSON     `json:"-"`
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

type UserShippingListResponseResult struct {
	ID      string                                `json:"id,required"`
	Address UserShippingListResponseResultAddress `json:"address,required"`
	JSON    userShippingListResponseResultJSON    `json:"-"`
}

// userShippingListResponseResultJSON contains the JSON metadata for the struct
// [UserShippingListResponseResult]
type userShippingListResponseResultJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingListResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserShippingListResponseResultAddress struct {
	City     string                                    `json:"city,required"`
	Country  string                                    `json:"country,required"`
	Name     string                                    `json:"name,required"`
	Province string                                    `json:"province,required"`
	Street1  string                                    `json:"street1,required"`
	Zip      string                                    `json:"zip,required"`
	Street2  string                                    `json:"street2"`
	JSON     userShippingListResponseResultAddressJSON `json:"-"`
}

// userShippingListResponseResultAddressJSON contains the JSON metadata for the
// struct [UserShippingListResponseResultAddress]
type userShippingListResponseResultAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Province    apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
	Street2     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingListResponseResultAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingListResponseResultAddressJSON) RawJSON() string {
	return r.raw
}

type UserShippingDeleteResponse struct {
	Result []UserShippingDeleteResponseResult `json:"result,required"`
	JSON   userShippingDeleteResponseJSON     `json:"-"`
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

type UserShippingDeleteResponseResult struct {
	ID      string                                  `json:"id,required"`
	Address UserShippingDeleteResponseResultAddress `json:"address,required"`
	JSON    userShippingDeleteResponseResultJSON    `json:"-"`
}

// userShippingDeleteResponseResultJSON contains the JSON metadata for the struct
// [UserShippingDeleteResponseResult]
type userShippingDeleteResponseResultJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserShippingDeleteResponseResultAddress struct {
	City     string                                      `json:"city,required"`
	Country  string                                      `json:"country,required"`
	Name     string                                      `json:"name,required"`
	Province string                                      `json:"province,required"`
	Street1  string                                      `json:"street1,required"`
	Zip      string                                      `json:"zip,required"`
	Street2  string                                      `json:"street2"`
	JSON     userShippingDeleteResponseResultAddressJSON `json:"-"`
}

// userShippingDeleteResponseResultAddressJSON contains the JSON metadata for the
// struct [UserShippingDeleteResponseResultAddress]
type userShippingDeleteResponseResultAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Province    apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
	Street2     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserShippingDeleteResponseResultAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userShippingDeleteResponseResultAddressJSON) RawJSON() string {
	return r.raw
}

type UserShippingNewParams struct {
	City     param.Field[string] `json:"city,required"`
	Country  param.Field[string] `json:"country,required"`
	Name     param.Field[string] `json:"name,required"`
	Province param.Field[string] `json:"province,required"`
	Street1  param.Field[string] `json:"street1,required"`
	Zip      param.Field[string] `json:"zip,required"`
	Street2  param.Field[string] `json:"street2"`
}

func (r UserShippingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserShippingDeleteParams struct {
	City     param.Field[string] `json:"city,required"`
	Country  param.Field[string] `json:"country,required"`
	Name     param.Field[string] `json:"name,required"`
	Province param.Field[string] `json:"province,required"`
	Street1  param.Field[string] `json:"street1,required"`
	Zip      param.Field[string] `json:"zip,required"`
	Street2  param.Field[string] `json:"street2"`
}

func (r UserShippingDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
