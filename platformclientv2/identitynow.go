package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Identitynow
type Identitynow struct { 
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Identitynow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
