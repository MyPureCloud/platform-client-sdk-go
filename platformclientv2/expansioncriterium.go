package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Expansioncriterium
type Expansioncriterium struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Threshold
	Threshold *float64 `json:"threshold,omitempty"`

}

func (o *Expansioncriterium) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Expansioncriterium
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Threshold *float64 `json:"threshold,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Threshold: o.Threshold,
		Alias:    (*Alias)(o),
	})
}

func (o *Expansioncriterium) UnmarshalJSON(b []byte) error {
	var ExpansioncriteriumMap map[string]interface{}
	err := json.Unmarshal(b, &ExpansioncriteriumMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ExpansioncriteriumMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Threshold, ok := ExpansioncriteriumMap["threshold"].(float64); ok {
		o.Threshold = &Threshold
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Expansioncriterium) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
