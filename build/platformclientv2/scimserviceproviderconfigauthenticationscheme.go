package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigauthenticationscheme) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
