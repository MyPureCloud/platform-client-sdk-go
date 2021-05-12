package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradepreviewresponse
type Shifttradepreviewresponse struct { 
	// Activities - List of activities that will make up the new shift if this shift trade is approved
	Activities *[]Shifttradeactivitypreviewresponse `json:"activities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradepreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
