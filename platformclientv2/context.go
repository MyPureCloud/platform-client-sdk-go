package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Context
type Context struct { 
	// Patterns - A list of one or more patterns to match.
	Patterns *[]Contextpattern `json:"patterns,omitempty"`

}

func (u *Context) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Context

	

	return json.Marshal(&struct { 
		Patterns *[]Contextpattern `json:"patterns,omitempty"`
		*Alias
	}{ 
		Patterns: u.Patterns,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Context) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
