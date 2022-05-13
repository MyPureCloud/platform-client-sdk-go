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

func (o *Webchatdeployment) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		AuthenticationRequired: o.AuthenticationRequired,
		
		AuthenticationUrl: o.AuthenticationUrl,
		
		Disabled: o.Disabled,
		
		WebChatConfig: o.WebChatConfig,
		
		AllowedDomains: o.AllowedDomains,
		
		Flow: o.Flow,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatdeployment) UnmarshalJSON(b []byte) error {
	var WebchatdeploymentMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatdeploymentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatdeploymentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebchatdeploymentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := WebchatdeploymentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if AuthenticationRequired, ok := WebchatdeploymentMap["authenticationRequired"].(bool); ok {
		o.AuthenticationRequired = &AuthenticationRequired
	}
    
	if AuthenticationUrl, ok := WebchatdeploymentMap["authenticationUrl"].(string); ok {
		o.AuthenticationUrl = &AuthenticationUrl
	}
    
	if Disabled, ok := WebchatdeploymentMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if WebChatConfig, ok := WebchatdeploymentMap["webChatConfig"].(map[string]interface{}); ok {
		WebChatConfigString, _ := json.Marshal(WebChatConfig)
		json.Unmarshal(WebChatConfigString, &o.WebChatConfig)
	}
	
	if AllowedDomains, ok := WebchatdeploymentMap["allowedDomains"].([]interface{}); ok {
		AllowedDomainsString, _ := json.Marshal(AllowedDomains)
		json.Unmarshal(AllowedDomainsString, &o.AllowedDomains)
	}
	
	if Flow, ok := WebchatdeploymentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if SelfUri, ok := WebchatdeploymentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
