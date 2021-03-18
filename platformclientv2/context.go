package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Context
type Context struct { 
	// Patterns - A list of one or more patterns to match.
	Patterns *[]Contextpattern `json:"patterns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Context) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
