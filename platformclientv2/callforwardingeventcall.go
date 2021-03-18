package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwardingeventcall
type Callforwardingeventcall struct { 
	// Targets
	Targets *[]Callforwardingeventtarget `json:"targets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callforwardingeventcall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
