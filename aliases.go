// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package terminal

import (
	"github.com/terminaldotshop/terminal-sdk-go/internal/apierror"
	"github.com/terminaldotshop/terminal-sdk-go/shared"
)

type Error = apierror.Error

// Physical address associated with a Terminal shop user.
//
// This is an alias to an internal type.
type Address = shared.Address

// Credit card used for payments in the Terminal shop.
//
// This is an alias to an internal type.
type Card = shared.Card

// Expiration of the card.
//
// This is an alias to an internal type.
type CardExpiration = shared.CardExpiration

// The current Terminal shop user's cart.
//
// This is an alias to an internal type.
type Cart = shared.Cart

// The subtotal and shipping amounts for the current user's cart.
//
// This is an alias to an internal type.
type CartAmount = shared.CartAmount

// An item in the current Terminal shop user's cart.
//
// This is an alias to an internal type.
type CartItem = shared.CartItem

// Shipping information for the current user's cart.
//
// This is an alias to an internal type.
type CartShipping = shared.CartShipping

// An order from the Terminal shop.
//
// This is an alias to an internal type.
type Order = shared.Order

// The subtotal and shipping amounts of the order.
//
// This is an alias to an internal type.
type OrderAmount = shared.OrderAmount

// This is an alias to an internal type.
type OrderItem = shared.OrderItem

// Shipping address of the order.
//
// This is an alias to an internal type.
type OrderShipping = shared.OrderShipping

// Tracking information of the order.
//
// This is an alias to an internal type.
type OrderTracking = shared.OrderTracking

// Product sold in the Terminal shop.
//
// This is an alias to an internal type.
type Product = shared.Product

// Whether the product must be or can be subscribed to.
//
// This is an alias to an internal type.
type ProductSubscription = shared.ProductSubscription

// This is an alias to an internal value.
const ProductSubscriptionAllowed = shared.ProductSubscriptionAllowed

// This is an alias to an internal value.
const ProductSubscriptionRequired = shared.ProductSubscriptionRequired

// Variant of a product in the Terminal shop.
//
// This is an alias to an internal type.
type ProductVariant = shared.ProductVariant

// Subscription to a Terminal shop product.
//
// This is an alias to an internal type.
type Subscription = shared.Subscription

// Frequency of the subscription.
//
// This is an alias to an internal type.
type SubscriptionFrequency = shared.SubscriptionFrequency

// This is an alias to an internal value.
const SubscriptionFrequencyFixed = shared.SubscriptionFrequencyFixed

// This is an alias to an internal value.
const SubscriptionFrequencyDaily = shared.SubscriptionFrequencyDaily

// This is an alias to an internal value.
const SubscriptionFrequencyWeekly = shared.SubscriptionFrequencyWeekly

// This is an alias to an internal value.
const SubscriptionFrequencyMonthly = shared.SubscriptionFrequencyMonthly

// This is an alias to an internal value.
const SubscriptionFrequencyYearly = shared.SubscriptionFrequencyYearly

// Subscription to a Terminal shop product.
//
// This is an alias to an internal type.
type SubscriptionParam = shared.SubscriptionParam

// A Terminal shop user. (We have users, btw.)
//
// This is an alias to an internal type.
type User = shared.User
