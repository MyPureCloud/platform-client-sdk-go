package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucopyschedulerequest
type Bucopyschedulerequest struct { 
	// Description - The description for the new schedule
	Description *string `json:"description,omitempty"`


	// WeekDate - The start weekDate for the new copy of the schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

func (u *Bucopyschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucopyschedulerequest

	
	WeekDate := new(string)
	if u.WeekDate != nil {
		*WeekDate = timeutil.Strftime(u.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	

	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Description: u.Description,
		
		WeekDate: WeekDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bucopyschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
