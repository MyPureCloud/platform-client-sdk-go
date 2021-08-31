package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Check
type Check struct { 
	// Result - The result of a check executed. This indicates if the check was successful or not.
	Result *string `json:"result,omitempty"`


	// VarType - The type of check executed.
	VarType *string `json:"type,omitempty"`

}

func (o *Check) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Check
	
	return json.Marshal(&struct { 
		Result *string `json:"result,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Result: o.Result,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Check) UnmarshalJSON(b []byte) error {
	var CheckMap map[string]interface{}
	err := json.Unmarshal(b, &CheckMap)
	if err != nil {
		return err
	}
	
	if Result, ok := CheckMap["result"].(string); ok {
		o.Result = &Result
	}
	
	if VarType, ok := CheckMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Check) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
