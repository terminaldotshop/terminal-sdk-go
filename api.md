# Shared Params Types

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#AddressParam">AddressParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Cart">Cart</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#ProductVariant">ProductVariant</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go/shared#User">User</a>

# Product

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="get /product">client.Product.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserMeResponse">UserMeResponse</a>

Methods:

- <code title="get /user/me">client.User.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserMeResponse">UserMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipping

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingNewResponse">UserShippingNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingListResponse">UserShippingListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingDeleteResponse">UserShippingDeleteResponse</a>

Methods:

- <code title="post /user/shipping">client.User.Shipping.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingNewParams">UserShippingNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingNewResponse">UserShippingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/shipping">client.User.Shipping.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingListResponse">UserShippingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /user/shipping/{id}">client.User.Shipping.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#UserShippingDeleteResponse">UserShippingDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Card

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>

Methods:

- <code title="post /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardNewResponse">CardNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card">client.Card.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CardListResponse">CardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cart

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetShippingResponse">CartSetShippingResponse</a>

Methods:

- <code title="get /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/card">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetCard">SetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardParams">CartSetCardParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetCardResponse">CartSetCardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/item">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetItem">SetItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemParams">CartSetItemParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart/shipping">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetShipping">SetShipping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetShippingParams">CartSetShippingParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetShippingResponse">CartSetShippingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
