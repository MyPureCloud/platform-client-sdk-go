package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetdeployment
type Widgetdeployment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - A human-readable description of this Deployment.
	Description *string `json:"description,omitempty"`


	// AuthenticationRequired - When true, the customer members starting a chat must be authenticated by supplying their JWT to the create operation.
	AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`


	// Disabled - When true, all create chat operations using this Deployment will be rejected.
	Disabled *bool `json:"disabled,omitempty"`


	// Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
	Flow *Domainentityref `json:"flow,omitempty"`


	// AllowedDomains - The list of domains that are approved to use this Deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`


	// ClientType - The type of display widget for which this Deployment is configured, which controls the administrator settings shown.
	ClientType *string `json:"clientType,omitempty"`


	// ClientConfig - The client configuration options that should be made available to the clients of this Deployment.
	ClientConfig *Widgetclientconfig `json:"clientConfig,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Widgetdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
