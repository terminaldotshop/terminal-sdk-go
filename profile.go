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

// ProfileService contains methods and other services that help with interacting
// with the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r *ProfileService) {
	r = &ProfileService{}
	r.Options = opts
	return
}

// Update the current user's profile.
func (r *ProfileService) Update(ctx context.Context, body ProfileUpdateParams, opts ...option.RequestOption) (res *ProfileUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get the current user's profile.
func (r *ProfileService) Me(ctx context.Context, opts ...option.RequestOption) (res *ProfileMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A Terminal shop user's profile. (We have users, btw.)
type Profile struct {
	// A Terminal shop user. (We have users, btw.)
	User ProfileUser `json:"user,required"`
	JSON profileJSON `json:"-"`
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

// A Terminal shop user. (We have users, btw.)
type ProfileUser struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Email address of the user.
	Email string `json:"email,required,nullable"`
	// The user's fingerprint, derived from their public SSH key.
	Fingerprint string `json:"fingerprint,required,nullable"`
	// Name of the user.
	Name string `json:"name,required,nullable"`
	// Stripe customer ID of the user.
	StripeCustomerID string          `json:"stripeCustomerID,required"`
	JSON             profileUserJSON `json:"-"`
}

// profileUserJSON contains the JSON metadata for the struct [ProfileUser]
type profileUserJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	Fingerprint      apijson.Field
	Name             apijson.Field
	StripeCustomerID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProfileUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUserJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponse struct {
	// A Terminal shop user's profile. (We have users, btw.)
	Data Profile                   `json:"data,required"`
	JSON profileUpdateResponseJSON `json:"-"`
}

// profileUpdateResponseJSON contains the JSON metadata for the struct
// [ProfileUpdateResponse]
type profileUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileMeResponse struct {
	// A Terminal shop user's profile. (We have users, btw.)
	Data Profile               `json:"data,required"`
	JSON profileMeResponseJSON `json:"-"`
}

// profileMeResponseJSON contains the JSON metadata for the struct
// [ProfileMeResponse]
type profileMeResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileMeResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateParams struct {
	Email param.Field[string] `json:"email,required" format:"email"`
	Name  param.Field[string] `json:"name,required"`
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
