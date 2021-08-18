package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Widgetdeployment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetdeployment

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		ClientType *string `json:"clientType,omitempty"`
		
		ClientConfig *Widgetclientconfig `json:"clientConfig,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		AuthenticationRequired: u.AuthenticationRequired,
		
		Disabled: u.Disabled,
		
		Flow: u.Flow,
		
		AllowedDomains: u.AllowedDomains,
		
		ClientType: u.ClientType,
		
		ClientConfig: u.ClientConfig,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Widgetdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
