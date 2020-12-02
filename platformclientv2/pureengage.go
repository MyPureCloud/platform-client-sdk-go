package platformclientv2
import (
	"encoding/json"
)

// Pureengage
type Pureengage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AutoProvisionUsers
	AutoProvisionUsers *bool `json:"autoProvisionUsers,omitempty"`


	// Certificate
	Certificate *string `json:"certificate,omitempty"`


	// SsoTargetURI
	SsoTargetURI *string `json:"ssoTargetURI,omitempty"`


	// IssuerURI
	IssuerURI *string `json:"issuerURI,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Pureengage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
