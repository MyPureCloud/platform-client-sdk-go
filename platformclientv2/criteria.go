package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Criteria
type Criteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

func (u *Criteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Criteria

	

	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Key: u.Key,
		
		Values: u.Values,
		
		ShouldIgnoreCase: u.ShouldIgnoreCase,
		
		Operator: u.Operator,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Criteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
