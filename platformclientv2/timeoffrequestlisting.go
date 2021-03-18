package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestlisting
type Timeoffrequestlisting struct { 
	// Entities - List of time off request look up objects
	Entities *[]Timeoffrequest `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
