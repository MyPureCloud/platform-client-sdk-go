package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationconfiguration - Configuration for an Integration
type Integrationconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
	Properties *interface{} `json:"properties,omitempty"`


	// Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
	Advanced *interface{} `json:"advanced,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
	Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Integrationconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationconfiguration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Properties *interface{} `json:"properties,omitempty"`
		
		Advanced *interface{} `json:"advanced,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		Properties: o.Properties,
		
		Advanced: o.Advanced,
		
		Notes: o.Notes,
		
		Credentials: o.Credentials,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Integrationconfiguration) UnmarshalJSON(b []byte) error {
	var IntegrationconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := IntegrationconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := IntegrationconfigurationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Version, ok := IntegrationconfigurationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Properties, ok := IntegrationconfigurationMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Advanced, ok := IntegrationconfigurationMap["advanced"].(map[string]interface{}); ok {
		AdvancedString, _ := json.Marshal(Advanced)
		json.Unmarshal(AdvancedString, &o.Advanced)
	}
	
	if Notes, ok := IntegrationconfigurationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if Credentials, ok := IntegrationconfigurationMap["credentials"].(map[string]interface{}); ok {
		CredentialsString, _ := json.Marshal(Credentials)
		json.Unmarshal(CredentialsString, &o.Credentials)
	}
	
	if SelfUri, ok := IntegrationconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
