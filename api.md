# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#SubscriptionParam">SubscriptionParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Cart">Cart</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#ProductVariant">ProductVariant</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go/shared#User">User</a>

# Product

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="get /product">client.Product.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserUpdateResponse">UserUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserInitResponse">UserInitResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserMeResponse">UserMeResponse</a>

Methods:

- <code title="put /user/me">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserUpdateResponse">UserUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/init">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserService.Init">Init</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserInitResponse">UserInitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/me">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserMeResponse">UserMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipping

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingNewResponse">UserShippingNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingListResponse">UserShippingListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingDeleteResponse">UserShippingDeleteResponse</a>

Methods:

- <code title="post /user/shipping">client.User.Shipping.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingNewParams">UserShippingNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingNewResponse">UserShippingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/shipping">client.User.Shipping.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingListResponse">UserShippingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /user/shipping/{id}">client.User.Shipping.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#UserShippingDeleteResponse">UserShippingDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Card

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardNewResponse">CardNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardListResponse">CardListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardDeleteResponse">CardDeleteResponse</a>

Methods:

- <code title="post /card">client.Card.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardNewResponse">CardNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card">client.Card.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardListResponse">CardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /card/{id}">client.Card.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CardDeleteResponse">CardDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cart

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartListResponse">CartListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetCardResponse">CartSetCardResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetItemResponse">CartSetItemResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetShippingResponse">CartSetShippingResponse</a>

Methods:

- <code title="get /cart">client.Cart.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartListResponse">CartListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/card">client.Cart.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartService.SetCard">SetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetCardParams">CartSetCardParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetCardResponse">CartSetCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/item">client.Cart.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartService.SetItem">SetItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetItemParams">CartSetItemParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetItemResponse">CartSetItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/shipping">client.Cart.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartService.SetShipping">SetShipping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetShippingParams">CartSetShippingParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#CartSetShippingResponse">CartSetShippingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Order

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderNewResponse">OrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderGetResponse">OrderGetResponse</a>

Methods:

- <code title="post /order">client.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderNewResponse">OrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order">client.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order/{id}">client.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#OrderGetResponse">OrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscription

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>

Methods:

- <code title="put /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /subscription/{id}">client.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go">terminal</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/terminal-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
