package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryfilter
type Auditqueryfilter struct { 
	// Property - Name of the property to filter.
	Property *string `json:"property,omitempty"`


	// Value - Value of the property to filter.
	Value *string `json:"value,omitempty"`

}

func (u *Auditqueryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryfilter

	

	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Property: u.Property,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
