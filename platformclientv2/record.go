package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Record
type Record struct { 
	// Name - The name of the record.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the record. (Example values:  MX, TXT, CNAME)
	VarType *string `json:"type,omitempty"`


	// Value - The value of the record.
	Value *string `json:"value,omitempty"`

}

func (o *Record) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Record
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Record) UnmarshalJSON(b []byte) error {
	var RecordMap map[string]interface{}
	err := json.Unmarshal(b, &RecordMap)
	if err != nil {
		return err
	}
	
	if Name, ok := RecordMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := RecordMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := RecordMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Record) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
