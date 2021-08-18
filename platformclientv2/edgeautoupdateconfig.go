package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
	// TimeZone - The timezone of the window in which any updates to the edges assigned to the site can be applied. The minimum size of the window is 2 hours.
	TimeZone *string `json:"timeZone,omitempty"`


	// Rrule - The recurrence rule for updating the Edges assigned to the site. The only supported frequencies are daily and weekly. Weekly frequencies require a day list with at least oneday specified. All other configurations are not supported.
	Rrule *string `json:"rrule,omitempty"`


	// Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	Start *time.Time `json:"start,omitempty"`


	// End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	End *time.Time `json:"end,omitempty"`

}

func (u *Edgeautoupdateconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeautoupdateconfig

	
	Start := new(string)
	if u.Start != nil {
		*Start = timeutil.Strftime(u.Start, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		Start = nil
	}
	
	End := new(string)
	if u.End != nil {
		*End = timeutil.Strftime(u.End, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		End = nil
	}
	

	return json.Marshal(&struct { 
		TimeZone *string `json:"timeZone,omitempty"`
		
		Rrule *string `json:"rrule,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		*Alias
	}{ 
		TimeZone: u.TimeZone,
		
		Rrule: u.Rrule,
		
		Start: Start,
		
		End: End,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgeautoupdateconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
