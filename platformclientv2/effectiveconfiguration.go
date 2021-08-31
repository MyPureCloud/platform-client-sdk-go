package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Effectiveconfiguration - Effective Configuration for an ClientApp. This is comprised of the integration specific configuration along with overrides specified in the integration type.
type Effectiveconfiguration struct { 
	// Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
	Advanced *map[string]interface{} `json:"advanced,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
	Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`

}

func (o *Effectiveconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Effectiveconfiguration
	
	return json.Marshal(&struct { 
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		Advanced *map[string]interface{} `json:"advanced,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`
		*Alias
	}{ 
		Properties: o.Properties,
		
		Advanced: o.Advanced,
		
		Name: o.Name,
		
		Notes: o.Notes,
		
		Credentials: o.Credentials,
		Alias:    (*Alias)(o),
	})
}

func (o *Effectiveconfiguration) UnmarshalJSON(b []byte) error {
	var EffectiveconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &EffectiveconfigurationMap)
	if err != nil {
		return err
	}
	
	if Properties, ok := EffectiveconfigurationMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Advanced, ok := EffectiveconfigurationMap["advanced"].(map[string]interface{}); ok {
		AdvancedString, _ := json.Marshal(Advanced)
		json.Unmarshal(AdvancedString, &o.Advanced)
	}
	
	if Name, ok := EffectiveconfigurationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Notes, ok := EffectiveconfigurationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if Credentials, ok := EffectiveconfigurationMap["credentials"].(map[string]interface{}); ok {
		CredentialsString, _ := json.Marshal(Credentials)
		json.Unmarshal(CredentialsString, &o.Credentials)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Effectiveconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
