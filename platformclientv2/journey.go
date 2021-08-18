package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journey
type Journey struct { 
	// Patterns - A list of zero or more patterns to match.
	Patterns *[]Journeypattern `json:"patterns,omitempty"`

}

func (u *Journey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journey

	

	return json.Marshal(&struct { 
		Patterns *[]Journeypattern `json:"patterns,omitempty"`
		*Alias
	}{ 
		Patterns: u.Patterns,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Journey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
