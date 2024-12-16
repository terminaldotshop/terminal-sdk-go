// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

// UserService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Update the current user.
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get initial app data, including user, products, cart, addresses, cards,
// subscriptions, and orders.
func (r *UserService) Init(ctx context.Context, opts ...option.RequestOption) (res *UserInitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the current user.
func (r *UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *UserMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A Terminal shop user. (We have users, btw.)
type User struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Email address of the user.
	Email string `json:"email,required,nullable"`
	// The user's fingerprint, derived from their public SSH key.
	Fingerprint string `json:"fingerprint,required,nullable"`
	// Name of the user.
	Name string `json:"name,required,nullable"`
	// Stripe customer ID of the user.
	StripeCustomerID string   `json:"stripeCustomerID,required"`
	JSON             userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	Fingerprint      apijson.Field
	Name             apijson.Field
	StripeCustomerID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserUpdateResponse struct {
	// A Terminal shop user. (We have users, btw.)
	Data User                   `json:"data,required"`
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
	Addresses []Address `json:"addresses,required"`
	Cards     []Card    `json:"cards,required"`
	// The current Terminal shop user's cart.
	Cart          Cart           `json:"cart,required"`
	Orders        []Order        `json:"orders,required"`
	Products      []Product      `json:"products,required"`
	Subscriptions []Subscription `json:"subscriptions,required"`
	// A Terminal shop user. (We have users, btw.)
	User User                     `json:"user,required"`
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

type UserMeResponse struct {
	// A Terminal shop user. (We have users, btw.)
	Data User               `json:"data,required"`
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
