package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactaddresstypeconditionsettings
type Contactaddresstypeconditionsettings struct { 
	// Operator - The operator to use when comparing the address types.
	Operator *string `json:"operator,omitempty"`


	// Value - The type value to compare against the contact column type.
	Value *string `json:"value,omitempty"`

}

func (o *Contactaddresstypeconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactaddresstypeconditionsettings
	
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

func (o *Contactaddresstypeconditionsettings) UnmarshalJSON(b []byte) error {
	var ContactaddresstypeconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ContactaddresstypeconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if Operator, ok := ContactaddresstypeconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactaddresstypeconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactaddresstypeconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
