package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfig - Defines a SCIM service provider's configuration.
type Scimserviceproviderconfig struct { 
	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// DocumentationUri - The HTTP-addressable URL that points to the service provider's documentation.
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// Patch - The \"patch\" configuration options.
	Patch *Scimserviceproviderconfigsimplefeature `json:"patch,omitempty"`


	// Filter - The \"filter\" configuration options.
	Filter *Scimserviceproviderconfigfilterfeature `json:"filter,omitempty"`


	// Etag - The \"etag\" configuration options.
	Etag *Scimserviceproviderconfigsimplefeature `json:"etag,omitempty"`


	// Sort - The \"sort\" configuration options.
	Sort *Scimserviceproviderconfigsimplefeature `json:"sort,omitempty"`


	// Bulk - The \"bulk\" configuration options.
	Bulk *Scimserviceproviderconfigbulkfeature `json:"bulk,omitempty"`


	// ChangePassword - The \"changePassword\" configuration options.
	ChangePassword *Scimserviceproviderconfigsimplefeature `json:"changePassword,omitempty"`


	// AuthenticationSchemes - The list of supported authentication schemes.
	AuthenticationSchemes *[]Scimserviceproviderconfigauthenticationscheme `json:"authenticationSchemes,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

func (u *Scimserviceproviderconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfig

	

	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		Patch *Scimserviceproviderconfigsimplefeature `json:"patch,omitempty"`
		
		Filter *Scimserviceproviderconfigfilterfeature `json:"filter,omitempty"`
		
		Etag *Scimserviceproviderconfigsimplefeature `json:"etag,omitempty"`
		
		Sort *Scimserviceproviderconfigsimplefeature `json:"sort,omitempty"`
		
		Bulk *Scimserviceproviderconfigbulkfeature `json:"bulk,omitempty"`
		
		ChangePassword *Scimserviceproviderconfigsimplefeature `json:"changePassword,omitempty"`
		
		AuthenticationSchemes *[]Scimserviceproviderconfigauthenticationscheme `json:"authenticationSchemes,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Schemas: u.Schemas,
		
		DocumentationUri: u.DocumentationUri,
		
		Patch: u.Patch,
		
		Filter: u.Filter,
		
		Etag: u.Etag,
		
		Sort: u.Sort,
		
		Bulk: u.Bulk,
		
		ChangePassword: u.ChangePassword,
		
		AuthenticationSchemes: u.AuthenticationSchemes,
		
		Meta: u.Meta,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
