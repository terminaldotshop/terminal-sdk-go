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
