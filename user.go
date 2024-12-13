// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/terminal-go/internal/apijson"
	"github.com/stainless-sdks/terminal-go/internal/param"
	"github.com/stainless-sdks/terminal-go/internal/requestconfig"
	"github.com/stainless-sdks/terminal-go/option"
	"github.com/stainless-sdks/terminal-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options  []option.RequestOption
	Shipping *UserShippingService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Shipping = NewUserShippingService(opts...)
	return
}

// Update the current user.
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get initial app data, including user, products, cart, addresses, cards,
// subscriptions, and orders.
func (r *UserService) Init(ctx context.Context, opts ...option.RequestOption) (res *UserInitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the current user.
func (r *UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *UserMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserUpdateResponse struct {
	// A Terminal shop user. (We have users, btw.)
	Data shared.User            `json:"data,required"`
	JSON userUpdateResponseJSON `json:"-"`
}

// userUpdateResponseJSON contains the JSON metadata for the struct
// [UserUpdateResponse]
type userUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type UserInitResponse struct {
	// Initial app data.
	Data UserInitResponseData `json:"data,required"`
	JSON userInitResponseJSON `json:"-"`
}

// userInitResponseJSON contains the JSON metadata for the struct
// [UserInitResponse]
type userInitResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInitResponseJSON) RawJSON() string {
	return r.raw
}

// Initial app data.
type UserInitResponseData struct {
	Addresses []UserInitResponseDataAddress `json:"addresses,required"`
	Cards     []shared.Card                 `json:"cards,required"`
	// The current Terminal shop user's cart.
	Cart          shared.Cart           `json:"cart,required"`
	Orders        []shared.Order        `json:"orders,required"`
	Products      []shared.Product      `json:"products,required"`
	Subscriptions []shared.Subscription `json:"subscriptions,required"`
	// A Terminal shop user. (We have users, btw.)
	User shared.User              `json:"user,required"`
	JSON userInitResponseDataJSON `json:"-"`
}

// userInitResponseDataJSON contains the JSON metadata for the struct
// [UserInitResponseData]
type userInitResponseDataJSON struct {
	Addresses     apijson.Field
	Cards         apijson.Field
	Cart          apijson.Field
	Orders        apijson.Field
	Products      apijson.Field
	Subscriptions apijson.Field
	User          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserInitResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInitResponseDataJSON) RawJSON() string {
	return r.raw
}

// Shipping address associated with a Terminal shop user.
type UserInitResponseDataAddress struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A physical address for shipping that sweet, sweet coffee to people's doorstep.
	Address shared.Address                  `json:"address,required"`
	JSON    userInitResponseDataAddressJSON `json:"-"`
}

// userInitResponseDataAddressJSON contains the JSON metadata for the struct
// [UserInitResponseDataAddress]
type userInitResponseDataAddressJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInitResponseDataAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInitResponseDataAddressJSON) RawJSON() string {
	return r.raw
}

type UserMeResponse struct {
	// A Terminal shop user. (We have users, btw.)
	Data shared.User        `json:"data,required"`
	JSON userMeResponseJSON `json:"-"`
}

// userMeResponseJSON contains the JSON metadata for the struct [UserMeResponse]
type userMeResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseJSON) RawJSON() string {
	return r.raw
}

type UserUpdateParams struct {
	// Email address of the user.
	Email param.Field[string] `json:"email"`
	// Name of the user.
	Name param.Field[string] `json:"name"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
