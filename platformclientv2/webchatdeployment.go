package platformclientv2
import (
	"encoding/json"
)

// Webchatdeployment
type Webchatdeployment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// AuthenticationRequired
	AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`


	// AuthenticationUrl - URL for third party service authenticating web chat clients. See https://github.com/MyPureCloud/authenticated-web-chat-server-examples
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// WebChatConfig
	WebChatConfig *Webchatconfig `json:"webChatConfig,omitempty"`


	// AllowedDomains
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`


	// Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
	Flow *Domainentityref `json:"flow,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatdeployment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
