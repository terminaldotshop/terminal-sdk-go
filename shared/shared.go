// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
)

// A physical address for shipping that sweet, sweet coffee to people's doorstep.
type Address struct {
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
	Street2 string      `json:"street2"`
	JSON    addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
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

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

// Credit card used for payments in the Terminal shop.
type Card struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Brand of the card.
	Brand string `json:"brand,required"`
	// Expiration of the card.
	Expiration CardExpiration `json:"expiration,required"`
	// Last four digits of the card.
	Last4 string   `json:"last4,required"`
	JSON  cardJSON `json:"-"`
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

// Expiration of the card.
type CardExpiration struct {
	// Expiration month of the card.
	Month int64 `json:"month,required"`
	// Expiration year of the card.
	Year int64              `json:"year,required"`
	JSON cardExpirationJSON `json:"-"`
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

// The current Terminal shop user's cart.
type Cart struct {
	// The subtotal and shipping amounts for the current user's cart.
	Amount CartAmount `json:"amount,required"`
	// An array of items in the current user's cart.
	Items []CartItem `json:"items,required"`
	// The subtotal of all items in the current user's cart.
	Subtotal float64 `json:"subtotal,required"`
	// ID of the card selected on the current user's cart.
	CardID string `json:"cardID"`
	// Shipping information for the current user's cart.
	Shipping CartShipping `json:"shipping"`
	// ID of the shipping address selected on the current user's cart.
	ShippingID string   `json:"shippingID"`
	JSON       cartJSON `json:"-"`
}

// cartJSON contains the JSON metadata for the struct [Cart]
type cartJSON struct {
	Amount      apijson.Field
	Items       apijson.Field
	Subtotal    apijson.Field
	CardID      apijson.Field
	Shipping    apijson.Field
	ShippingID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartJSON) RawJSON() string {
	return r.raw
}

// The subtotal and shipping amounts for the current user's cart.
type CartAmount struct {
	// Subtotal of the current user's cart, in cents (USD).
	Subtotal int64 `json:"subtotal,required"`
	// Shipping amount of the current user's cart, in cents (USD).
	Shipping int64          `json:"shipping"`
	JSON     cartAmountJSON `json:"-"`
}

// cartAmountJSON contains the JSON metadata for the struct [CartAmount]
type cartAmountJSON struct {
	Subtotal    apijson.Field
	Shipping    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartAmountJSON) RawJSON() string {
	return r.raw
}

// An item in the current Terminal shop user's cart.
type CartItem struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// ID of the product variant for this item in the current user's cart.
	ProductVariantID string `json:"productVariantID,required"`
	// Quantity of the item in the current user's cart.
	Quantity float64 `json:"quantity,required"`
	// Subtotal of the item in the current user's cart, in cents (USD).
	Subtotal int64        `json:"subtotal,required"`
	JSON     cartItemJSON `json:"-"`
}

// cartItemJSON contains the JSON metadata for the struct [CartItem]
type cartItemJSON struct {
	ID               apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	Subtotal         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CartItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartItemJSON) RawJSON() string {
	return r.raw
}

// Shipping information for the current user's cart.
type CartShipping struct {
	// Shipping service name.
	Service string `json:"service"`
	// Shipping timeframe provided by the shipping carrier.
	Timeframe string           `json:"timeframe"`
	JSON      cartShippingJSON `json:"-"`
}

// cartShippingJSON contains the JSON metadata for the struct [CartShipping]
type cartShippingJSON struct {
	Service     apijson.Field
	Timeframe   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartShipping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartShippingJSON) RawJSON() string {
	return r.raw
}

// An order from the Terminal shop.
type Order struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The subtotal and shipping amounts of the order.
	Amount OrderAmount `json:"amount,required"`
	// Items in the order.
	Items []OrderItem `json:"items,required"`
	// A physical address for shipping that sweet, sweet coffee to people's doorstep.
	Shipping Address `json:"shipping,required"`
	// Tracking information of the order.
	Tracking OrderTracking `json:"tracking,required"`
	// Zero-based index of the order for this user only.
	Index float64   `json:"index"`
	JSON  orderJSON `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
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
	Shipping float64 `json:"shipping,required"`
	// Subtotal amount of the order, in cents (USD).
	Subtotal float64         `json:"subtotal,required"`
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
	Amount float64 `json:"amount,required"`
	// Quantity of the item in the order.
	Quantity float64 `json:"quantity,required"`
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

// Tracking information of the order.
type OrderTracking struct {
	// Tracking number of the order.
	Number string `json:"number"`
	// Shipping service of the order.
	Service string `json:"service"`
	// Tracking URL of the order.
	URL  string            `json:"url"`
	JSON orderTrackingJSON `json:"-"`
}

// orderTrackingJSON contains the JSON metadata for the struct [OrderTracking]
type orderTrackingJSON struct {
	Number      apijson.Field
	Service     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrderTracking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderTrackingJSON) RawJSON() string {
	return r.raw
}

// Product sold in the Terminal shop.
type Product struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Description of the product.
	Description string `json:"description,required"`
	// Name of the product.
	Name string `json:"name,required"`
	// List of variants of the product.
	Variants []ProductVariant `json:"variants,required"`
	// Order of the product used when displaying a sorted list of products.
	Order int64 `json:"order"`
	// Whether the product must be or can be subscribed to.
	Subscription ProductSubscription `json:"subscription"`
	// Tags for the product.
	Tags map[string]string `json:"tags"`
	JSON productJSON       `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Variants     apijson.Field
	Order        apijson.Field
	Subscription apijson.Field
	Tags         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

// Whether the product must be or can be subscribed to.
type ProductSubscription string

const (
	ProductSubscriptionAllowed  ProductSubscription = "allowed"
	ProductSubscriptionRequired ProductSubscription = "required"
)

func (r ProductSubscription) IsKnown() bool {
	switch r {
	case ProductSubscriptionAllowed, ProductSubscriptionRequired:
		return true
	}
	return false
}

// Variant of a product in the Terminal shop.
type ProductVariant struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Name of the product variant.
	Name string `json:"name,required"`
	// Price of the product variant in cents (USD).
	Price float64            `json:"price,required"`
	JSON  productVariantJSON `json:"-"`
}

// productVariantJSON contains the JSON metadata for the struct [ProductVariant]
type productVariantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productVariantJSON) RawJSON() string {
	return r.raw
}

// Subscription to a Terminal shop product.
type Subscription struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// ID of the card used for the subscription.
	CardID string `json:"cardID,required"`
	// Frequency of the subscription.
	Frequency SubscriptionFrequency `json:"frequency,required"`
	// ID of the product variant being subscribed to.
	ProductVariantID string `json:"productVariantID,required"`
	// Quantity of the subscription.
	Quantity int64 `json:"quantity,required"`
	// ID of the shipping address used for the subscription.
	ShippingID string           `json:"shippingID,required"`
	JSON       subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID               apijson.Field
	CardID           apijson.Field
	Frequency        apijson.Field
	ProductVariantID apijson.Field
	Quantity         apijson.Field
	ShippingID       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// Frequency of the subscription.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyFixed   SubscriptionFrequency = "fixed"
	SubscriptionFrequencyDaily   SubscriptionFrequency = "daily"
	SubscriptionFrequencyWeekly  SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly SubscriptionFrequency = "monthly"
	SubscriptionFrequencyYearly  SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyFixed, SubscriptionFrequencyDaily, SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// Subscription to a Terminal shop product.
type SubscriptionParam struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID param.Field[string] `json:"id,required"`
	// ID of the card used for the subscription.
	CardID param.Field[string] `json:"cardID,required"`
	// Frequency of the subscription.
	Frequency param.Field[SubscriptionFrequency] `json:"frequency,required"`
	// ID of the product variant being subscribed to.
	ProductVariantID param.Field[string] `json:"productVariantID,required"`
	// Quantity of the subscription.
	Quantity param.Field[int64] `json:"quantity,required"`
	// ID of the shipping address used for the subscription.
	ShippingID param.Field[string] `json:"shippingID,required"`
}

func (r SubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Terminal shop user. (We have users, btw.)
type User struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Email address of the user.
	Email string `json:"email,required,nullable"`
	// The user's fingerprint, derived from their public SSH key.
	Fingerprint string `json:"fingerprint,required,nullable"`
	// Name of the user.
	Name string `json:"name,required,nullable"`
	// Stripe customer ID of the user.
	StripeCustomerID string   `json:"stripeCustomerID,required"`
	JSON             userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID               apijson.Field
	Email            apijson.Field
	Fingerprint      apijson.Field
	Name             apijson.Field
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
