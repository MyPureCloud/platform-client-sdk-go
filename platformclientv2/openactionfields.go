package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openactionfields
type Openactionfields struct { 
	// OpenAction - The specific type of the open action.
	OpenAction *Domainentityref `json:"openAction,omitempty"`


	// ConfigurationFields - Custom fields defined in the schema referenced by the open action type selected.
	ConfigurationFields *map[string]interface{} `json:"configurationFields,omitempty"`

}

func (o *Openactionfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openactionfields
	
	return json.Marshal(&struct { 
		OpenAction *Domainentityref `json:"openAction,omitempty"`
		
		ConfigurationFields *map[string]interface{} `json:"configurationFields,omitempty"`
		*Alias
	}{ 
		OpenAction: o.OpenAction,
		
		ConfigurationFields: o.ConfigurationFields,
		Alias:    (*Alias)(o),
	})
}

func (o *Openactionfields) UnmarshalJSON(b []byte) error {
	var OpenactionfieldsMap map[string]interface{}
	err := json.Unmarshal(b, &OpenactionfieldsMap)
	if err != nil {
		return err
	}
	
	if OpenAction, ok := OpenactionfieldsMap["openAction"].(map[string]interface{}); ok {
		OpenActionString, _ := json.Marshal(OpenAction)
		json.Unmarshal(OpenActionString, &o.OpenAction)
	}
	
	if ConfigurationFields, ok := OpenactionfieldsMap["configurationFields"].(map[string]interface{}); ok {
		ConfigurationFieldsString, _ := json.Marshal(ConfigurationFields)
		json.Unmarshal(ConfigurationFieldsString, &o.ConfigurationFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openactionfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
