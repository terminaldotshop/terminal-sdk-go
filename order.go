// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/requestconfig"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terminal-sdk-go/shared"
)

// OrderService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options []option.RequestOption
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r *OrderService) {
	r = &OrderService{}
	r.Options = opts
	return
}

// Create an order from the current user's cart.
func (r *OrderService) New(ctx context.Context, opts ...option.RequestOption) (res *OrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List the orders associated with the current user.
func (r *OrderService) List(ctx context.Context, opts ...option.RequestOption) (res *OrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the order with the given ID.
func (r *OrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OrderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("order/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrderNewResponse struct {
	// An order from the Terminal shop.
	Data shared.Order         `json:"data,required"`
	JSON orderNewResponseJSON `json:"-"`
}

// orderNewResponseJSON contains the JSON metadata for the struct
// [OrderNewResponse]
type orderNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrderListResponse struct {
	// List of orders.
	Data []shared.Order        `json:"data,required"`
	JSON orderListResponseJSON `json:"-"`
}

// orderListResponseJSON contains the JSON metadata for the struct
// [OrderListResponse]
type orderListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderListResponseJSON) RawJSON() string {
	return r.raw
}

type OrderGetResponse struct {
	// An order from the Terminal shop.
	Data shared.Order         `json:"data,required"`
	JSON orderGetResponseJSON `json:"-"`
}

// orderGetResponseJSON contains the JSON metadata for the struct
// [OrderGetResponse]
type orderGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderGetResponseJSON) RawJSON() string {
	return r.raw
}
