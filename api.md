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

# Cart

Response Types:

- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>
- <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>

Methods:

- <code title="get /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartListResponse">CartListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cart">client.Cart.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartService.SetItem">SetItem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemParams">CartSetItemParams</a>) (<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go">terminal</a>.<a href="https://pkg.go.dev/github.com/terminaldotshop/terminal-sdk-go#CartSetItemResponse">CartSetItemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
