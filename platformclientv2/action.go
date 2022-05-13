package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Action
type Action struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// IntegrationId - The ID of the integration for which this action is associated
	IntegrationId *string `json:"integrationId,omitempty"`


	// Category - Category of Action
	Category *string `json:"category,omitempty"`


	// Contract - Action contract
	Contract *Actioncontract `json:"contract,omitempty"`


	// Version - Version of this action
	Version *int `json:"version,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Action) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Action
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Contract *Actioncontract `json:"contract,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationId: o.IntegrationId,
		
		Category: o.Category,
		
		Contract: o.Contract,
		
		Version: o.Version,
		
		Secure: o.Secure,
		
		Config: o.Config,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Action) UnmarshalJSON(b []byte) error {
	var ActionMap map[string]interface{}
	err := json.Unmarshal(b, &ActionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IntegrationId, ok := ActionMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if Category, ok := ActionMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Contract, ok := ActionMap["contract"].(map[string]interface{}); ok {
		ContractString, _ := json.Marshal(Contract)
		json.Unmarshal(ContractString, &o.Contract)
	}
	
	if Version, ok := ActionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Secure, ok := ActionMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
    
	if Config, ok := ActionMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if SelfUri, ok := ActionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Action) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
