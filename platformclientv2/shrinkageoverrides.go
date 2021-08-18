package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shrinkageoverrides
type Shrinkageoverrides struct { 
	// Clear - Set true to clear the shrinkage interval overrides
	Clear *bool `json:"clear,omitempty"`


	// Values - List of interval shrinkage overrides
	Values *[]Shrinkageoverride `json:"values,omitempty"`

}

func (u *Shrinkageoverrides) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shrinkageoverrides

	

	return json.Marshal(&struct { 
		Clear *bool `json:"clear,omitempty"`
		
		Values *[]Shrinkageoverride `json:"values,omitempty"`
		*Alias
	}{ 
		Clear: u.Clear,
		
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shrinkageoverrides) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
