package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchessummaryresponse
type Shifttradematchessummaryresponse struct { 
	// Entities
	Entities *[]Weekshifttradematchessummaryresponse `json:"entities,omitempty"`

}

func (u *Shifttradematchessummaryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchessummaryresponse

	

	return json.Marshal(&struct { 
		Entities *[]Weekshifttradematchessummaryresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradematchessummaryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
