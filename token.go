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
)

// TokenService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	return
}

// Create a personal access token.
func (r *TokenService) New(ctx context.Context, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List the current user's personal access tokens.
func (r *TokenService) List(ctx context.Context, opts ...option.RequestOption) (res *TokenListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the personal access token with the given ID.
func (r *TokenService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("token/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the personal access token with the given ID.
func (r *TokenService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("token/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A personal access token used to access the Terminal API. If you leak this,
// expect large sums of coffee to be ordered on your credit card.
type Token struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Personal access token (obfuscated).
	Token string `json:"token,required"`
	// The created time for the token.
	Created string    `json:"created,required"`
	JSON    tokenJSON `json:"-"`
}

// tokenJSON contains the JSON metadata for the struct [Token]
type tokenJSON struct {
	ID          apijson.Field
	Token       apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Token) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponse struct {
	Data TokenNewResponseData `json:"data,required"`
	JSON tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseData struct {
	// Personal token ID.
	ID string `json:"id,required"`
	// Personal access token. Include this in the Authorization header
	// (`Bearer <token>`) when accessing the Terminal API.
	Token string                   `json:"token,required"`
	JSON  tokenNewResponseDataJSON `json:"-"`
}

// tokenNewResponseDataJSON contains the JSON metadata for the struct
// [TokenNewResponseData]
type tokenNewResponseDataJSON struct {
	ID          apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type TokenListResponse struct {
	// List of personal access tokens.
	Data []Token               `json:"data,required"`
	JSON tokenListResponseJSON `json:"-"`
}

// tokenListResponseJSON contains the JSON metadata for the struct
// [TokenListResponse]
type tokenListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenListResponseJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponse struct {
	Data TokenDeleteResponseData `json:"data,required"`
	JSON tokenDeleteResponseJSON `json:"-"`
}

// tokenDeleteResponseJSON contains the JSON metadata for the struct
// [TokenDeleteResponse]
type tokenDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseData string

const (
	TokenDeleteResponseDataOk TokenDeleteResponseData = "ok"
)

func (r TokenDeleteResponseData) IsKnown() bool {
	switch r {
	case TokenDeleteResponseDataOk:
		return true
	}
	return false
}

type TokenGetResponse struct {
	// A personal access token used to access the Terminal API. If you leak this,
	// expect large sums of coffee to be ordered on your credit card.
	Data Token                `json:"data,required"`
	JSON tokenGetResponseJSON `json:"-"`
}

// tokenGetResponseJSON contains the JSON metadata for the struct
// [TokenGetResponse]
type tokenGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseJSON) RawJSON() string {
	return r.raw
}
