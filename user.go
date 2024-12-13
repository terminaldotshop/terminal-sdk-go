// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terminal-sdk-go/shared"
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
	path := "user/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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
