package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatususerparam
type Userroutingstatususerparam struct { 
	// Key
	Key *string `json:"key,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *Userroutingstatususerparam) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatususerparam
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingstatususerparam) UnmarshalJSON(b []byte) error {
	var UserroutingstatususerparamMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatususerparamMap)
	if err != nil {
		return err
	}
	
	if Key, ok := UserroutingstatususerparamMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if Value, ok := UserroutingstatususerparamMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatususerparam) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
