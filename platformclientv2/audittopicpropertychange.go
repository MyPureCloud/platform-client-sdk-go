package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicpropertychange
type Audittopicpropertychange struct { 
	// Property
	Property *string `json:"property,omitempty"`


	// OldValues
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues
	NewValues *[]string `json:"newValues,omitempty"`

}

func (u *Audittopicpropertychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicpropertychange

	

	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Property: u.Property,
		
		OldValues: u.OldValues,
		
		NewValues: u.NewValues,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Audittopicpropertychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
