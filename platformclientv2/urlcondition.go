package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Urlcondition
type Urlcondition struct { 
	// Values - The URL condition value.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

func (u *Urlcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Urlcondition

	

	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Values: u.Values,
		
		Operator: u.Operator,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Urlcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
