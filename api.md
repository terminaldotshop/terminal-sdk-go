# Product

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductVariant">ProductVariant</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductGetResponse">ProductGetResponse</a>

Methods:

- <code title="get /product">client.Product.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /product/{id}">client.Product.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductGetResponse">ProductGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Profile">Profile</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileUpdateResponse">ProfileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileMeResponse">ProfileMeResponse</a>

Methods:

- <code title="put /profile">client.Profile.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileUpdateParams">ProfileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileUpdateResponse">ProfileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /profile">client.Profile.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProfileMeResponse">ProfileMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Address

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewResponse">AddressNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressListResponse">AddressListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressDeleteResponse">AddressDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressGetResponse">AddressGetResponse</a>

Methods:

- <code title="post /address">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewParams">AddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewResponse">AddressNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /address">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressListResponse">AddressListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /address/{id}">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressDeleteResponse">AddressDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /address/{id}">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressGetResponse">AddressGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Card

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardDeleteResponse">CardDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardCollectResponse">CardCollectResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardGetResponse">CardGetResponse</a>

Methods:

- <code title="post /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /card/{id}">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardDeleteResponse">CardDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /card/collect">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.Collect">Collect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardCollectResponse">CardCollectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card/{id}">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardGetResponse">CardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cart

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Cart">Cart</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartClearResponse">CartClearResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartConvertResponse">CartConvertResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartGetResponse">CartGetResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartRedeemGiftCardResponse">CartRedeemGiftCardResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartRemoveGiftCardResponse">CartRemoveGiftCardResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressResponse">CartSetAddressResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>

Methods:

- <code title="delete /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.Clear">Clear</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartClearResponse">CartClearResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cart/convert">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartConvertParams">CartConvertParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartConvertResponse">CartConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartGetResponse">CartGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/gift-card">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.RedeemGiftCard">RedeemGiftCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartRedeemGiftCardParams">CartRedeemGiftCardParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartRedeemGiftCardResponse">CartRedeemGiftCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cart/gift-card">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.RemoveGiftCard">RemoveGiftCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartRemoveGiftCardResponse">CartRemoveGiftCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/address">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetAddress">SetAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressParams">CartSetAddressParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressResponse">CartSetAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/card">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetCard">SetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardParams">CartSetCardParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/item">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetItem">SetItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemParams">CartSetItemParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Order

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderNewResponse">OrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderGetResponse">OrderGetResponse</a>

Methods:

- <code title="post /order">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderNewResponse">OrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order/{id}">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderGetResponse">OrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscription

Params Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionParam">SubscriptionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionGetResponse">SubscriptionGetResponse</a>

Methods:

- <code title="post /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /subscription/{id}">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscription/{id}">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionGetResponse">SubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Token

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#Token">Token</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenNewResponse">TokenNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenListResponse">TokenListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenDeleteResponse">TokenDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenGetResponse">TokenGetResponse</a>

Methods:

- <code title="post /token">client.Token.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenNewResponse">TokenNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /token">client.Token.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenListResponse">TokenListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /token/{id}">client.Token.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenDeleteResponse">TokenDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /token/{id}">client.Token.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#TokenGetResponse">TokenGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#App">App</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppNewResponse">AppNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppListResponse">AppListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppDeleteResponse">AppDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppGetResponse">AppGetResponse</a>

Methods:

- <code title="post /app">client.App.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppNewParams">AppNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppNewResponse">AppNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /app">client.App.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppListResponse">AppListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /app/{id}">client.App.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppDeleteResponse">AppDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /app/{id}">client.App.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AppGetResponse">AppGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Email

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#EmailNewResponse">EmailNewResponse</a>

Methods:

- <code title="post /email">client.Email.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#EmailService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#EmailNewParams">EmailNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#EmailNewResponse">EmailNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# View

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ViewInitResponse">ViewInitResponse</a>

Methods:

- <code title="get /view/init">client.View.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ViewService.Init">Init</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ViewInitResponse">ViewInitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
