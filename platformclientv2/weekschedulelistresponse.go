package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulelistresponse - Week schedule list
type Weekschedulelistresponse struct { 
	// Entities
	Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`

}

func (u *Weekschedulelistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedulelistresponse

	

	return json.Marshal(&struct { 
		Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Weekschedulelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
