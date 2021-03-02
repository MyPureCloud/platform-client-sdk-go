package platformclientv2
import (
	"encoding/json"
)

// Genericsaml
type Genericsaml struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// IssuerURI
	IssuerURI *string `json:"issuerURI,omitempty"`


	// SsoTargetURI
	SsoTargetURI *string `json:"ssoTargetURI,omitempty"`


	// Certificate
	Certificate *string `json:"certificate,omitempty"`


	// Certificates
	Certificates *[]string `json:"certificates,omitempty"`


	// RelyingPartyIdentifier
	RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`


	// LogoImageData
	LogoImageData *string `json:"logoImageData,omitempty"`


	// EndpointCompression
	EndpointCompression *bool `json:"endpointCompression,omitempty"`


	// NameIdentifierFormat
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Genericsaml) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
