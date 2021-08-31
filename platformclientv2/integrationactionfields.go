package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationactionfields
type Integrationactionfields struct { 
	// IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
	IntegrationAction *Integrationaction `json:"integrationAction,omitempty"`


	// RequestMappings - Collection of Request Mappings to use
	RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`

}

func (o *Integrationactionfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationactionfields
	
	return json.Marshal(&struct { 
		IntegrationAction *Integrationaction `json:"integrationAction,omitempty"`
		
		RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`
		*Alias
	}{ 
		IntegrationAction: o.IntegrationAction,
		
		RequestMappings: o.RequestMappings,
		Alias:    (*Alias)(o),
	})
}

func (o *Integrationactionfields) UnmarshalJSON(b []byte) error {
	var IntegrationactionfieldsMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationactionfieldsMap)
	if err != nil {
		return err
	}
	
	if IntegrationAction, ok := IntegrationactionfieldsMap["integrationAction"].(map[string]interface{}); ok {
		IntegrationActionString, _ := json.Marshal(IntegrationAction)
		json.Unmarshal(IntegrationActionString, &o.IntegrationAction)
	}
	
	if RequestMappings, ok := IntegrationactionfieldsMap["requestMappings"].([]interface{}); ok {
		RequestMappingsString, _ := json.Marshal(RequestMappings)
		json.Unmarshal(RequestMappingsString, &o.RequestMappings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationactionfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
