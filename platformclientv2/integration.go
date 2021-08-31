package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integration - Details for an Integration
type Integration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Type of the integration
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// IntendedState - Configured state of the integration.
	IntendedState *string `json:"intendedState,omitempty"`


	// Config - Configuration information for the integration.
	Config *Integrationconfigurationinfo `json:"config,omitempty"`


	// ReportedState - Last reported status of the integration.
	ReportedState *Integrationstatusinfo `json:"reportedState,omitempty"`


	// Attributes - Read-only attributes for the integration.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Integration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationType *Integrationtype `json:"integrationType,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		IntendedState *string `json:"intendedState,omitempty"`
		
		Config *Integrationconfigurationinfo `json:"config,omitempty"`
		
		ReportedState *Integrationstatusinfo `json:"reportedState,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationType: o.IntegrationType,
		
		Notes: o.Notes,
		
		IntendedState: o.IntendedState,
		
		Config: o.Config,
		
		ReportedState: o.ReportedState,
		
		Attributes: o.Attributes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Integration) UnmarshalJSON(b []byte) error {
	var IntegrationMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := IntegrationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := IntegrationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if IntegrationType, ok := IntegrationMap["integrationType"].(map[string]interface{}); ok {
		IntegrationTypeString, _ := json.Marshal(IntegrationType)
		json.Unmarshal(IntegrationTypeString, &o.IntegrationType)
	}
	
	if Notes, ok := IntegrationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if IntendedState, ok := IntegrationMap["intendedState"].(string); ok {
		o.IntendedState = &IntendedState
	}
	
	if Config, ok := IntegrationMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if ReportedState, ok := IntegrationMap["reportedState"].(map[string]interface{}); ok {
		ReportedStateString, _ := json.Marshal(ReportedState)
		json.Unmarshal(ReportedStateString, &o.ReportedState)
	}
	
	if Attributes, ok := IntegrationMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if SelfUri, ok := IntegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
