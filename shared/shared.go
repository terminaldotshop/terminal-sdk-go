// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/terminaldotshop/terminal-sdk-go/internal/apijson"
	"github.com/terminaldotshop/terminal-sdk-go/internal/param"
)

type Address struct {
	City     string      `json:"city,required"`
	Country  string      `json:"country,required"`
	Name     string      `json:"name,required"`
	Province string      `json:"province,required"`
	Street1  string      `json:"street1,required"`
	Zip      string      `json:"zip,required"`
	Street2  string      `json:"street2"`
	JSON     addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Province    apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
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

type AddressParam struct {
	City     param.Field[string] `json:"city,required"`
	Country  param.Field[string] `json:"country,required"`
	Name     param.Field[string] `json:"name,required"`
	Province param.Field[string] `json:"province,required"`
	Street1  param.Field[string] `json:"street1,required"`
	Zip      param.Field[string] `json:"zip,required"`
	Street2  param.Field[string] `json:"street2"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Month int64              `json:"month,required"`
	Year  int64              `json:"year,required"`
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

type Cart struct {
	Amount     CartAmount `json:"amount,required"`
	Items      []CartItem `json:"items,required"`
	Subtotal   int64      `json:"subtotal,required"`
	CardID     string     `json:"cardID"`
	ShippingID string     `json:"shippingID"`
	JSON       cartJSON   `json:"-"`
}

// cartJSON contains the JSON metadata for the struct [Cart]
type cartJSON struct {
	Amount      apijson.Field
	Items       apijson.Field
	Subtotal    apijson.Field
	CardID      apijson.Field
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

type CartAmount struct {
	Shipping int64          `json:"shipping,required"`
	Subtotal int64          `json:"subtotal,required"`
	JSON     cartAmountJSON `json:"-"`
}

// cartAmountJSON contains the JSON metadata for the struct [CartAmount]
type cartAmountJSON struct {
	Shipping    apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CartAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cartAmountJSON) RawJSON() string {
	return r.raw
}

type CartItem struct {
	ID               string       `json:"id,required"`
	ProductVariantID string       `json:"productVariantID,required"`
	Quantity         int64        `json:"quantity,required"`
	Subtotal         int64        `json:"subtotal,required"`
	JSON             cartItemJSON `json:"-"`
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

type Order struct {
	ID       string        `json:"id,required"`
	Amount   OrderAmount   `json:"amount,required"`
	Items    []OrderItem   `json:"items,required"`
	Shipping OrderShipping `json:"shipping,required"`
	Tracking OrderTracking `json:"tracking,required"`
	JSON     orderJSON     `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Items       apijson.Field
	Shipping    apijson.Field
	Tracking    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

type OrderAmount struct {
	Shipping float64         `json:"shipping,required"`
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
	ID               string        `json:"id,required"`
	Amount           float64       `json:"amount,required"`
	Quantity         int64         `json:"quantity,required"`
	Description      string        `json:"description"`
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

type OrderShipping struct {
	City     string            `json:"city,required"`
	Country  string            `json:"country,required"`
	Name     string            `json:"name,required"`
	Province string            `json:"province,required"`
	Street1  string            `json:"street1,required"`
	Zip      string            `json:"zip,required"`
	Street2  string            `json:"street2"`
	JSON     orderShippingJSON `json:"-"`
}

// orderShippingJSON contains the JSON metadata for the struct [OrderShipping]
type orderShippingJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Province    apijson.Field
	Street1     apijson.Field
	Zip         apijson.Field
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

type OrderTracking struct {
	Number string            `json:"number"`
	URL    string            `json:"url"`
	JSON   orderTrackingJSON `json:"-"`
}

// orderTrackingJSON contains the JSON metadata for the struct [OrderTracking]
type orderTrackingJSON struct {
	Number      apijson.Field
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

type Product struct {
	ID          string           `json:"id,required"`
	Description string           `json:"description,required"`
	Name        string           `json:"name,required"`
	Variants    []ProductVariant `json:"variants,required"`
	JSON        productJSON      `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

type ProductVariant struct {
	ID    string             `json:"id,required"`
	Name  string             `json:"name,required"`
	Price int64              `json:"price,required"`
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

type Shipping struct {
	ID      string       `json:"id,required"`
	Address Address      `json:"address,required"`
	JSON    shippingJSON `json:"-"`
}

// shippingJSON contains the JSON metadata for the struct [Shipping]
type shippingJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Shipping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shippingJSON) RawJSON() string {
	return r.raw
}

type User struct {
	ID               string   `json:"id,required"`
	Email            string   `json:"email,required,nullable"`
	Fingerprint      string   `json:"fingerprint,required,nullable"`
	Name             string   `json:"name,required,nullable"`
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
