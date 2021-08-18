package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scheduleinterval
type Scheduleinterval struct { 
	// Start - The scheduled start time as an ISO-8601 string, i.e yyyy-MM-ddTHH:mm:ss.SSSZ
	Start *string `json:"start,omitempty"`


	// End - The scheduled end time as an ISO-8601 string, i.e. yyyy-MM-ddTHH:mm:ss.SSSZ
	End *string `json:"end,omitempty"`

}

func (u *Scheduleinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scheduleinterval

	

	return json.Marshal(&struct { 
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		*Alias
	}{ 
		Start: u.Start,
		
		End: u.End,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scheduleinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
