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

// EmailService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	Options []option.RequestOption
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r *EmailService) {
	r = &EmailService{}
	r.Options = opts
	return
}

// Subscribe to email updates from Terminal.
func (r *EmailService) New(ctx context.Context, body EmailNewParams, opts ...option.RequestOption) (res *EmailNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmailNewResponse struct {
	Data EmailNewResponseData `json:"data,required"`
	JSON emailNewResponseJSON `json:"-"`
}

// emailNewResponseJSON contains the JSON metadata for the struct
// [EmailNewResponse]
type emailNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailNewResponseJSON) RawJSON() string {
	return r.raw
}

type EmailNewResponseData string

const (
	EmailNewResponseDataOk EmailNewResponseData = "ok"
)

func (r EmailNewResponseData) IsKnown() bool {
	switch r {
	case EmailNewResponseDataOk:
		return true
	}
	return false
}

type EmailNewParams struct {
	// Email address to subscribe to Terminal updates with.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r EmailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
