package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserrole - Defines a user role.
type Scimuserrole struct { 
	// Value - The role of the Genesys Cloud user.
	Value *string `json:"value,omitempty"`

}

func (o *Scimuserrole) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimuserrole
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimuserrole) UnmarshalJSON(b []byte) error {
	var ScimuserroleMap map[string]interface{}
	err := json.Unmarshal(b, &ScimuserroleMap)
	if err != nil {
		return err
	}
	
	if Value, ok := ScimuserroleMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimuserrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
