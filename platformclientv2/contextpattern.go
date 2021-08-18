package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contextpattern
type Contextpattern struct { 
	// Criteria - A list of one or more criteria to satisfy.
	Criteria *[]Entitytypecriteria `json:"criteria,omitempty"`

}

func (u *Contextpattern) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contextpattern

	

	return json.Marshal(&struct { 
		Criteria *[]Entitytypecriteria `json:"criteria,omitempty"`
		*Alias
	}{ 
		Criteria: u.Criteria,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contextpattern) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
