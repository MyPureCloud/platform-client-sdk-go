package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsefilter - Used to filter response queries
type Responsefilter struct { 
	// Name - Field to filter on. Allowed values are 'name' and 'libraryId.
	Name *string `json:"name,omitempty"`


	// Operator - Filter operation: IN, EQUALS, NOTEQUALS.
	Operator *string `json:"operator,omitempty"`


	// Values - Values to filter on.
	Values *[]string `json:"values,omitempty"`

}

func (u *Responsefilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsefilter

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Operator: u.Operator,
		
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Responsefilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
