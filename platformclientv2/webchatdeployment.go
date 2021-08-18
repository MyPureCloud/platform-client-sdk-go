package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Webchatdeployment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatdeployment

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`
		
		AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		WebChatConfig *Webchatconfig `json:"webChatConfig,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		AuthenticationRequired: u.AuthenticationRequired,
		
		AuthenticationUrl: u.AuthenticationUrl,
		
		Disabled: u.Disabled,
		
		WebChatConfig: u.WebChatConfig,
		
		AllowedDomains: u.AllowedDomains,
		
		Flow: u.Flow,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
