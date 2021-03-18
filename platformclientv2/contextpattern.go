package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contextpattern
type Contextpattern struct { 
	// Criteria - A list of one or more criteria to satisfy.
	Criteria *[]Entitytypecriteria `json:"criteria,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contextpattern) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
