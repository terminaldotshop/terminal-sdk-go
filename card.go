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

// CardService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	Options []option.RequestOption
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r *CardService) {
	r = &CardService{}
	r.Options = opts
	return
}

func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *CardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CardService) List(ctx context.Context, opts ...option.RequestOption) (res *CardListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CardNewResponse struct {
	Result string              `json:"result,required"`
	JSON   cardNewResponseJSON `json:"-"`
}

// cardNewResponseJSON contains the JSON metadata for the struct [CardNewResponse]
type cardNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardNewResponseJSON) RawJSON() string {
	return r.raw
}

type CardListResponse struct {
	Result []shared.Card        `json:"result,required"`
	JSON   cardListResponseJSON `json:"-"`
}

// cardListResponseJSON contains the JSON metadata for the struct
// [CardListResponse]
type cardListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardListResponseJSON) RawJSON() string {
	return r.raw
}

type CardNewParams struct {
	Token param.Field[string] `json:"token,required"`
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
