package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemsusereventsnotificationdelta
type Workitemsusereventsnotificationdelta struct { 
	// Op
	Op *string `json:"op,omitempty"`


	// Field
	Field *string `json:"field,omitempty"`


	// OldValue
	OldValue *string `json:"oldValue,omitempty"`


	// NewValue
	NewValue *string `json:"newValue,omitempty"`

}

func (o *Workitemsusereventsnotificationdelta) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workitemsusereventsnotificationdelta
	
	return json.Marshal(&struct { 
		Op *string `json:"op,omitempty"`
		
		Field *string `json:"field,omitempty"`
		
		OldValue *string `json:"oldValue,omitempty"`
		
		NewValue *string `json:"newValue,omitempty"`
		*Alias
	}{ 
		Op: o.Op,
		
		Field: o.Field,
		
		OldValue: o.OldValue,
		
		NewValue: o.NewValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Workitemsusereventsnotificationdelta) UnmarshalJSON(b []byte) error {
	var WorkitemsusereventsnotificationdeltaMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemsusereventsnotificationdeltaMap)
	if err != nil {
		return err
	}
	
	if Op, ok := WorkitemsusereventsnotificationdeltaMap["op"].(string); ok {
		o.Op = &Op
	}
    
	if Field, ok := WorkitemsusereventsnotificationdeltaMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if OldValue, ok := WorkitemsusereventsnotificationdeltaMap["oldValue"].(string); ok {
		o.OldValue = &OldValue
	}
    
	if NewValue, ok := WorkitemsusereventsnotificationdeltaMap["newValue"].(string); ok {
		o.NewValue = &NewValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemsusereventsnotificationdelta) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
