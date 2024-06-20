// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
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

func (r *UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *UserMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type User struct {
	ID               string   `json:"id,required"`
	Email            string   `json:"email,required,nullable"`
	Fingerprint      string   `json:"fingerprint,required,nullable"`
	StripeCustomerID string   `json:"stripeCustomerID,required"`
	JSON             userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	Fingerprint      apijson.Field
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

type UserMeResponse struct {
	Result User               `json:"result,required"`
	JSON   userMeResponseJSON `json:"-"`
}

// userMeResponseJSON contains the JSON metadata for the struct [UserMeResponse]
type userMeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseJSON) RawJSON() string {
	return r.raw
}
