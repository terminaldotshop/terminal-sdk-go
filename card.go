// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
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

func (r *CardService) List(ctx context.Context, opts ...option.RequestOption) (res *CardListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Card struct {
	ID         string         `json:"id,required"`
	Brand      string         `json:"brand,required"`
	Expiration CardExpiration `json:"expiration,required"`
	Last4      string         `json:"last4,required"`
	JSON       cardJSON       `json:"-"`
}

// cardJSON contains the JSON metadata for the struct [Card]
type cardJSON struct {
	ID          apijson.Field
	Brand       apijson.Field
	Expiration  apijson.Field
	Last4       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardJSON) RawJSON() string {
	return r.raw
}

type CardExpiration struct {
	Month float64            `json:"month,required"`
	Year  float64            `json:"year,required"`
	JSON  cardExpirationJSON `json:"-"`
}

// cardExpirationJSON contains the JSON metadata for the struct [CardExpiration]
type cardExpirationJSON struct {
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardExpirationJSON) RawJSON() string {
	return r.raw
}

type CardListResponse struct {
	Result []Card               `json:"result,required"`
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
