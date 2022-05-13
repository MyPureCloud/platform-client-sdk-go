package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Limit
type Limit struct { 
	// Key - The limit key
	Key *string `json:"key,omitempty"`


	// Value - The limit value
	Value *float64 `json:"value,omitempty"`

}

func (o *Limit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Limit
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Limit) UnmarshalJSON(b []byte) error {
	var LimitMap map[string]interface{}
	err := json.Unmarshal(b, &LimitMap)
	if err != nil {
		return err
	}
	
	if Key, ok := LimitMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if Value, ok := LimitMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Limit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
