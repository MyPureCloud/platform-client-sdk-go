package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resultcounters
type Resultcounters struct { 
	// Success
	Success *int `json:"success,omitempty"`


	// Failure
	Failure *int `json:"failure,omitempty"`

}

func (u *Resultcounters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resultcounters

	

	return json.Marshal(&struct { 
		Success *int `json:"success,omitempty"`
		
		Failure *int `json:"failure,omitempty"`
		*Alias
	}{ 
		Success: u.Success,
		
		Failure: u.Failure,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Resultcounters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
