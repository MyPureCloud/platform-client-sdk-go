package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmodetrigger - Basic identifying information about a trigger
type Testmodetrigger struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the trigger
	Name *string `json:"name,omitempty"`


	// Enabled - Whether or not the trigger is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Testmodetrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmodetrigger
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Enabled: o.Enabled,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Testmodetrigger) UnmarshalJSON(b []byte) error {
	var TestmodetriggerMap map[string]interface{}
	err := json.Unmarshal(b, &TestmodetriggerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TestmodetriggerMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TestmodetriggerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := TestmodetriggerMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SelfUri, ok := TestmodetriggerMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Testmodetrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
