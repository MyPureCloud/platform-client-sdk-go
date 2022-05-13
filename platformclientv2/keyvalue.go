package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Keyvalue
type Keyvalue struct { 
	// Key - Key for free-form data.
	Key *string `json:"key,omitempty"`


	// Value - Value for free-form data.
	Value *string `json:"value,omitempty"`

}

func (o *Keyvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Keyvalue
	
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

func (o *Keyvalue) UnmarshalJSON(b []byte) error {
	var KeyvalueMap map[string]interface{}
	err := json.Unmarshal(b, &KeyvalueMap)
	if err != nil {
		return err
	}
	
	if Key, ok := KeyvalueMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if Value, ok := KeyvalueMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Keyvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
