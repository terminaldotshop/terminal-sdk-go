// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

// ViewService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r *ViewService) {
	r = &ViewService{}
	r.Options = opts
	return
}

// Get initial app data, including user, products, cart, addresses, cards,
// subscriptions, and orders.
func (r *ViewService) Init(ctx context.Context, opts ...option.RequestOption) (res *ViewInitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "view/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ViewInitResponse struct {
	// Initial app data.
	Data ViewInitResponseData `json:"data,required"`
	JSON viewInitResponseJSON `json:"-"`
}

// viewInitResponseJSON contains the JSON metadata for the struct
// [ViewInitResponse]
type viewInitResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewInitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewInitResponseJSON) RawJSON() string {
	return r.raw
}

// Initial app data.
type ViewInitResponseData struct {
	Addresses []Address `json:"addresses,required"`
	Apps      []App     `json:"apps,required"`
	Cards     []Card    `json:"cards,required"`
	// The current Terminal shop user's cart.
	Cart     Cart      `json:"cart,required"`
	Orders   []Order   `json:"orders,required"`
	Products []Product `json:"products,required"`
	// A Terminal shop user's profile. (We have users, btw.)
	Profile       Profile                  `json:"profile,required"`
	Subscriptions []Subscription           `json:"subscriptions,required"`
	Tokens        []Token                  `json:"tokens,required"`
	JSON          viewInitResponseDataJSON `json:"-"`
}

// viewInitResponseDataJSON contains the JSON metadata for the struct
// [ViewInitResponseData]
type viewInitResponseDataJSON struct {
	Addresses     apijson.Field
	Apps          apijson.Field
	Cards         apijson.Field
	Cart          apijson.Field
	Orders        apijson.Field
	Products      apijson.Field
	Profile       apijson.Field
	Subscriptions apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ViewInitResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewInitResponseDataJSON) RawJSON() string {
	return r.raw
}
