package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Items
type Items struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Pattern
	Pattern *string `json:"pattern,omitempty"`

}

func (o *Items) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Items
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Pattern: o.Pattern,
		Alias:    (*Alias)(o),
	})
}

func (o *Items) UnmarshalJSON(b []byte) error {
	var ItemsMap map[string]interface{}
	err := json.Unmarshal(b, &ItemsMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ItemsMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Pattern, ok := ItemsMap["pattern"].(string); ok {
		o.Pattern = &Pattern
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Items) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
