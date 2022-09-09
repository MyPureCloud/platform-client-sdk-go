package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactaddressconditionsettings
type Contactaddressconditionsettings struct { 
	// Operator - The operator to use when comparing address values.
	Operator *string `json:"operator,omitempty"`


	// Value - The value to compare against the contact's address.
	Value *string `json:"value,omitempty"`

}

func (o *Contactaddressconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactaddressconditionsettings
	
	return json.Marshal(&struct { 
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Operator: o.Operator,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactaddressconditionsettings) UnmarshalJSON(b []byte) error {
	var ContactaddressconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ContactaddressconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if Operator, ok := ContactaddressconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactaddressconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactaddressconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
