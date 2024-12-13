# Shared Params Types

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#SubscriptionParam">SubscriptionParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Cart">Cart</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#ProductVariant">ProductVariant</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#User">User</a>

# Product

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="get /product">client.Product.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserUpdateResponse">UserUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserInitResponse">UserInitResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserMeResponse">UserMeResponse</a>

Methods:

- <code title="put /user/me">client.User.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserUpdateResponse">UserUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/init">client.User.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserService.Init">Init</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserInitResponse">UserInitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/me">client.User.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserMeResponse">UserMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Address

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewResponse">AddressNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressListResponse">AddressListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressDeleteResponse">AddressDeleteResponse</a>

Methods:

- <code title="post /address">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewParams">AddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressNewResponse">AddressNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /address">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressListResponse">AddressListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /address/{id}">client.Address.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#AddressDeleteResponse">AddressDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Card

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardDeleteResponse">CardDeleteResponse</a>

Methods:

- <code title="post /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /card/{id}">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardDeleteResponse">CardDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cart

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressResponse">CartSetAddressResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>

Methods:

- <code title="get /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/address">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetAddress">SetAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressParams">CartSetAddressParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetAddressResponse">CartSetAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/card">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetCard">SetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardParams">CartSetCardParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/item">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetItem">SetItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemParams">CartSetItemParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Order

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderNewResponse">OrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderGetResponse">OrderGetResponse</a>

Methods:

- <code title="post /order">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderNewResponse">OrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /order/{id}">client.Order.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#OrderGetResponse">OrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscription

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>

Methods:

- <code title="put /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscription">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /subscription/{id}">client.Subscription.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
