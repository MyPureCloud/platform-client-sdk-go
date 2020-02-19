package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
