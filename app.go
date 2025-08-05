// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomterminaldotshopterminalsdkgo

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

// AppService contains methods and other services that help with interacting with
// the terminal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Create an app.
func (r *AppService) New(ctx context.Context, body AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "app"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the current user's registered apps.
func (r *AppService) List(ctx context.Context, opts ...option.RequestOption) (res *AppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "app"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the app with the given ID.
func (r *AppService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("app/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the app with the given ID.
func (r *AppService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("app/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A Terminal App used for configuring an OAuth 2.0 client.
type App struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Name of the app.
	Name string `json:"name,required"`
	// Redirect URI of the app.
	RedirectUri string `json:"redirectURI,required"`
	// OAuth 2.0 client secret of the app (obfuscated).
	Secret string  `json:"secret,required"`
	JSON   appJSON `json:"-"`
}

// appJSON contains the JSON metadata for the struct [App]
type appJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RedirectUri apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *App) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appJSON) RawJSON() string {
	return r.raw
}

type AppNewResponse struct {
	Data AppNewResponseData `json:"data,required"`
	JSON appNewResponseJSON `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseData struct {
	// OAuth 2.0 client ID.
	ID string `json:"id,required"`
	// OAuth 2.0 client secret.
	Secret string                 `json:"secret,required"`
	JSON   appNewResponseDataJSON `json:"-"`
}

// appNewResponseDataJSON contains the JSON metadata for the struct
// [AppNewResponseData]
type appNewResponseDataJSON struct {
	ID          apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppListResponse struct {
	// List of apps.
	Data []App               `json:"data,required"`
	JSON appListResponseJSON `json:"-"`
}

// appListResponseJSON contains the JSON metadata for the struct [AppListResponse]
type appListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponse struct {
	Data AppDeleteResponseData `json:"data,required"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseData string

const (
	AppDeleteResponseDataOk AppDeleteResponseData = "ok"
)

func (r AppDeleteResponseData) IsKnown() bool {
	switch r {
	case AppDeleteResponseDataOk:
		return true
	}
	return false
}

type AppGetResponse struct {
	// A Terminal App used for configuring an OAuth 2.0 client.
	Data App                `json:"data,required"`
	JSON appGetResponseJSON `json:"-"`
}

// appGetResponseJSON contains the JSON metadata for the struct [AppGetResponse]
type appGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewParams struct {
	Name        param.Field[string] `json:"name,required"`
	RedirectUri param.Field[string] `json:"redirectURI,required"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
