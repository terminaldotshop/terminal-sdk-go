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

// Create an order without a cart. The order will be placed immediately.
func (r *OrderService) New(ctx context.Context, body OrderNewParams, opts ...option.RequestOption) (res *OrderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// An order from the Terminal shop.
type Order struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The subtotal and shipping amounts of the order.
	Amount OrderAmount `json:"amount,required"`
	// Date the order was created.
	Created string `json:"created,required"`
	// Items in the order.
	Items []OrderItem `json:"items,required"`
	// Shipping address of the order.
	Shipping OrderShipping `json:"shipping,required"`
	// Tracking information of the order.
	Tracking OrderTracking `json:"tracking,required"`
	// Zero-based index of the order for this user only.
	Index int64     `json:"index"`
	JSON  orderJSON `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Created     apijson.Field
	Items       apijson.Field
	Shipping    apijson.Field
	Tracking    apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

// The subtotal and shipping amounts of the order.
type OrderAmount struct {
	// Shipping amount of the order, in cents (USD).
	Shipping int64 `json:"shipping,required"`
	// Subtotal amount of the order, in cents (USD).
	Subtotal int64           `json:"subtotal,required"`
	JSON     orderAmountJSON `json:"-"`
}

// orderAmountJSON contains the JSON metadata for the struct [OrderAmount]
type orderAmountJSON struct {
	Shipping    apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderAmountJSON) RawJSON() string {
	return r.raw
}

type OrderItem struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Amount of the item in the order, in cents (USD).
	Amount int64 `json:"amount,required"`
	// Quantity of the item in the order.
	Quantity int64 `json:"quantity,required"`
	// Description of the item in the order.
	Description string `json:"description"`
	// ID of the product variant of the item in the order.
	ProductVariantID string        `json:"productVariantID"`
	JSON             orderItemJSON `json:"-"`
}

// orderItemJSON contains the JSON metadata for the struct [OrderItem]
type orderItemJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Quantity         apijson.Field
	Description      apijson.Field
	ProductVariantID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrderItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderItemJSON) RawJSON() string {
	return r.raw
}

// Shipping address of the order.
type OrderShipping struct {
	// City of the address.
	City string `json:"city,required"`
	// ISO 3166-1 alpha-2 country code of the address.
	Country string `json:"country,required"`
	// The recipient's name.
	Name string `json:"name,required"`
	// Street of the address.
	Street1 string `json:"street1,required"`
	// Zip code of the address.
	Zip string `json:"zip,required"`
	// Phone number of the recipient.
	Phone string `json:"phone"`
	// Province or state of the address.
	Province string `json:"province"`
	// Apartment, suite, etc. of the address.
	Street2 string            `json:"street2"`
	JSON    orderShippingJSON `json:"-"`
}

// orderShippingJSON contains the JSON metadata for the struct [OrderShipping]
type orderShippingJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
	Phone       apijson.Field
	Province    apijson.Field
	Street2     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderShipping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderShippingJSON) RawJSON() string {
	return r.raw
}

// Tracking information of the order.
type OrderTracking struct {
	// Tracking number of the order.
	Number string `json:"number"`
	// Shipping service of the order.
	Service string `json:"service"`
	// Current tracking status of the shipment.
	Status OrderTrackingStatus `json:"status"`
	// Additional details about the tracking status.
	StatusDetails string `json:"statusDetails"`
	// When the tracking status was last updated.
	StatusUpdatedAt string `json:"statusUpdatedAt"`
	// Tracking URL of the order.
	URL  string            `json:"url"`
	JSON orderTrackingJSON `json:"-"`
}

// orderTrackingJSON contains the JSON metadata for the struct [OrderTracking]
type orderTrackingJSON struct {
	Number          apijson.Field
	Service         apijson.Field
	Status          apijson.Field
	StatusDetails   apijson.Field
	StatusUpdatedAt apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OrderTracking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderTrackingJSON) RawJSON() string {
	return r.raw
}

// Current tracking status of the shipment.
type OrderTrackingStatus string

const (
	OrderTrackingStatusPreTransit OrderTrackingStatus = "PRE_TRANSIT"
	OrderTrackingStatusTransit    OrderTrackingStatus = "TRANSIT"
	OrderTrackingStatusDelivered  OrderTrackingStatus = "DELIVERED"
	OrderTrackingStatusReturned   OrderTrackingStatus = "RETURNED"
	OrderTrackingStatusFailure    OrderTrackingStatus = "FAILURE"
	OrderTrackingStatusUnknown    OrderTrackingStatus = "UNKNOWN"
)

func (r OrderTrackingStatus) IsKnown() bool {
	switch r {
	case OrderTrackingStatusPreTransit, OrderTrackingStatusTransit, OrderTrackingStatusDelivered, OrderTrackingStatusReturned, OrderTrackingStatusFailure, OrderTrackingStatusUnknown:
		return true
	}
	return false
}

type OrderNewResponse struct {
	// Order ID.
	Data string               `json:"data,required"`
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
	Data []Order               `json:"data,required"`
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
	Data Order                `json:"data,required"`
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

type OrderNewParams struct {
	// Shipping address ID.
	AddressID param.Field[string] `json:"addressID,required"`
	// Card ID.
	CardID param.Field[string] `json:"cardID,required"`
	// Product variants to include in the order, along with their quantities.
	Variants param.Field[map[string]int64] `json:"variants,required"`
}

func (r OrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
