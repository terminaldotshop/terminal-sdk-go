// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"errors"
	"fmt"
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

// Attach a credit card (tokenized via Stripe) to the current user.
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *CardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the credit cards associated with the current user.
func (r *CardService) List(ctx context.Context, opts ...option.RequestOption) (res *CardListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a credit card associated with the current user.
func (r *CardService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CardDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("card/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CardNewResponse struct {
	// ID of the card.
	Data string              `json:"data,required"`
	JSON cardNewResponseJSON `json:"-"`
}

// cardNewResponseJSON contains the JSON metadata for the struct [CardNewResponse]
type cardNewResponseJSON struct {
	Data        apijson.Field
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
	// List of cards associated with the user.
	Data []shared.Card        `json:"data,required"`
	JSON cardListResponseJSON `json:"-"`
}

// cardListResponseJSON contains the JSON metadata for the struct
// [CardListResponse]
type cardListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardListResponseJSON) RawJSON() string {
	return r.raw
}

type CardDeleteResponse struct {
	Data CardDeleteResponseData `json:"data,required"`
	JSON cardDeleteResponseJSON `json:"-"`
}

// cardDeleteResponseJSON contains the JSON metadata for the struct
// [CardDeleteResponse]
type cardDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CardDeleteResponseData string

const (
	CardDeleteResponseDataOk CardDeleteResponseData = "ok"
)

func (r CardDeleteResponseData) IsKnown() bool {
	switch r {
	case CardDeleteResponseDataOk:
		return true
	}
	return false
}

type CardNewParams struct {
	// Stripe card token. Learn how to
	// [create one here](https://docs.stripe.com/api/tokens/create_card).
	Token param.Field[string] `json:"token,required"`
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
