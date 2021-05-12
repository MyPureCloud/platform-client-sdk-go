package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulelistresponse - Week schedule list
type Weekschedulelistresponse struct { 
	// Entities
	Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekschedulelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
