package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluconfusionmatrixcolumn
type Nluconfusionmatrixcolumn struct { 
	// Name - The name of the intent for the column.
	Name *string `json:"name,omitempty"`


	// Value - The confusion value between the intents
	Value *float32 `json:"value,omitempty"`

}

func (o *Nluconfusionmatrixcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluconfusionmatrixcolumn
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluconfusionmatrixcolumn) UnmarshalJSON(b []byte) error {
	var NluconfusionmatrixcolumnMap map[string]interface{}
	err := json.Unmarshal(b, &NluconfusionmatrixcolumnMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NluconfusionmatrixcolumnMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Value, ok := NluconfusionmatrixcolumnMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
