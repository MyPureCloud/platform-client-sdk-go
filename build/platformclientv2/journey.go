package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Journey
type Journey struct { 
	// Patterns - A list of zero or more patterns to match.
	Patterns *[]Journeypattern `json:"patterns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
