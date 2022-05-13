package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulermessagetypeseverity
type Schedulermessagetypeseverity struct { 
	// VarType - The type of the message
	VarType *string `json:"type,omitempty"`


	// Severity - The severity of the message
	Severity *string `json:"severity,omitempty"`

}

func (o *Schedulermessagetypeseverity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulermessagetypeseverity
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Severity *string `json:"severity,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Severity: o.Severity,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulermessagetypeseverity) UnmarshalJSON(b []byte) error {
	var SchedulermessagetypeseverityMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulermessagetypeseverityMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SchedulermessagetypeseverityMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Severity, ok := SchedulermessagetypeseverityMap["severity"].(string); ok {
		o.Severity = &Severity
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulermessagetypeseverity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
