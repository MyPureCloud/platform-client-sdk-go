package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingannotationflow
type Recordingannotationflow struct { 
	// Name - The flow name
	Name *string `json:"name,omitempty"`


	// Id - The flow Id
	Id *string `json:"id,omitempty"`

}

func (o *Recordingannotationflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingannotationflow
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingannotationflow) UnmarshalJSON(b []byte) error {
	var RecordingannotationflowMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingannotationflowMap)
	if err != nil {
		return err
	}
	
	if Name, ok := RecordingannotationflowMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Id, ok := RecordingannotationflowMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingannotationflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
