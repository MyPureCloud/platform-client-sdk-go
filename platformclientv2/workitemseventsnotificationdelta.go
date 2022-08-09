package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemseventsnotificationdelta
type Workitemseventsnotificationdelta struct { 
	// Op
	Op *string `json:"op,omitempty"`


	// Field
	Field *string `json:"field,omitempty"`


	// OldValue
	OldValue *string `json:"oldValue,omitempty"`


	// NewValue
	NewValue *string `json:"newValue,omitempty"`

}

func (o *Workitemseventsnotificationdelta) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workitemseventsnotificationdelta
	
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

func (o *Workitemseventsnotificationdelta) UnmarshalJSON(b []byte) error {
	var WorkitemseventsnotificationdeltaMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemseventsnotificationdeltaMap)
	if err != nil {
		return err
	}
	
	if Op, ok := WorkitemseventsnotificationdeltaMap["op"].(string); ok {
		o.Op = &Op
	}
    
	if Field, ok := WorkitemseventsnotificationdeltaMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if OldValue, ok := WorkitemseventsnotificationdeltaMap["oldValue"].(string); ok {
		o.OldValue = &OldValue
	}
    
	if NewValue, ok := WorkitemseventsnotificationdeltaMap["newValue"].(string); ok {
		o.NewValue = &NewValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemseventsnotificationdelta) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
