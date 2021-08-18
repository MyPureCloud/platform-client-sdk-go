package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigauthenticationscheme - Defines an authentication scheme in the SCIM service provider's configuration.
type Scimserviceproviderconfigauthenticationscheme struct { 
	// Name - The name of the authentication scheme, for example, HTTP Basic.
	Name *string `json:"name,omitempty"`


	// Description - The description of the authentication scheme.
	Description *string `json:"description,omitempty"`


	// SpecUri - The HTTP-addressable URL that points to the authentication scheme's specification.
	SpecUri *string `json:"specUri,omitempty"`


	// DocumentationUri - The HTTP-addressable URL that points to the authentication scheme's usage documentation.
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// VarType - The type of authentication scheme.
	VarType *string `json:"type,omitempty"`


	// Primary - Indicates whether this authentication scheme is the primary method of authentication.
	Primary *bool `json:"primary,omitempty"`

}

func (u *Scimserviceproviderconfigauthenticationscheme) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigauthenticationscheme

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SpecUri *string `json:"specUri,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		SpecUri: u.SpecUri,
		
		DocumentationUri: u.DocumentationUri,
		
		VarType: u.VarType,
		
		Primary: u.Primary,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigauthenticationscheme) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
