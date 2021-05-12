package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitactivitycodelisting - List of BusinessUnitActivityCode
type Businessunitactivitycodelisting struct { 
	// Entities
	Entities *[]Businessunitactivitycode `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Businessunitactivitycodelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
